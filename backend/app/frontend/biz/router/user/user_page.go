// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "byte_go/backend/app/frontend/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/add_role", append(_addroleMw(), user.AddRole)...)
		_user.POST("/login", append(_loginMw(), user.Login)...)
		_user.POST("/logout", append(_logoutMw(), user.Logout)...)
		_user.POST("/register", append(_registerMw(), user.Register)...)
		_user.POST("/remove_role", append(_removeroleMw(), user.RemoveRole)...)
		{
			_delete := _user.Group("/delete", _deleteMw()...)
			_delete.POST("/:user_id", append(_deleteuserMw(), user.DeleteUser)...)
		}
		{
			_get := _user.Group("/get", _getMw()...)
			_get.GET("/:user_id", append(_getuserinfoMw(), user.GetUserInfo)...)
		}
		{
			_update := _user.Group("/update", _updateMw()...)
			_update.POST("/:user_id", append(_updateuserMw(), user.UpdateUser)...)
		}
	}
}
