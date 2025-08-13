package system

import (
	"feishuReboot/pkg/dto"
	"feishuReboot/pkg/handle"
	"feishuReboot/pkg/model"
	"feishuReboot/pkg/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaConfigApi struct{}

// SaveMediaConfig 保存媒体配置
// @Summary 保存媒体配置
// @Description 保存新的媒体配置信息
// @Tags 媒体配置
// @Accept json
// @Produce json
// @Param request body dto.SaveMediaConfigRequest true "媒体配置信息"
// @Success 200 {object} handle.Response{data=model.MediaConfig} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 401 {object} handle.Response "未授权"
// @Security BearerAuth
// @Router /mainModule/saveMediaConfig [post]
func (m *MediaConfigApi) SaveMediaConfig(c *gin.Context) {
	var req dto.SaveMediaConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	mediaConfig := &model.MediaConfig{
		Name:      req.Name,
		Url:       req.Url,
		Status:    req.Status,
		Code:      req.Code,
		Type:      req.Type,
		MediaDesc: req.MediaDesc,
	}

	if err := repo.SaveMediaConfig(mediaConfig); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(mediaConfig))
}

// UpdateMediaConfig 更新媒体配置
// @Summary 更新媒体配置
// @Description 更新已有的媒体配置信息
// @Tags 媒体配置
// @Accept json
// @Produce json
// @Param request body dto.UpdateMediaConfigRequest true "媒体配置更新信息"
// @Success 200 {object} handle.Response{data=model.MediaConfig} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Failure 401 {object} handle.Response "未授权"
// @Security BearerAuth
// @Router /mainModule/updateMediaConfig [post]
func (m *MediaConfigApi) UpdateMediaConfig(c *gin.Context) {
	var req dto.UpdateMediaConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	mediaConfig := &model.MediaConfig{
		ID:        req.ID,
		Name:      req.Name,
		Url:       req.Url,
		Status:    req.Status,
		Code:      req.Code,
		Type:      req.Type,
		MediaDesc: req.MediaDesc,
	}

	if err := repo.UpdateMediaConfig(mediaConfig); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(mediaConfig))
}

// DelMediaConfig 删除媒体配置
// @Summary 删除媒体配置
// @Description 根据ID删除媒体配置
// @Tags 媒体配置
// @Accept json
// @Produce json
// @Param id path int true "媒体配置ID"
// @Success 200 {object} handle.Response "删除成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /mainModule/delMediaConfig/{id} [delete]
func (m *MediaConfigApi) DelMediaConfig(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteMediaConfig(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetMediaConfig 获取媒体配置详情
// @Summary 获取媒体配置详情
// @Description 根据ID获取媒体配置的详细信息
// @Tags 媒体配置
// @Accept json
// @Produce json
// @Param id path int true "媒体配置ID"
// @Success 200 {object} handle.Response{data=model.MediaConfig} "获取成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /mainModule/getMediaConfig/{id} [get]
func (m *MediaConfigApi) GetMediaConfig(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	mediaConfig, err := repo.GetMediaConfigByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(mediaConfig))
}

// GetMediaConfigPage 分页获取媒体配置
// @Summary 分页获取媒体配置
// @Description 分页获取媒体配置列表
// @Tags 媒体配置
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} handle.Response{data=dto.MediaConfigPageResponse} "获取成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /mainModule/getMediaConfigPage [get]
func (m *MediaConfigApi) GetMediaConfigPage(c *gin.Context) {
	var req dto.MediaConfigPageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := repo.GetMediaConfigPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.MediaConfigPageResponse{
		List:  make([]dto.MediaConfigResponse, len(list)),
		Total: total,
	}

	for i, item := range list {
		response.List[i] = dto.MediaConfigResponse{
			ID:        item.ID,
			Name:      item.Name,
			Url:       item.Url,
			Status:    item.Status,
			Code:      item.Code,
			Type:      item.Type,
			MediaDesc: item.MediaDesc,
		}
	}

	c.JSON(http.StatusOK, handle.Success(response))
}
