package service

import (
	"encoding/json"
	"errors"
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"gin-vben-admin/dao"
	"gin-vben-admin/dto"
	"go.uber.org/zap"
	"strconv"
	"time"
)
// 获取角色所有菜单列表
func GetRoleMenuTree(RoleId int) (menuTreeList []*dto.SysMenu, err error) {
	if val,err1:=common.CACHE.Get("sideMenu").Result();err1!=nil{
		var roleMenuList []*dao.SysMenu
		err=common.DB.Model(&dao.SysRole{ID:RoleId}).Association("SysMenu").Find(&roleMenuList).Error
		if err!=nil{
			return
		}
		newMenuList,_ :=getSideMenuList(roleMenuList)
		menuTreeList=generateMenuTree(newMenuList)
		res,_:=json.Marshal(menuTreeList)
		errRedis:=common.CACHE.Set("sideMenu"+strconv.Itoa(RoleId),res,60*60*24*time.Second).Err()
		//设置缓存，如果缓存失败，打印错误信息到日志
		if errRedis!=nil{
			common.LOG.Warn("Cache Redis error",zap.Any("warn",errRedis))
		}

		return
	}else{
		err =json.Unmarshal([]byte(val),&menuTreeList)
		return
	}
}
func GetRoleList(q dto.QuerySysRole) (roles []*dto.SysRole,total int, err error) {

	var daoRole []*dao.SysRole
    query:=common.DB.Model(&daoRole)
	if q.Status!="" {
		query = query.Where("status = ?", q.Status).Count(&total)
	}
	if  q.RoleName!="" {
		query = query.Where("role_name like ?", "%"+q.RoleName+"%").Count(&total)
	}
	if q.PageSize!=""&&q.Page!="" {
		PageSize := utils.StrToInt(q.PageSize)
		Page := utils.StrToInt(q.Page)
		query = query.Limit(PageSize).Offset((Page - 1) * PageSize)
	}

	err =query.Find(&daoRole).Error
	if err==nil {
		for _, v := range daoRole {
			var role dto.SysRole
			var menu []*dao.SysMenu
			err = common.DB.Model(&v).Association("SysMenu").Find(&menu).Error
			var mu []string
			for _, m := range menu {

				mu = append(mu, strconv.Itoa(int(m.ID)))
			}
			role.ID = strconv.Itoa(v.ID)
			role.Menu = mu
			role.RoleName = v.RoleName
			role.RoleValue = v.RoleValue
			role.Remark = v.Remark
			role.OrderNo = strconv.Itoa(v.OrderNo)
			role.Status = v.Status
			role.CreatedAt = utils.TimeParseStr(v.CreatedAt, "FULL")

			roles = append(roles, &role)
		}
	}
	return
}




func SaveRole(r  dto.SysRole) (err error) {
	var role dao.SysRole
	var menu []*dao.SysMenu
	role.ID=utils.StrToInt(r.ID)
	role.Status=r.Status
	role.OrderNo=utils.StrToInt(r.OrderNo)
	role.Remark=r.Remark
	role.RoleName=r.RoleName
	role.RoleValue=r.RoleValue
	err = common.DB.Save(&role).Error
	if err!=nil{
		return
	}
	//遍历pid

	err = common.DB.Where("id in (?)",r.Menu).Find(&menu).Error

	if err!=nil{
		return
	}
	err =common.DB.Model(&role).Association("SysMenu").Replace(&menu).Error
	if err!=nil{
		return
	}

		sysCasbin,err:=GetRoleCasbinRule(&role,menu)
		//先删除casbin，再添加
		if r.ID!=""{
			common.CASBIN.RemoveFilteredPolicy(0,r.ID)
		}
		for _,casbinRule:=range sysCasbin{
			if !common.CASBIN.AddPolicy(casbinRule.RoleId,casbinRule.Path,casbinRule.Method){ //return
				continue
			}
		}

	common.CACHE.Del("sideMenu"+r.ID)
	return

}


func DelRole(roleId string) (err error) {
	var role dao.SysRole
	role.ID=utils.StrToInt(roleId)
	err =common.DB.Model(&role).Association("SysMenu").Clear().Error
	if err!=nil{
		return
	}
	err=common.DB.Delete(&role).Error
	if err!=nil{
		return
	}
	if common.CASBIN.RemoveFilteredPolicy(0,roleId){
		common.CACHE.Del("sideMenu"+roleId)
		return
	}else{
		return errors.New("casbin has no record")
	}
}

func GetRoleCasbinRule(role *dao.SysRole,roleMenuList []*dao.SysMenu)(sysCasbin []dto.SysCasbin,err error){
	var menus []*dao.SysMenu
	menusMap :=make(map[int]*dao.SysMenu)
	newMenusMap:=make( map[int]*dao.SysMenu)
	err = common.DB.Where("status = ?",1).Find(&menus).Error
	for _,menu:= range menus{
		menusMap[menu.ID]=menu
	}
	for _,m:=range roleMenuList{
		findMenuParentMap(menusMap,m,newMenusMap)
	}

		for _,mm:=range newMenusMap {
			var casbin dto.SysCasbin
			casbin.RoleId = strconv.Itoa(role.ID)
			casbin.Path = mm.ApiPath
			switch  mm.ApiMethod{
			case 1:
				casbin.Method = "GET"
			case 2:
				casbin.Method = "POST"
			case 3:
				casbin.Method = "PUT"
			case 4:
				casbin.Method = "DELETE"
			}
			sysCasbin = append(sysCasbin, casbin)

		}
		//加入通用的鉴权的path
	sysCasbin=append(sysCasbin,dto.SysCasbin{RoleId:strconv.Itoa(role.ID),Path: "/system/user/menu",Method: "GET"})
	sysCasbin=append(sysCasbin,dto.SysCasbin{RoleId:strconv.Itoa(role.ID),Path: "/system/user/info",Method: "GET"})
	//sysCasbin=append(sysCasbin,dto.SysCasbin{RoleId:strconv.Itoa(role.ID),Path: "/login",Method: "POST"})
	//sysCasbin=append(sysCasbin,dto.SysCasbin{RoleId:strconv.Itoa(role.ID),Path: "/logout",Method: "POST"})

	return
}

