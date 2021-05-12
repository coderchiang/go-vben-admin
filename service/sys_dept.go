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
func GetDeptTree(q dto.QuerySysDept) (DeptTree []*dto.SysDept, total int, err error) {
	var dept []*dao.SysDept
	query:= common.DB.Model(&dept)
	if q.Status!=""{
		query = query.Where("status = ?", q.Status)
	}
	if q.DeptName!="" {
		query = query.Where("dept_name like ?", "%"+q.DeptName+"%")
	}
	err=query.Order("order_no asc").Order("level desc").Find(&dept).Error

	DeptTree=generateDeptTree(dept)
	total=len(DeptTree)
	return
}

func SaveDept(dept *dto.SysDept) error{
	var newDept,ParentDept dao.SysDept

	//判断pid并赋值
	if dept.ParentDept!=""{
		res,ok:=strconv.Atoi(dept.ParentDept)
		if ok==nil{
			newDept.Pid=res
		}else{
			newDept.Pid=utils.StrToInt(dept.Pid)
		}
		//Level值加1
		common.DB.Where("id = ?", newDept.Pid).Find(&ParentDept)
		newDept.Level=ParentDept.Level+1
	}else{
		newDept.Pid=0
		newDept.Level=1
	}
	newDept.ID=utils.StrToInt(dept.ID)
	if newDept.Pid==newDept.ID{
		return errors.New("父子节点不能相同")
	}
	newDept.OrderNo=dept.OrderNo
	newDept.Remark=dept.Remark
	newDept.DeptName=dept.DeptName
	newDept.Status=utils.StrToInt(dept.Status)
	newDept.CreatedAt=time.Now()
	return common.DB.Model(&newDept).Save(&newDept).Error
}
func DelDept(id interface{}) error{
	if err:=common.DB.Where("pid = ?", id).Delete(&dao.SysDept{}).Error;err==nil{
		return common.DB.Where("id = ?", id).Delete(&dao.SysDept{}).Error
	}else{
		return err
	}
}




//形成树状结构
func generateDeptTree(origin []*dao.SysDept) (deptSlice []*dto.SysDept) {

	deptMap := make(map[string]*dto.SysDept)
    //重新组装输出Map数据
	for _, dept := range origin {
		var oDept =&dto.SysDept{}
		oDept.ID=strconv.Itoa(dept.ID)
		oDept.Pid=strconv.Itoa(dept.Pid)
		oDept.OrderNo=dept.OrderNo
		oDept.DeptName=dept.DeptName
		oDept.Remark=dept.Remark
		oDept.CreateTime=utils.TimeParseStr(dept.CreatedAt,"FULL")
		oDept.Status=strconv.Itoa(dept.Status)
		deptMap[oDept.ID] = oDept

	}
	//父子节点形成树结构
	for _, de := range origin {
		if de.Pid==0 {
			deptMap[strconv.Itoa(de.ID)].ParentDept="顶级目录"
		}
		if  deptMap[strconv.Itoa(de.Pid)]!=nil&&deptMap[strconv.Itoa(de.ID)]!=nil {

			deptMap[strconv.Itoa(de.ID)].ParentDept=deptMap[strconv.Itoa(de.Pid)].DeptName
			deptMap[strconv.Itoa(de.Pid)].Children = append(deptMap[strconv.Itoa(de.Pid)].Children,
				deptMap[strconv.Itoa(de.ID)])
			delete(deptMap,strconv.Itoa(de.ID))
		}
	}
	//遍历menuMap形成切片
	for _,dept := range deptMap{
		deptSlice=append(deptSlice,dept)
	}

	//menuMap排序
	orderDeptSlice(deptSlice)
	return
}
func orderDeptSlice(slice []*dto.SysDept){
	//切片根据sort值冒泡排序
	for i := 0; i < len(slice); i++ {
		if slice[i].Children!=nil{
			orderDeptSlice(slice[i].Children)
		}
		for j := i + 1; j < len(slice); j++ {
			if slice[i].OrderNo > slice[j].OrderNo {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}






