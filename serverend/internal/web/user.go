package web

import (
	"demo-golang/microservice/internal/domain"
	"demo-golang/microservice/internal/service"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	group.POST("/profile", t.Profile)
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

	u := domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := t.userSvc.Login(ctx, u)
	switch {
	case err == nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 86400,
		})
		err = sess.Save()
		ctx.JSON(http.StatusOK, "登录成功")
	case errors.Is(err, service.ErrInvalidUserOrPassword):
		ctx.JSON(http.StatusOK, err.Error())
	default:
		ctx.JSON(http.StatusOK, "系统错误")
	}
}

func (t *UserHandler) Profile(ctx *gin.Context) {

}

func (t *UserHandler) Logout(ctx *gin.Context) {}
