package casbin

import (
	"byte_go/backend/app/front/biz/dal/mysql"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strconv"
	"strings"
)

var (
	enforcer *casbin.CachedEnforcer
)

func InitCasbin() {
	tmp, _ := gormadapter.NewAdapterByDB(mysql.DB)
	m, err := model.NewModelFromFile("./casbin/model.pml")
	if err != nil {
		panic(err)
	}
	enforcer, err = casbin.NewCachedEnforcer(m, tmp)
	err = enforcer.LoadPolicy()
	if err != nil {
		return
	}
	enforcer.AddFunction("MyPolicy", KeyMatchFunc)
}

func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return (bool)(KeyMatch(name1, name2)), nil
}

func AddRoleForUser(uid uint32, role string) {
	uId := strconv.Itoa(int(uid))
	_, err := enforcer.AddRoleForUser(uId, role)
	if err != nil {
		panic(err)
	}
	_ = enforcer.SavePolicy()
}

func DeleteRolesForUser(uid uint32, role string) {
	uId := strconv.Itoa(int(uid))
	_, err := enforcer.DeleteRolesForUser(uId, role)
	if err != nil {
		panic(err)
	}
	_ = enforcer.SavePolicy()
}

func GetRolesForUser(uid uint32) []string {
	uId := strconv.Itoa(int(uid))
	roles, err := enforcer.GetRolesForUser(uId)
	if err != nil {
		panic(err)
	}
	return roles
}

func DeleteUser(uid uint32) {
	uId := strconv.Itoa(int(uid))
	_, err := enforcer.DeleteUser(uId)
	if err != nil {
		panic(err)
	}
	_ = enforcer.SavePolicy()
}

func HasRoleForUser(uid uint32, role string) bool {
	uId := strconv.Itoa(int(uid))
	has, err := enforcer.HasRoleForUser(uId, role)
	if err != nil {
		panic(err)
	}
	return has
}

func AddPolicy(uid uint32, path string, method string) {
	uId := GetRolesForUser(uid)[0]
	_, err := enforcer.AddPolicy(uId, path, method)
	if err != nil {
		panic(err)
	}
	_ = enforcer.SavePolicy()
}

func CheckPermission(uid uint32, path string, method string) bool {
	uId := GetRolesForUser(uid)[0]
	ok, err := enforcer.Enforce(uId, path, method)
	if err != nil {
		panic(err)
	}
	return ok
}
