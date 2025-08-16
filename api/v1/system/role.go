package system

import (
	"go-web-template/global"
	"go-web-template/pkg/dto"
	"go-web-template/pkg/handle"
	"go-web-template/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

// CreateRole 创建角色
// @Summary 创建角色
// @Description 创建新角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param request body dto.CreateRoleRequest true "创建角色请求参数"
// @Success 200 {object} handle.Response{data=dto.RoleResponse} "创建成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Security BearerAuth
// @Router /role [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var req dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查角色名称和编码是否已存在
	var count int64
	global.DB.Model(&model.Role{}).Where("name = ? OR code = ?", req.Name, req.Code).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色名称或编码已存在"))
		return
	}

	// 创建角色
	role := &model.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
		Sort:        req.Sort,
	}

	if err := global.DB.Create(role).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "创建角色失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToRoleResponse(role)))
}

// GetRoleList 获取角色列表
// @Summary 获取角色列表
// @Description 获取所有角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]dto.RoleResponse} "获取成功"
// @Security BearerAuth
// @Router /role [get]
func (r *RoleApi) GetRoleList(c *gin.Context) {
	var roles []model.Role
	if err := global.DB.Order("sort ASC, id ASC").Find(&roles).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "查询角色列表失败"))
		return
	}

	var roleResponses []dto.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, convertToRoleResponse(&role))
	}

	c.JSON(http.StatusOK, handle.Success(roleResponses))
}

// GetRole 获取角色详情
// @Summary 获取角色详情
// @Description 根据角色ID获取角色详情
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} handle.Response{data=dto.RoleResponse} "获取成功"
// @Failure 404 {object} handle.Response "角色不存在"
// @Security BearerAuth
// @Router /role/{id} [get]
func (r *RoleApi) GetRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色ID格式错误"))
		return
	}

	var role model.Role
	if err := global.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色不存在"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToRoleResponse(&role)))
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Description 更新角色信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param request body dto.UpdateRoleRequest true "更新角色请求参数"
// @Success 200 {object} handle.Response{data=dto.RoleResponse} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 404 {object} handle.Response "角色不存在"
// @Security BearerAuth
// @Router /role/{id} [put]
func (r *RoleApi) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色ID格式错误"))
		return
	}

	var req dto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 查找角色
	var role model.Role
	if err := global.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色不存在"))
		return
	}

	// 检查角色名称和编码是否已被其他角色使用
	var count int64
	global.DB.Model(&model.Role{}).Where("(name = ? OR code = ?) AND id != ?", req.Name, req.Code, id).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色名称或编码已被其他角色使用"))
		return
	}

	// 更新角色信息
	role.Name = req.Name
	role.Code = req.Code
	role.Description = req.Description
	role.Status = req.Status
	role.Sort = req.Sort

	if err := global.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新角色失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToRoleResponse(&role)))
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} handle.Response "删除成功"
// @Failure 404 {object} handle.Response "角色不存在"
// @Security BearerAuth
// @Router /role/{id} [delete]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色ID格式错误"))
		return
	}

	// 检查角色是否存在
	var role model.Role
	if err := global.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "角色不存在"))
		return
	}

	// 检查是否有用户使用该角色
	var userRoleCount int64
	global.DB.Model(&model.UserRole{}).Where("role_id = ?", id).Count(&userRoleCount)
	if userRoleCount > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该角色下还有用户，无法删除"))
		return
	}

	// 删除角色权限关联
	global.DB.Where("role_id = ?", id).Delete(&model.RolePermission{})

	// 删除角色
	if err := global.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除角色失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// convertToRoleResponse 转换为角色响应格式
func convertToRoleResponse(role *model.Role) dto.RoleResponse {
	return dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      role.Status,
		Sort:        role.Sort,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}
