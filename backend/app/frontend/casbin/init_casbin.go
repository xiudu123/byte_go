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

const (
	UserRole     = "user"
	AdminRole    = "admin"
	MerchantRole = "merchant"
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
	//enforcer.AddFunction("MyPolicy", KeyMatchFunc)
	if mysql.NeedDate {
		initPagePermission()
	}
}

func initPagePermission() {
	_ = AddPolicyRole(UserRole, "user/register", "post")
	_ = AddPolicyRole(UserRole, "user/login", "post")
	_ = AddPolicyRole(UserRole, "user/logout", "post")
	_ = AddPolicyRole(UserRole, "user/get/*", "get")
	_ = AddPolicyRole(UserRole, "user/update", "post")
	_ = AddPolicyRole(UserRole, "user/delete/*", "post")
	_ = AddPolicyRole(UserRole, "product/get/*", "get")
	_ = AddPolicyRole(UserRole, "product/list", "get")
	_ = AddPolicyRole(UserRole, "product/search", "get")
	_ = AddPolicyRole(UserRole, "cart/add", "post")
	_ = AddPolicyRole(UserRole, "cart/get", "get")
	_ = AddPolicyRole(UserRole, "cart/empty", "post")
	_ = AddPolicyRole(UserRole, "cart/place", "post")
	_ = AddPolicyRole(UserRole, "cart/list", "get")
	_ = AddPolicyRole(UserRole, "cart/mark_paid", "post")
	_ = AddPolicyRole(UserRole, "payment/charge", "post")
	_ = AddPolicyRole(UserRole, "checkout", "post")
	_ = AddPolicyRole(AdminRole, "user/delete", "post")
	_ = AddPolicyRole(MerchantRole, "product/add", "post")
	_ = AddPolicyRole(MerchantRole, "product/delete", "get")
	_ = AddPolicyRole(MerchantRole, "product/update", "post")
}

func AddPolicyRole(role string, path string, method string) (err error) {
	_, err = enforcer.AddPolicy(role, path, method)
	return err
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

func AddRoleForUser(uid uint32, role string) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.AddRoleForUser(uId, role)
	if err != nil {
		return err
	}
	err = enforcer.SavePolicy()
	return err
}

func DeleteRoleForUser(uid uint32, role string) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteRoleForUser(uId, role)
	if err != nil {
		return err
	}
	err = enforcer.SavePolicy()
	return err
}

func DeleteRolesForUser(uid uint32) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteRolesForUser(uId)
	if err != nil {
		return err
	}
	err = enforcer.SavePolicy()
	return err
}

func GetRolesForUser(uid uint32) ([]string, error) {
	uId := strconv.Itoa(int(uid))
	roles, err := enforcer.GetRolesForUser(uId)
	return roles, err
}

func DeleteUser(uid uint32) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteUser(uId)
	if err != nil {
		return err
	}
	err = enforcer.SavePolicy()
	return err
}

func HasRoleForUser(uid uint32, role string) (bool, error) {
	uId := strconv.Itoa(int(uid))
	has, err := enforcer.HasRoleForUser(uId, role)
	return has, err
}

func CheckPermission(uid uint32, path string, method string) (bool, error) {
	uIds, err := GetRolesForUser(uid)
	if err != nil {
		return false, err
	}
	for _, uId := range uIds {
		ok, err := enforcer.Enforce(uId, path, method)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}
