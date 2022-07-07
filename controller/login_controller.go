package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zph/constants"
	api_error "zph/error"
	"zph/logger"
	"zph/models/request"
	"zph/service"
)

var log = logger.Log

type UserController struct {
	BaseController
	userService service.UserService
	loginLogService service.LoginLogService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
		loginLogService: service.NewLoginLogService(),
	}
}

func (c *UserController)getJwtToken(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	splits := strings.Split(header, " ")
	if len(splits) == 2 && strings.ToLower(splits[0]) == "bearer" {
		return splits[1]
	} else {
		return ""
	}
}

func (c *UserController)LoginController(ctx *gin.Context){
	jwtToken := c.getJwtToken(ctx)
	req, err := parseLoginParams(ctx)
	if err != nil {
		go c.loginLogService.CreateNewLoginLog(ctx, false)
		panic(api_error.ParamError)
	}
	resp := c.userService.VerifyToken(ctx, req, jwtToken)
	ctx.Set(constants.OperationKey, constants.LoginOperation.Value())
	c.Success(ctx, resp)
}

func parseLoginParams(c *gin.Context) (*request.UserLoginRequest, error) {
	req := &request.UserLoginRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		log.Errorf("parseLoginParams failed, err is: %+v", err)
		return nil, err
	}
	return req, nil
}

