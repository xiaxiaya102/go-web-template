package system

import (
	"go-web-template/global"
	"go-web-template/pkg/buserr"
	"go-web-template/pkg/dto"
	"go-web-template/pkg/handle"
	"go-web-template/pkg/model"
	"go-web-template/pkg/utils"
	"net/http"
	"time"

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

	username := req.UserName
	password := req.Password
	if username == "" || password == "" {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	}

	// 查询用户
	var user model.User
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	}

	// 检查用户状态
	if user.Status != 1 {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "用户已被禁用"))
		return
	}

	// 生成JWT token
	accessToken, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "生成token失败"))
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "生成刷新token失败"))
		return
	}

	// 更新用户登录信息
	user.LastLogin = time.Now()
	user.LoginCount++
	global.DB.Save(&user)

	c.JSON(http.StatusOK, handle.Success(dto.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(global.System.JWT.Expire.Seconds()),
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
	// JWT是无状态的，客户端删除token即可
	// 这里可以将token加入黑名单（如果需要的话）
	c.JSON(http.StatusOK, handle.Success(nil))
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 使用刷新token获取新的访问token
// @Tags 基础接口
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "刷新token请求参数"
// @Success 200 {object} handle.Response{data=dto.LoginResponse} "刷新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 401 {object} handle.Response "刷新token无效"
// @Router /refresh [post]
func (b *BaseApi) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 刷新token
	accessToken, refreshToken, err := utils.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "刷新token无效"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(dto.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(global.System.JWT.Expire.Seconds()),
	}))
}
