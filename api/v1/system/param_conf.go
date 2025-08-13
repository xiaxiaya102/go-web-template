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

type ParamConfApi struct{}

// SaveParamConf 保存参数配置
// @Summary 保存参数配置
// @Description 保存新的参数配置信息
// @Tags 参数配置
// @Accept json
// @Produce json
// @Param request body dto.SaveParamConfRequest true "参数配置信息"
// @Success 200 {object} handle.Response{data=model.ParamConf} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /paramConf/saveParamConf [post]
func (p *ParamConfApi) SaveParamConf(c *gin.Context) {
	var req dto.SaveParamConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查配置是否已存在
	exists, err := repo.CheckParamConfExists(req.Key, req.Param)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查配置失败"))
		return
	}
	if exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置已存在"))
		return
	}

	paramConf := &model.ParamConf{
		Key:        req.Key,
		Param:      req.Param,
		ParamValue: req.ParamValue,
	}

	if err := repo.SaveParamConf(paramConf); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(paramConf))
}

// UpdateParamConf 更新参数配置
func (p *ParamConfApi) UpdateParamConf(c *gin.Context) {
	var req dto.UpdateParamConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查配置是否存在
	existingConf, err := repo.GetParamConfByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置不存在"))
		return
	}

	// 如果key或param发生变化，检查新的组合是否已存在
	if existingConf.Key != req.Key || existingConf.Param != req.Param {
		exists, err := repo.CheckParamConfExists(req.Key, req.Param)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查配置失败"))
			return
		}
		if exists {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置已存在"))
			return
		}
	}

	paramConf := &model.ParamConf{
		ID:         req.ID,
		Key:        req.Key,
		Param:      req.Param,
		ParamValue: req.ParamValue,
	}

	if err := repo.UpdateParamConf(paramConf); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(paramConf))
}

// UpdateParamValue 更新参数值
func (p *ParamConfApi) UpdateParamValue(c *gin.Context) {
	var req dto.UpdateParamValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.UpdateParamValue(req.ID, req.ParamValue); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新参数值失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// BatchUpdateParamConf 批量更新参数配置
func (p *ParamConfApi) BatchUpdateParamConf(c *gin.Context) {
	var req dto.BatchUpdateParamConfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if len(req.Configs) == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置列表不能为空"))
		return
	}

	var paramConfs []model.ParamConf
	for _, config := range req.Configs {
		paramConfs = append(paramConfs, model.ParamConf{
			ID:         config.ID,
			Key:        config.Key,
			Param:      config.Param,
			ParamValue: config.ParamValue,
		})
	}

	if err := repo.BatchUpdateParamConf(paramConfs); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "批量更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelParamConf 删除参数配置
func (p *ParamConfApi) DelParamConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteParamConf(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelParamConfByKey 根据配置键删除所有相关配置
func (p *ParamConfApi) DelParamConfByKey(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置键不能为空"))
		return
	}

	if err := repo.DeleteParamConfByKey(key); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetParamConf 根据ID获取参数配置
func (p *ParamConfApi) GetParamConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	paramConf, err := repo.GetParamConfByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.ParamConfResponse{
		ID:         paramConf.ID,
		Key:        paramConf.Key,
		Param:      paramConf.Param,
		ParamValue: paramConf.ParamValue,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetParamConfByKey 根据配置键获取参数配置列表
func (p *ParamConfApi) GetParamConfByKey(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置键不能为空"))
		return
	}

	paramConfs, err := repo.GetParamConfByKey(key)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var params []dto.ParamConfItemResponse
	for _, paramConf := range paramConfs {
		params = append(params, dto.ParamConfItemResponse{
			ID:         paramConf.ID,
			Param:      paramConf.Param,
			ParamValue: paramConf.ParamValue,
		})
	}

	response := dto.ParamConfByKeyResponse{
		Key:    key,
		Params: params,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetParamConfAll 获取所有参数配置
func (p *ParamConfApi) GetParamConfAll(c *gin.Context) {
	paramConfs, err := repo.GetParamConfAll()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ParamConfResponse
	for _, paramConf := range paramConfs {
		responses = append(responses, dto.ParamConfResponse{
			ID:         paramConf.ID,
			Key:        paramConf.Key,
			Param:      paramConf.Param,
			ParamValue: paramConf.ParamValue,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetParamConfPage 分页获取参数配置列表
func (p *ParamConfApi) GetParamConfPage(c *gin.Context) {
	var req dto.ParamConfPageRequest
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

	paramConfs, total, err := repo.GetParamConfPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ParamConfResponse
	for _, paramConf := range paramConfs {
		responses = append(responses, dto.ParamConfResponse{
			ID:         paramConf.ID,
			Key:        paramConf.Key,
			Param:      paramConf.Param,
			ParamValue: paramConf.ParamValue,
		})
	}

	response := dto.ParamConfPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// QueryParamConf 根据查询条件获取参数配置列表
func (p *ParamConfApi) QueryParamConf(c *gin.Context) {
	var req dto.ParamConfQueryRequest
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

	paramConfs, total, err := repo.GetParamConfByQuery(req.Key, req.Param, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ParamConfResponse
	for _, paramConf := range paramConfs {
		responses = append(responses, dto.ParamConfResponse{
			ID:         paramConf.ID,
			Key:        paramConf.Key,
			Param:      paramConf.Param,
			ParamValue: paramConf.ParamValue,
		})
	}

	response := dto.ParamConfPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAllParamKeys 获取所有配置键
func (p *ParamConfApi) GetAllParamKeys(c *gin.Context) {
	keys, err := repo.GetAllParamKeys()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取配置键失败"))
		return
	}

	response := dto.ParamConfKeysResponse{
		Keys: keys,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetParamValue 获取参数值
func (p *ParamConfApi) GetParamValue(c *gin.Context) {
	key := c.Query("key")
	param := c.Query("param")

	if key == "" || param == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "配置键和参数名不能为空"))
		return
	}

	paramValue, err := repo.GetParamValue(key, param)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取参数值失败"))
		return
	}

	response := map[string]string{
		"key":         key,
		"param":       param,
		"param_value": paramValue,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// SetParamValue 设置参数值
func (p *ParamConfApi) SetParamValue(c *gin.Context) {
	var req struct {
		Key        string `json:"key" binding:"required"`
		Param      string `json:"param" binding:"required"`
		ParamValue string `json:"param_value"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.SetParamValue(req.Key, req.Param, req.ParamValue); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置参数值失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// SearchParamConf 搜索参数配置
func (p *ParamConfApi) SearchParamConf(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索关键词不能为空"))
		return
	}

	paramConfs, err := repo.SearchParamConf(keyword)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索失败"))
		return
	}

	var responses []dto.ParamConfResponse
	for _, paramConf := range paramConfs {
		responses = append(responses, dto.ParamConfResponse{
			ID:         paramConf.ID,
			Key:        paramConf.Key,
			Param:      paramConf.Param,
			ParamValue: paramConf.ParamValue,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}
