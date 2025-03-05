package web

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"microservicedemo/internal/domain"
	"microservicedemo/internal/service"
	"net/http"
)

type UserHandler struct {
	userSvc *service.UserService
}

func NewUserHandler(userSvc *service.UserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (t *UserHandler) RegisterRoutes(server *gin.Engine) {
	group := server.Group("/user")
	group.POST("/signup", t.Signup)
	group.POST("/login", t.Login)
	group.POST("/signup_email", t.SignupEmail)
	group.POST("/login_email", t.LoginEmail)
	group.POST("/signup_sms", t.SignupSms)
	group.POST("/login_sms", t.LoginSms)
	group.GET("/profile", t.Profile)
	group.POST("/logout", t.Logout)
}

type SignupReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (t *UserHandler) Signup(ctx *gin.Context) {
	var req SignupReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "密码和确认密码不一致")
		return
	}

	u := domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := t.userSvc.Signup(ctx, u)
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, "注册成功")
	case errors.Is(err, service.ErrUserDuplicate):
		ctx.JSON(http.StatusOK, err.Error())
	default:
		ctx.JSON(http.StatusOK, "系统错误")
	}
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (t *UserHandler) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	u, err := t.userSvc.Login(ctx, req.Username, req.Password)
	switch {
	case err == nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 86400,
		})
		err = sess.Save()
		if err != nil {
			ctx.JSON(http.StatusOK, "系统错误")
		}
		ctx.JSON(http.StatusOK, "登录成功")
	case errors.Is(err, service.ErrWrongPassword):
		ctx.JSON(http.StatusOK, err.Error())
	default:
		ctx.JSON(http.StatusOK, "系统错误")
	}
}

func (t *UserHandler) SignupEmail(ctx *gin.Context) {

}
func (t *UserHandler) LoginEmail(ctx *gin.Context) {

}

func (t *UserHandler) SignupSms(ctx *gin.Context) {

}
func (t *UserHandler) LoginSms(c *gin.Context) {

}

type ProfileResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (t *UserHandler) Profile(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	userId, _ := sess.Get("userId").(int64)
	u, err := t.userSvc.FindByUserId(ctx, userId)
	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, ProfileResp{
			Id:       u.Id,
			Username: u.Username,
		})
	default:
		ctx.JSON(http.StatusOK, "系统错误")
	}
}

func (t *UserHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	err := session.Save()
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
	}
}
