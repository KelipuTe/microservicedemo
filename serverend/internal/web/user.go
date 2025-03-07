package web

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"microservicedemo/internal/domain"
	"microservicedemo/internal/service"
	"microservicedemo/internal/service/verifycode"
	"net/http"
)

type UserHandler struct {
	userSvc    *service.UserService
	emailVcSvc *verifycode.EmailVerifyCodeService
	smsVcSvc   *verifycode.SMSVerifyCodeService
}

func NewUserHandler(u *service.UserService, e *verifycode.EmailVerifyCodeService, s *verifycode.SMSVerifyCodeService) *UserHandler {
	return &UserHandler{
		userSvc:    u,
		emailVcSvc: e,
		smsVcSvc:   s,
	}
}

func (t *UserHandler) RegisterRoutes(server *gin.Engine) {
	group := server.Group("/user")

	group.POST("/signup", t.Signup)
	group.POST("/login", t.Login)
	group.GET("/profile", t.Profile)
	group.POST("/logout", t.Logout)

	group.POST("/signup_email", t.SignupEmail)
	group.POST("/login_email", t.LoginEmail)

	group.POST("/signup_sms", t.SignupSms)
	group.POST("/login_sms", t.LoginSms)
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
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("userId", u.Id)
	sess.Options(sessions.Options{
		MaxAge: 86400,
	})
	err = sess.Save()
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, "登录成功")
}

type ProfileResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (t *UserHandler) Profile(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	userId, _ := sess.Get("userId").(int64)
	u, err := t.userSvc.FindByUserId(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, ProfileResp{
		Id:       u.Id,
		Username: u.Username,
	})
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
		return
	}
}

type SignupEmailReq struct {
	Email string `json:"email"`
}

func (t *UserHandler) SignupEmail(ctx *gin.Context) {
	var req SignupEmailReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	err := t.emailVcSvc.Send(ctx, "SignupEmail", req.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, "验证码已发送")
}

type LoginEmailReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (t *UserHandler) LoginEmail(ctx *gin.Context) {
	var req LoginEmailReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	err := t.emailVcSvc.Verify(ctx, "SignupEmail", req.Email, req.Code)
	if err != nil {
		switch {
		case errors.Is(err, verifycode.ErrVerifyCodeNotFound):
		case errors.Is(err, verifycode.ErrWrongVerifyCode):
		case errors.Is(err, verifycode.ErrExpiredVerifyCode):
			ctx.JSON(http.StatusOK, err.Error())
		default:
			ctx.JSON(http.StatusOK, "系统错误")
		}
		return
	}

	d, err := t.userSvc.LoginEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("userId", d.Id)
	sess.Options(sessions.Options{
		MaxAge: 86400,
	})
	err = sess.Save()
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, "登录成功")

}

type SignupSmsReq struct {
	Phone string `json:"phone"`
}

// SignupSms 只发送验证码
func (t *UserHandler) SignupSms(ctx *gin.Context) {
	var req SignupSmsReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	err := t.smsVcSvc.Send(ctx, "SignupSms", req.Phone)

	switch {
	case err == nil:
		ctx.JSON(http.StatusOK, "验证码已发送")
	default:
		ctx.JSON(http.StatusOK, "系统错误")
	}
}

type LoginSmsReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// LoginSms 校验验证码、注册用户或者用户登录
func (t *UserHandler) LoginSms(ctx *gin.Context) {
	var req LoginSmsReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	//验证码
	err := t.smsVcSvc.Verify(ctx, "SignupSms", req.Phone, req.Code)

	if err != nil {
		switch {
		case errors.Is(err, verifycode.ErrVerifyCodeNotFound):
		case errors.Is(err, verifycode.ErrWrongVerifyCode):
		case errors.Is(err, verifycode.ErrExpiredVerifyCode):
			ctx.JSON(http.StatusOK, err.Error())
		default:
			ctx.JSON(http.StatusOK, "系统错误")
		}
		return
	}

	//注册或登录
	d, err := t.userSvc.LoginPhone(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("userId", d.Id)
	sess.Options(sessions.Options{
		MaxAge: 86400,
	})
	err = sess.Save()
	if err != nil {
		ctx.JSON(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, "登录成功")
}
