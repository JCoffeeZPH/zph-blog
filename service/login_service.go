package service

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"zph/dao"
	"zph/lib/cache"
	"zph/logger"
	"zph/models/bo"
	"zph/models/request"
	"zph/models/response"
	"zph/models/util"
	"zph/session"
	"zph/utils"
)

var log = logger.Log

type UserService interface {
	VerifyToken(ctx *gin.Context, req *request.UserLoginRequest, jwtToken string) *response.LoginUserResponse
	GetUserByUsername(username string) *bo.User
}

type UserServiceImpl struct {
	userDao dao.UserDao
	loginLogService LoginLogService
}

func NewUserService() UserService {
	return &UserServiceImpl{
		userDao: dao.NewUserDao(),
		loginLogService: NewLoginLogService(),
	}
}

func (service *UserServiceImpl) VerifyToken(ctx *gin.Context, req *request.UserLoginRequest, jwtToken string) *response.LoginUserResponse {
	var session *bo.Session
	if jwtToken != "" {
		s := util.ParseSession(jwtToken)
		if s != nil && s.Username == req.Username{
			session = s
		}
	}

	if session == nil {
		user := service.userDao.GetUser(req.Username, req.Password)
		if user == nil {
			log.Errorf("VerifyToken not found user")
			go service.loginLogService.CreateNewLoginLog(ctx, false)
			return nil
		}
		sessionId := uuid.NewV1().String()
		secret := uuid.NewV1().String()
		s := &bo.Session{
			SessionId: sessionId,
			UserId: user.Id,
			Secret: secret,
			Username: user.Username,
			Nickname: user.Nickname,
			Email: user.Email,
			Avatar: user.Avatar,
			Role: user.Role,
			UpdateTime: user.UpdateTime,
			CreateTime: user.CreateTime,
		}
		session = s
	}

	go cache.SaveSessionIntoRedis(session)
	go service.loginLogService.CreateNewLoginLog(ctx, true)
	service.setContext(ctx, session)
	return service.createLoginResp(session)
}

func (service *UserServiceImpl) createLoginResp(session *bo.Session) *response.LoginUserResponse {
	return &response.LoginUserResponse{
		UserId:     session.UserId,
		Username:   session.Username,
		Nickname:   session.Nickname,
		Avatar:     session.Avatar,
		Email:      session.Email,
		Token:      util.GenerateJwtToken(session.SessionId, session.Secret),
		Role:       session.Role,
		UpdateTime: utils.TimeFormat(session.UpdateTime),
		CreateTime: utils.TimeFormat(session.CreateTime),
	}
}

func (service *UserServiceImpl) setContext(ctx *gin.Context, s *bo.Session)  {
	session.SetUserId(ctx, s.UserId)
	session.SetUsername(ctx, s.Username)
	session.SetNickname(ctx, s.Nickname)
	session.SetEmail(ctx, s.Email)
	session.SetAvatar(ctx, s.Avatar)
	session.SetRole(ctx, s.Role)
}

func (service *UserServiceImpl) GetUserByUsername(username string) *bo.User{
	user := service.userDao.GetUserByUsername(username)
	return &bo.User{
		Id: user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar: user.Avatar,
		Email: user.Email,
		CreateTime: utils.TimeFormat(user.CreateTime),
		UpdateTime: utils.TimeFormat(user.UpdateTime),
		Role: user.Role,
	}
}
