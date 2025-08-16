package system

import (
	"go-web-template/global"
	"go-web-template/pkg/dto"
	"go-web-template/pkg/handle"
	"go-web-template/pkg/model"
	"go-web-template/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "创建用户请求参数"
// @Success 200 {object} handle.Response{data=dto.UserResponse} "创建成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Security BearerAuth
// @Router /user [post]
func (u *UserApi) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查用户名是否已存在
	var count int64
	global.DB.Model(&model.User{}).Where("username = ? OR user_id = ?", req.Username, req.UserID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户名或用户ID已存在"))
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "密码加密失败"))
		return
	}

	// 创建用户
	user := &model.User{
		UserID:   req.UserID,
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
		Remark:   req.Remark,
	}

	if err := global.DB.Create(user).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "创建用户失败"))
		return
	}

	// 分配角色
	if len(req.RoleIDs) > 0 {
		for _, roleID := range req.RoleIDs {
			userRole := &model.UserRole{
				UserID: user.ID,
				RoleID: roleID,
			}
			global.DB.Create(userRole)
		}
	}

	c.JSON(http.StatusOK, handle.Success(convertToUserResponse(user)))
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param username query string false "用户名（模糊搜索）"
// @Param status query int false "状态筛选"
// @Success 200 {object} handle.Response{data=dto.UserListResponse} "获取成功"
// @Security BearerAuth
// @Router /user [get]
func (u *UserApi) GetUserList(c *gin.Context) {
	var req dto.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询
	query := global.DB.Model(&model.User{})

	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}

	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 获取总数
	var total int64
	query.Count(&total)

	// 分页查询
	var users []model.User
	offset := (req.Page - 1) * req.PageSize
	if err := query.Offset(offset).Limit(req.PageSize).Find(&users).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "查询用户列表失败"))
		return
	}

	// 转换响应
	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, convertToUserResponse(&user))
	}

	response := dto.UserListResponse{
		List:  userResponses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据用户ID获取用户详情
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} handle.Response{data=dto.UserResponse} "获取成功"
// @Failure 404 {object} handle.Response "用户不存在"
// @Security BearerAuth
// @Router /user/{id} [get]
func (u *UserApi) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户ID格式错误"))
		return
	}

	var user model.User
	if err := global.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户不存在"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToUserResponse(&user)))
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param request body dto.UpdateUserRequest true "更新用户请求参数"
// @Success 200 {object} handle.Response{data=dto.UserResponse} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 404 {object} handle.Response "用户不存在"
// @Security BearerAuth
// @Router /user/{id} [put]
func (u *UserApi) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户ID格式错误"))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 查找用户
	var user model.User
	if err := global.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户不存在"))
		return
	}

	// 检查用户名是否已被其他用户使用
	var count int64
	global.DB.Model(&model.User{}).Where("(username = ? OR user_id = ?) AND id != ?", req.Username, req.UserID, id).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户名或用户ID已被其他用户使用"))
		return
	}

	// 更新用户信息
	user.UserID = req.UserID
	user.Username = req.Username
	user.Email = req.Email
	user.Phone = req.Phone
	user.Avatar = req.Avatar
	user.Status = req.Status
	user.Remark = req.Remark

	if err := global.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新用户失败"))
		return
	}

	// 更新用户角色
	// 先删除原有角色关联
	global.DB.Where("user_id = ?", user.ID).Delete(&model.UserRole{})

	// 添加新的角色关联
	if len(req.RoleIDs) > 0 {
		for _, roleID := range req.RoleIDs {
			userRole := &model.UserRole{
				UserID: user.ID,
				RoleID: roleID,
			}
			global.DB.Create(userRole)
		}
	}

	c.JSON(http.StatusOK, handle.Success(convertToUserResponse(&user)))
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} handle.Response "删除成功"
// @Failure 404 {object} handle.Response "用户不存在"
// @Security BearerAuth
// @Router /user/{id} [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户ID格式错误"))
		return
	}

	// 检查用户是否存在
	var user model.User
	if err := global.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户不存在"))
		return
	}

	// 删除用户角色关联
	global.DB.Where("user_id = ?", id).Delete(&model.UserRole{})

	// 删除用户
	if err := global.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除用户失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前用户密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body dto.ChangePasswordRequest true "修改密码请求参数"
// @Success 200 {object} handle.Response "修改成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Security BearerAuth
// @Router /user/change-password [post]
func (u *UserApi) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户未登录"))
		return
	}

	// 查找用户
	var user model.User
	if err := global.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户不存在"))
		return
	}

	// 验证旧密码
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "旧密码错误"))
		return
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "密码加密失败"))
		return
	}

	// 更新密码
	user.Password = hashedPassword
	if err := global.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "修改密码失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// convertToUserResponse 转换为用户响应格式
func convertToUserResponse(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:         user.ID,
		UserID:     user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		Avatar:     user.Avatar,
		Status:     user.Status,
		LastLogin:  user.LastLogin,
		LoginCount: user.LoginCount,
		Remark:     user.Remark,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		// TODO: 添加角色信息
		Roles: []string{},
	}
}
