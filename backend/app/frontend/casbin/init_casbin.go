package casbin

import (
	"byte_go/backend/app/frontend/biz/dal/mysql"
	"byte_go/backend/app/frontend/conf"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"path/filepath"
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
	m, err := model.NewModelFromFile(filepath.Join(os.Getenv("PWD"), "casbin/model.pml"))
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
	_ = AddPolicyRole(UserRole, "user/register", "POST")
	_ = AddPolicyRole(UserRole, "user/login", "POST")
	_ = AddPolicyRole(UserRole, "user/logout", "POST")
	_ = AddPolicyRole(UserRole, "user/get/*", "GET")
	_ = AddPolicyRole(UserRole, "user/update", "POST")
	_ = AddPolicyRole(UserRole, "products/get/*", "GET")
	_ = AddPolicyRole(UserRole, "products/list", "GET")
	_ = AddPolicyRole(UserRole, "products/search", "GET")
	_ = AddPolicyRole(UserRole, "cart/add", "POST")
	_ = AddPolicyRole(UserRole, "cart/get", "GET")
	_ = AddPolicyRole(UserRole, "cart/empty", "POST")
	_ = AddPolicyRole(UserRole, "cart/place", "POST")
	_ = AddPolicyRole(UserRole, "cart/list", "GET")
	_ = AddPolicyRole(UserRole, "cart/mark_paid", "POST")
	_ = AddPolicyRole(UserRole, "payment/charge", "POST")
	_ = AddPolicyRole(UserRole, "checkout", "POST")
	_ = AddPolicyRole(AdminRole, "user/delete/*", "POST")
	_ = AddPolicyRole(AdminRole, "user/add_role", "POST")
	_ = AddPolicyRole(AdminRole, "user/remove_role", "POST")
	_ = AddPolicyRole(MerchantRole, "products/create", "POST")
	_ = AddPolicyRole(MerchantRole, "products/delete/*", "GET")
	_ = AddPolicyRole(MerchantRole, "products/update/*", "POST")
	if conf.GetEnv() == "test" {
		_ = AddRoleForUser(1, AdminRole)
		_ = AddRoleForUser(2, UserRole)
		_ = AddRoleForUser(2, MerchantRole)
		_ = AddRoleForUser(3, UserRole)
	}
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

func checkRole(role string) error {
	if role == UserRole || role == AdminRole || role == MerchantRole {
		return nil
	}
	return errors.New("role is not valid")
}

func AddPolicyRole(role string, path string, method string) (err error) {
	if err = checkRole(role); err != nil {
		hlog.Errorf("role %s is not valid, err=%v", role, err)
		return err
	}
	_, err = enforcer.AddPolicy(role, path, method)
	if err != nil {
		hlog.Errorf("add policy %s %s %s failed, err=%v", role, path, method, err)
		return err
	}
	return nil
}

func AddRoleForUser(uid uint32, role string) (err error) {
	if err = checkRole(role); err != nil {
		hlog.Errorf("role %s is not valid, err=%v", role, err)
		return err
	}
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.AddRoleForUser(uId, role)
	if err != nil {
		hlog.Errorf("add role %s for user %s failed, err=%v", role, uId, err)
		return err
	}
	err = enforcer.SavePolicy()
	if err != nil {
		hlog.Errorf("save policy failed, err=%v", err)
		return err
	}
	return nil
}

func DeleteRoleForUser(uid uint32, role string) (err error) {
	if err = checkRole(role); err != nil {
		hlog.Errorf("role %s is not valid, err=%v", role, err)
		return err
	}
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteRoleForUser(uId, role)
	if err != nil {
		hlog.Errorf("delete role %s for user %s failed, err=%v", role, uId, err)
		return err
	}
	err = enforcer.SavePolicy()
	if err != nil {
		hlog.Errorf("save policy failed, err=%v", err)
		return err
	}
	return nil
}

func DeleteRolesForUser(uid uint32) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteRolesForUser(uId)
	if err != nil {
		hlog.Errorf("delete roles for user %s failed, err=%v", uId, err)
		return err
	}
	err = enforcer.SavePolicy()
	if err != nil {
		hlog.Errorf("save policy failed, err=%v", err)
		return err
	}
	return nil
}

func GetRolesForUser(uid uint32) ([]string, error) {
	uId := strconv.Itoa(int(uid))
	roles, err := enforcer.GetRolesForUser(uId)
	if err != nil {
		hlog.Errorf("get roles for user %s failed, err=%v", uId, err)
		return nil, err
	}
	return roles, nil
}

func DeleteUser(uid uint32) (err error) {
	uId := strconv.Itoa(int(uid))
	_, err = enforcer.DeleteUser(uId)
	if err != nil {
		hlog.Errorf("delete user %s failed, err=%v", uId, err)
		return err
	}
	err = enforcer.SavePolicy()
	if err != nil {
		hlog.Errorf("save policy failed, err=%v", err)
		return err
	}
	return nil
}

func HasRoleForUser(uid uint32, role string) (bool, error) {
	uId := strconv.Itoa(int(uid))
	has, err := enforcer.HasRoleForUser(uId, role)
	if err != nil {
		hlog.Errorf("has role %s for user %s failed, err=%v", role, uId, err)
		return false, err
	}
	return has, nil
}

func CheckPermission(uid uint32, path string, method string) (bool, error) {
	uIds, err := GetRolesForUser(uid)
	if err != nil {
		hlog.Errorf("get roles for user %d failed, err=%v", uid, err)
		return false, err
	}
	for _, uId := range uIds {
		ok, err := enforcer.Enforce(uId, path, method)
		if err != nil {
			hlog.Errorf("enforce failed, err=%v", err)
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}
