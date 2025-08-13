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

type ThresholdConfApi struct{}

// SaveThresholdConf 保存阈值配置
// @Summary 保存阈值配置
// @Description 保存新的阈值配置信息
// @Tags 阈值配置
// @Accept json
// @Produce json
// @Param request body dto.SaveThresholdConfRequest true "阈值配置信息"
// @Success 200 {object} handle.Response{data=model.ThresholdConf} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /thresholdConf/saveThresholdConf [post]
func (t *ThresholdConfApi) SaveThresholdConf(c *gin.Context) {
	var req dto.SaveThresholdConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查阈值配置是否已存在
	exists, err := repo.CheckThresholdConfExists(req.ParamDesc, req.SophonType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查阈值配置失败"))
		return
	}
	if exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该参数描述和Sophon类型的阈值配置已存在"))
		return
	}

	thresholdConf := &model.ThresholdConf{
		ParamDesc:  req.ParamDesc,
		ParamValue: req.ParamValue,
		SophonType: req.SophonType,
	}

	if err := repo.SaveThresholdConf(thresholdConf); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(thresholdConf))
}

// UpdateThresholdConf 更新阈值配置
func (t *ThresholdConfApi) UpdateThresholdConf(c *gin.Context) {
	var req dto.UpdateThresholdConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查阈值配置是否存在
	existingConf, err := repo.GetThresholdConfByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "阈值配置不存在"))
		return
	}

	// 如果参数描述或Sophon类型发生变化，检查新的组合是否已存在
	if existingConf.ParamDesc != req.ParamDesc || existingConf.SophonType != req.SophonType {
		exists, err := repo.CheckThresholdConfExists(req.ParamDesc, req.SophonType)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查阈值配置失败"))
			return
		}
		if exists {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该参数描述和Sophon类型的阈值配置已存在"))
			return
		}
	}

	thresholdConf := &model.ThresholdConf{
		ID:         req.ID,
		ParamDesc:  req.ParamDesc,
		ParamValue: req.ParamValue,
		SophonType: req.SophonType,
	}

	if err := repo.UpdateThresholdConf(thresholdConf); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(thresholdConf))
}

// UpdateThresholdValue 更新阈值参数值
func (t *ThresholdConfApi) UpdateThresholdValue(c *gin.Context) {
	var req dto.UpdateThresholdValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.UpdateThresholdValue(req.ID, req.ParamValue); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新阈值失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// BatchUpdateThresholdConf 批量更新阈值配置
func (t *ThresholdConfApi) BatchUpdateThresholdConf(c *gin.Context) {
	var req dto.BatchUpdateThresholdConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if len(req.Thresholds) == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "阈值配置列表不能为空"))
		return
	}

	var thresholdConfs []model.ThresholdConf
	for _, thresholdReq := range req.Thresholds {
		thresholdConfs = append(thresholdConfs, model.ThresholdConf{
			ID:         thresholdReq.ID,
			ParamDesc:  thresholdReq.ParamDesc,
			ParamValue: thresholdReq.ParamValue,
			SophonType: thresholdReq.SophonType,
		})
	}

	if err := repo.BatchUpdateThresholdConf(thresholdConfs); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "批量更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelThresholdConf 删除阈值配置
func (t *ThresholdConfApi) DelThresholdConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteThresholdConf(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetThresholdConf 根据ID获取阈值配置
func (t *ThresholdConfApi) GetThresholdConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	thresholdConf, err := repo.GetThresholdConfByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.ThresholdConfResponse{
		ID:         thresholdConf.ID,
		ParamDesc:  thresholdConf.ParamDesc,
		ParamValue: thresholdConf.ParamValue,
		SophonType: thresholdConf.SophonType,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetThresholdConfAll 获取所有阈值配置
func (t *ThresholdConfApi) GetThresholdConfAll(c *gin.Context) {
	thresholdConfs, err := repo.GetThresholdConfAll()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ThresholdConfResponse
	for _, thresholdConf := range thresholdConfs {
		responses = append(responses, dto.ThresholdConfResponse{
			ID:         thresholdConf.ID,
			ParamDesc:  thresholdConf.ParamDesc,
			ParamValue: thresholdConf.ParamValue,
			SophonType: thresholdConf.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetThresholdConfPage 分页获取阈值配置列表
func (t *ThresholdConfApi) GetThresholdConfPage(c *gin.Context) {
	var req dto.ThresholdConfPageRequest
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

	thresholdConfs, total, err := repo.GetThresholdConfPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ThresholdConfResponse
	for _, thresholdConf := range thresholdConfs {
		responses = append(responses, dto.ThresholdConfResponse{
			ID:         thresholdConf.ID,
			ParamDesc:  thresholdConf.ParamDesc,
			ParamValue: thresholdConf.ParamValue,
			SophonType: thresholdConf.SophonType,
		})
	}

	response := dto.ThresholdConfPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetThresholdConfBySophonType 根据Sophon类型获取阈值配置列表
func (t *ThresholdConfApi) GetThresholdConfBySophonType(c *gin.Context) {
	sophonTypeStr := c.Param("sophonType")
	sophonType, err := strconv.Atoi(sophonTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	thresholdConfs, err := repo.GetThresholdConfBySophonType(sophonType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var thresholds []dto.ThresholdConfResponse
	for _, thresholdConf := range thresholdConfs {
		thresholds = append(thresholds, dto.ThresholdConfResponse{
			ID:         thresholdConf.ID,
			ParamDesc:  thresholdConf.ParamDesc,
			ParamValue: thresholdConf.ParamValue,
			SophonType: thresholdConf.SophonType,
		})
	}

	response := dto.ThresholdConfBySophonTypeResponse{
		SophonType: sophonType,
		Thresholds: thresholds,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// QueryThresholdConf 根据查询条件获取阈值配置列表
func (t *ThresholdConfApi) QueryThresholdConf(c *gin.Context) {
	var req dto.ThresholdConfQueryRequest
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

	thresholdConfs, total, err := repo.GetThresholdConfByQuery(
		req.ParamDesc,
		req.SophonType,
		req.MinValue,
		req.MaxValue,
		req.Page,
		req.PageSize,
	)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ThresholdConfResponse
	for _, thresholdConf := range thresholdConfs {
		responses = append(responses, dto.ThresholdConfResponse{
			ID:         thresholdConf.ID,
			ParamDesc:  thresholdConf.ParamDesc,
			ParamValue: thresholdConf.ParamValue,
			SophonType: thresholdConf.SophonType,
		})
	}

	response := dto.ThresholdConfPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetThresholdConfStatistics 获取阈值配置统计
func (t *ThresholdConfApi) GetThresholdConfStatistics(c *gin.Context) {
	statistics, err := repo.GetThresholdConfStatistics()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取统计失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(statistics))
}

// SearchThresholdConf 搜索阈值配置
func (t *ThresholdConfApi) SearchThresholdConf(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索关键词不能为空"))
		return
	}

	thresholdConfs, err := repo.SearchThresholdConf(keyword)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索失败"))
		return
	}

	var responses []dto.ThresholdConfResponse
	for _, thresholdConf := range thresholdConfs {
		responses = append(responses, dto.ThresholdConfResponse{
			ID:         thresholdConf.ID,
			ParamDesc:  thresholdConf.ParamDesc,
			ParamValue: thresholdConf.ParamValue,
			SophonType: thresholdConf.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetThresholdConfLists 获取参数描述和Sophon类型列表
func (t *ThresholdConfApi) GetThresholdConfLists(c *gin.Context) {
	paramDescs, err := repo.GetAllParamDescs()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取参数描述列表失败"))
		return
	}

	sophonTypes, err := repo.GetAllSophonTypes()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取Sophon类型列表失败"))
		return
	}

	response := dto.ThresholdConfListResponse{
		ParamDescs:  paramDescs,
		SophonTypes: sophonTypes,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetThresholdValue 获取阈值参数值
func (t *ThresholdConfApi) GetThresholdValue(c *gin.Context) {
	paramDesc := c.Query("param_desc")
	sophonTypeStr := c.Query("sophon_type")

	if paramDesc == "" || sophonTypeStr == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数描述和Sophon类型不能为空"))
		return
	}

	sophonType, err := strconv.Atoi(sophonTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "Sophon类型参数错误"))
		return
	}

	paramValue, err := repo.GetThresholdValue(paramDesc, sophonType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取阈值失败"))
		return
	}

	response := map[string]interface{}{
		"param_desc":  paramDesc,
		"param_value": paramValue,
		"sophon_type": sophonType,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// SetThresholdValue 设置阈值参数值
func (t *ThresholdConfApi) SetThresholdValue(c *gin.Context) {
	var req struct {
		ParamDesc  string  `json:"param_desc" binding:"required"`
		ParamValue float64 `json:"param_value"`
		SophonType int     `json:"sophon_type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.SetThresholdValue(req.ParamDesc, req.ParamValue, req.SophonType); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置阈值失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}
