package user

import (
	"byte_go/backend/app/frontend/casbin"
	"byte_go/backend/app/frontend/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"byte_go/backend/app/frontend/hertz_gen/frontend/common_hertz"
	user "byte_go/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user.DeleteUserReq) (resp *common_hertz.Empty, err error) {

	_, err = rpc.UserClient.DeleteUser(h.Context, &rpcUser.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "delete user [%d] failed, err: %v", req.UserId, err.Error())
		return nil, err
	}

	if err = casbin.DeleteRolesForUser(req.UserId); err != nil {
		hlog.CtxErrorf(h.Context, "delete roles for user [%d] failed, err: %v", req.UserId, err.Error())
		return nil, err
	}

	return
}
