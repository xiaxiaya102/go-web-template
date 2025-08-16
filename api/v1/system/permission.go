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

type PermissionApi struct{}

// CreatePermission 创建权限
// @Summary 创建权限
// @Description 创建新权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param request body dto.CreatePermissionRequest true "创建权限请求参数"
// @Success 200 {object} handle.Response{data=dto.PermissionResponse} "创建成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Security BearerAuth
// @Router /permission [post]
func (p *PermissionApi) CreatePermission(c *gin.Context) {
	var req dto.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查权限名称和编码是否已存在
	var count int64
	global.DB.Model(&model.Permission{}).Where("name = ? OR code = ?", req.Name, req.Code).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限名称或编码已存在"))
		return
	}

	// 创建权限
	permission := &model.Permission{
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		ParentID:    req.ParentID,
		Path:        req.Path,
		Component:   req.Component,
		Icon:        req.Icon,
		Sort:        req.Sort,
		Status:      req.Status,
		Description: req.Description,
	}

	if err := global.DB.Create(permission).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "创建权限失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToPermissionResponse(permission)))
}

// GetPermissionList 获取权限列表
// @Summary 获取权限列表
// @Description 获取权限树形列表
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]dto.PermissionResponse} "获取成功"
// @Security BearerAuth
// @Router /permission [get]
func (p *PermissionApi) GetPermissionList(c *gin.Context) {
	var permissions []model.Permission
	if err := global.DB.Order("sort ASC, id ASC").Find(&permissions).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "查询权限列表失败"))
		return
	}

	// 构建树形结构
	permissionTree := buildPermissionTree(permissions, 0)
	c.JSON(http.StatusOK, handle.Success(permissionTree))
}

// GetPermission 获取权限详情
// @Summary 获取权限详情
// @Description 根据权限ID获取权限详情
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path int true "权限ID"
// @Success 200 {object} handle.Response{data=dto.PermissionResponse} "获取成功"
// @Failure 404 {object} handle.Response "权限不存在"
// @Security BearerAuth
// @Router /permission/{id} [get]
func (p *PermissionApi) GetPermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限ID格式错误"))
		return
	}

	var permission model.Permission
	if err := global.DB.First(&permission, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限不存在"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToPermissionResponse(&permission)))
}

// UpdatePermission 更新权限
// @Summary 更新权限
// @Description 更新权限信息
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path int true "权限ID"
// @Param request body dto.CreatePermissionRequest true "更新权限请求参数"
// @Success 200 {object} handle.Response{data=dto.PermissionResponse} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 404 {object} handle.Response "权限不存在"
// @Security BearerAuth
// @Router /permission/{id} [put]
func (p *PermissionApi) UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限ID格式错误"))
		return
	}

	var req dto.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 查找权限
	var permission model.Permission
	if err := global.DB.First(&permission, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限不存在"))
		return
	}

	// 检查权限名称和编码是否已被其他权限使用
	var count int64
	global.DB.Model(&model.Permission{}).Where("(name = ? OR code = ?) AND id != ?", req.Name, req.Code, id).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限名称或编码已被其他权限使用"))
		return
	}

	// 更新权限信息
	permission.Name = req.Name
	permission.Code = req.Code
	permission.Type = req.Type
	permission.ParentID = req.ParentID
	permission.Path = req.Path
	permission.Component = req.Component
	permission.Icon = req.Icon
	permission.Sort = req.Sort
	permission.Status = req.Status
	permission.Description = req.Description

	if err := global.DB.Save(&permission).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新权限失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(convertToPermissionResponse(&permission)))
}

// DeletePermission 删除权限
// @Summary 删除权限
// @Description 删除权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param id path int true "权限ID"
// @Success 200 {object} handle.Response "删除成功"
// @Failure 404 {object} handle.Response "权限不存在"
// @Security BearerAuth
// @Router /permission/{id} [delete]
func (p *PermissionApi) DeletePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限ID格式错误"))
		return
	}

	// 检查权限是否存在
	var permission model.Permission
	if err := global.DB.First(&permission, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "权限不存在"))
		return
	}

	// 检查是否有子权限
	var childCount int64
	global.DB.Model(&model.Permission{}).Where("parent_id = ?", id).Count(&childCount)
	if childCount > 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该权限下还有子权限，无法删除"))
		return
	}

	// 删除角色权限关联
	global.DB.Where("permission_id = ?", id).Delete(&model.RolePermission{})

	// 删除权限
	if err := global.DB.Delete(&permission).Error; err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除权限失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// buildPermissionTree 构建权限树
func buildPermissionTree(permissions []model.Permission, parentID uint) []dto.PermissionResponse {
	var tree []dto.PermissionResponse

	for _, permission := range permissions {
		if permission.ParentID == parentID {
			permissionResp := convertToPermissionResponse(&permission)
			permissionResp.Children = buildPermissionTree(permissions, permission.ID)
			tree = append(tree, permissionResp)
		}
	}

	return tree
}

// convertToPermissionResponse 转换为权限响应格式
func convertToPermissionResponse(permission *model.Permission) dto.PermissionResponse {
	return dto.PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Code:        permission.Code,
		Type:        permission.Type,
		ParentID:    permission.ParentID,
		Path:        permission.Path,
		Component:   permission.Component,
		Icon:        permission.Icon,
		Sort:        permission.Sort,
		Status:      permission.Status,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}
