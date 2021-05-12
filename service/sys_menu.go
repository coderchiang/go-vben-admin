package service

import (
	"errors"
	"gin-vben-admin/common"
	"gin-vben-admin/common/utils"
	"gin-vben-admin/dao"
	"gin-vben-admin/dto"
	"strconv"
	"time"
)




//条件查询
func GetMenuTree(q dto.QuerySysMenu) (newmenus []*dto.SysMenu, total int, err error) {
	var menus []*dao.SysMenu
	query:=common.DB.Model(&menus)
	 if q.Status!=""{
		 query = query.Where("status = ?", q.Status)
	}
	if q.Name!=""{
		query = query.Where("cn_title like ?", "%"+q.Name+"%")
	}
		err =query.Find(&menus).Error

	newmenus=generateMenuTree(menus)
	total=len(newmenus)
	return
}

func SaveMenu(menu *dto.SysMenu) error{
	var newMenu,ParentMenu dao.SysMenu
	newMenu.ID=utils.StrToInt(menu.ID)
	newMenu.Type= utils.StrToInt(menu.MenuType)
	//判断pid并赋值

	if menu.ParentName!=""{
		res,ok:=strconv.Atoi(menu.ParentName)
		if ok==nil{
			newMenu.ParentId=res
		}else{
			newMenu.ParentId=utils.StrToInt(menu.Pid)
		}

		common.DB.Where("id = ?", newMenu.ParentId).Find(&ParentMenu)
		newMenu.Level=ParentMenu.Level+1
	}else{
		newMenu.ParentId=0
		newMenu.Level=1
	}
	if newMenu.ID!=0&&newMenu.ParentId==newMenu.ID{

		return errors.New("父子节点不能相同")
	}
	//判断菜单类型并赋值path和component
	if newMenu.Type==0{
		//path:=strings.Split(menu.ApiPath,"/")
		newMenu.Path=menu.Path
		newMenu.Component="LAYOUT"

	}else{
		//path:=strings.Split(menu.ApiPath,"/")
		newMenu.Path=menu.Path
		newMenu.Component=menu.Component
	}
	newMenu.Sort=menu.Sort
	newMenu.Status=utils.StrToInt(menu.Status)
	newMenu.CreatedAt=time.Now()
	newMenu.Icon=menu.Icon
	newMenu.ApiMethod=utils.StrToInt(menu.ApiMethod)
	newMenu.CnTitle=menu.MenuName
	newMenu.ApiPath=menu.ApiPath
	newMenu.ApiMethod=utils.StrToInt(menu.ApiMethod)
	newMenu.IsExt=utils.StrToInt(menu.IsExt)
	newMenu.Keepalive=utils.StrToInt(menu.Keepalive)

	return common.DB.Save(&newMenu).Error

}

func DelMenu(id interface{}) error{
	if err:=common.DB.Where("parent_id = ?", id).Delete(&dao.SysMenu{}).Error;err==nil{
		common.CACHE.Del("sideMenu")
		return common.DB.Where("id = ?", id).Delete(&dao.SysMenu{}).Error
	}else{
		return err
	}
}



func getSideMenuList(origin []*dao.SysMenu)(srcMenu[]*dao.SysMenu,err error){
	var menus []*dao.SysMenu
	menusMap :=make(map[int]*dao.SysMenu)
	newMenusMap:=make( map[int]*dao.SysMenu)
	err = common.DB.Where("status = ?",1).Where("type in (?)",[]int{0,1}).Find(&menus).Error
	for _,menu:= range menus{
		menusMap[menu.ID]=menu
	}
	for _,m:=range origin{
		findMenuParentMap(menusMap,m,newMenusMap)
	}
	for _,mm:=range newMenusMap{
		if mm.Type==0||mm.Type==1{
			srcMenu=append(srcMenu,mm)
		}

	}
	return
}



func findMenuParentMap(menusMap map[int]*dao.SysMenu,m *dao.SysMenu,newMenusMap map[int]*dao.SysMenu){

		newMenusMap[m.ID] =m
	if menusMap[m.ParentId]!=nil&&m.ParentId!=0{
		findMenuParentMap(menusMap,menusMap[m.ParentId],newMenusMap)
	}
	return
}


//形成树状结构
func generateMenuTree(menuList []*dao.SysMenu) (menuSlice []*dto.SysMenu) {
	func (slice  []*dao.SysMenu){
		//Level值冒泡降序
		for i := 0; i < len(slice); i++ {
			for j := i + 1; j < len(slice); j++ {
				if slice[i].Level < slice[j].Level {
					slice[i], slice[j] = slice[j], slice[i]
				}
			}
		}

	}(menuList)
	menuMap := make(map[string]*dto.SysMenu)
    //重新组装输出数据
	for _, menu := range menuList {
		var oMenu =&dto.SysMenu{}
		oMenu.ID=strconv.Itoa(menu.ID)
		oMenu.Pid=strconv.Itoa(menu.ParentId)
		oMenu.Sort=menu.Sort
		oMenu.MenuType=strconv.Itoa(menu.Type)
		oMenu.Path=menu.Path
		oMenu.Keepalive=strconv.Itoa(menu.Keepalive)
		oMenu.ApiPath=menu.ApiPath
		oMenu.ApiMethod=strconv.Itoa(menu.ApiMethod)
		oMenu.Component=menu.Component

		oMenu.CreateTime=utils.TimeParseStr(menu.CreatedAt,"FULL")
		oMenu.Status=strconv.Itoa(menu.Status)
		oMenu.Icon=menu.Icon
		if menu.Flag==0{
			oMenu.Meta.Title=menu.CnTitle
		}else{
			oMenu.Meta.Title=menu.EnTitle
		}

		oMenu.Name=oMenu.Meta.Title
		oMenu.MenuName=oMenu.Meta.Title
		oMenu.Meta.Affix=menu.Affix
		oMenu.Meta.HideTab=false
		//根据menu的类型是否显示
		oMenu.Meta.IgnoreKeepAlive=true
		oMenu.Meta.Icon=menu.Icon
		menuMap[oMenu.ID] = oMenu

	}
	//父子节点形成树结构
	for _, menu := range menuList {
		if menu.ParentId==0 {
			menuMap[strconv.Itoa(int(menu.ID))].ParentName="顶级目录"
		}
		if  menuMap[strconv.Itoa(menu.ParentId)]!=nil&&menuMap[strconv.Itoa(int(menu.ID))]!=nil {

			menuMap[strconv.Itoa(int(menu.ID))].ParentName=menuMap[strconv.Itoa(menu.ParentId)].Meta.Title
			menuMap[strconv.Itoa(menu.ParentId)].Children = append(menuMap[strconv.Itoa(menu.ParentId)].Children,
				menuMap[strconv.Itoa(int(menu.ID))])
			//Type排序后先删除Type大的节点
			delete(menuMap,strconv.Itoa(int(menu.ID)))
		}
	}
	//遍历menuMap形成切片
	for _,menu := range menuMap{
		menuSlice=append(menuSlice,menu)
	}
	orderSlice(menuSlice)

	return
}

//menuMap排序
func orderSlice(slice []*dto.SysMenu){
	//切片根据sort值冒泡排序
	for i := 0; i < len(slice); i++ {
		if slice[i].Children!=nil{
			orderSlice(slice[i].Children)
		}
		for j := i + 1; j < len(slice); j++ {
			if slice[i].Sort > slice[j].Sort {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}





