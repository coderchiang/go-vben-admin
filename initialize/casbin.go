package initialize

import (
	"gin-vben-admin/common"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"strings"
)

func InitCasbin()  {
	a := gormadapter.NewAdapterByDB(common.DB)
	e := casbin.NewEnforcer(common.CONFIG.Casbin.ModelPath, a)
	e.EnableLog(true)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)

	common.CASBIN=e
	return
}


func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}


func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}