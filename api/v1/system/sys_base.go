package system

import (
	"feishuReboot/pkg/buserr"
	"feishuReboot/pkg/dto"
	"feishuReboot/pkg/handle"
	"feishuReboot/pkg/repo"
	"feishuReboot/pkg/service"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 基础接口
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "登录请求参数"
// @Success 200 {object} handle.Response{data=dto.LoginResponse} "登录成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 401 {object} handle.Response "用户名或密码错误"
// @Router /login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	userName := req.UserName
	password := req.Password
	if userName == "" || password == "" {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	}

	user, _ := repo.QueryUserWithName(userName)
	var token string
	if user == nil || user.Password != password { // 验证密码
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	} else {
		now := time.Now()
		if now.After(user.ExpireTime) {
			token = strings.ReplaceAll(uuid.New().String(), "-", "")
			user.Token = token
			user.LoginTime = now
			user.ExpireTime = now.Add(time.Hour * 2)
			repo.UpdateUser(user)
		} else {
			token = user.Token
		}
	}
	service.SetUser(token, user)

	c.JSON(http.StatusOK, handle.Success(dto.LoginResponse{
		Token: token,
	}))
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出接口
// @Tags 基础接口
// @Accept json
// @Produce json
// @Param request body dto.LogoutRequest true "登出请求参数"
// @Success 200 {object} handle.Response "登出成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 401 {object} handle.Response "未授权"
// @Security BearerAuth
// @Router /logout [post]
func (b *BaseApi) Logout(c *gin.Context) {
	var req dto.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	user, _ := repo.QueryUserWithName("admin")
	user.Token = ""
	user.ExpireTime = time.Now()
	repo.UpdateUser(user)
	service.RemoveUser(req.Token)
	c.JSON(http.StatusOK, handle.Success(nil))

	// if req.Token != "" {
	// 	user, err := repo.QueryUserWithToken(req.Token)
	// 	if err == nil && user != nil {
	// 		user.Token = ""
	// 		user.ExpireTime = time.Now()
	// 		repo.UpdateUser(user)
	// 		service.RemoveUser(req.Token)
	// 		c.JSON(http.StatusOK, handle.Success(nil))

	// 	} else {
	// 		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的token"))
	// 		return
	// 	}
	// } else {
	// 	c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的token"))
	// }
}
