package system

import (
	"feishuReboot/pkg/dto"
	"feishuReboot/pkg/handle"
	"feishuReboot/pkg/model"
	"feishuReboot/pkg/repo"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ModelsInfoApi struct{}

// SaveModelsInfo 保存模型信息
// @Summary 保存模型信息
// @Description 保存新的模型信息
// @Tags 模型信息
// @Accept json
// @Produce json
// @Param request body dto.SaveModelsInfoRequest true "模型信息"
// @Success 200 {object} handle.Response{data=model.ModelsInfo} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /modelsInfo/saveModelsInfo [post]
func (m *ModelsInfoApi) SaveModelsInfo(c *gin.Context) {
	var req dto.SaveModelsInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查模型名称是否已存在
	exists, err := repo.CheckModelsInfoExists(req.Name)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查模型失败"))
		return
	}
	if exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "模型名称已存在"))
		return
	}

	// 如果没有提供日期时间，使用当前时间
	if req.DateTime == "" {
		req.DateTime = time.Now().Format("2006-01-02 15:04:05")
	}

	modelsInfo := &model.ModelsInfo{
		Name:        req.Name,
		Algorithm:   req.Algorithm,
		DateTime:    req.DateTime,
		Description: req.Description,
		SophonType:  req.SophonType,
		Classify:    req.Classify,
		Specialty:   req.Specialty,
	}

	if err := repo.SaveModelsInfo(modelsInfo); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(modelsInfo))
}

// UpdateModelsInfo 更新模型信息
func (m *ModelsInfoApi) UpdateModelsInfo(c *gin.Context) {
	var req dto.UpdateModelsInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查模型是否存在
	existingModel, err := repo.GetModelsInfoByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "模型不存在"))
		return
	}

	// 如果名称发生变化，检查新名称是否已被其他模型使用
	if existingModel.Name != req.Name {
		exists, err := repo.CheckModelsInfoExists(req.Name)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查模型失败"))
			return
		}
		if exists {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "模型名称已存在"))
			return
		}
	}

	modelsInfo := &model.ModelsInfo{
		ID:          req.ID,
		Name:        req.Name,
		Algorithm:   req.Algorithm,
		DateTime:    req.DateTime,
		Description: req.Description,
		SophonType:  req.SophonType,
		Classify:    req.Classify,
		Specialty:   req.Specialty,
	}

	if err := repo.UpdateModelsInfo(modelsInfo); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(modelsInfo))
}

// BatchUpdateModelsInfo 批量更新模型信息
func (m *ModelsInfoApi) BatchUpdateModelsInfo(c *gin.Context) {
	var req dto.BatchUpdateModelsInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if len(req.Models) == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "模型列表不能为空"))
		return
	}

	var modelsInfos []model.ModelsInfo
	for _, modelReq := range req.Models {
		modelsInfos = append(modelsInfos, model.ModelsInfo{
			ID:          modelReq.ID,
			Name:        modelReq.Name,
			Algorithm:   modelReq.Algorithm,
			DateTime:    modelReq.DateTime,
			Description: modelReq.Description,
			SophonType:  modelReq.SophonType,
			Classify:    modelReq.Classify,
			Specialty:   modelReq.Specialty,
		})
	}

	if err := repo.BatchUpdateModelsInfo(modelsInfos); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "批量更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelModelsInfo 删除模型信息
func (m *ModelsInfoApi) DelModelsInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteModelsInfo(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetModelsInfo 根据ID获取模型信息
// @Summary 根据ID获取模型信息
// @Description 根据模型ID获取详细信息
// @Tags 模型信息
// @Accept json
// @Produce json
// @Param id path int true "模型ID"
// @Success 200 {object} handle.Response{data=dto.ModelsInfoResponse} "获取成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /modelsInfo/getModelsInfo/{id} [get]
func (m *ModelsInfoApi) GetModelsInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	modelsInfo, err := repo.GetModelsInfoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.ModelsInfoResponse{
		ID:          modelsInfo.ID,
		Name:        modelsInfo.Name,
		Algorithm:   modelsInfo.Algorithm,
		DateTime:    modelsInfo.DateTime,
		Description: modelsInfo.Description,
		SophonType:  modelsInfo.SophonType,
		Classify:    modelsInfo.Classify,
		Specialty:   modelsInfo.Specialty,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetModelsInfoAll 获取所有模型信息
// @Summary 获取所有模型信息
// @Description 获取系统中所有的模型信息列表
// @Tags 模型信息
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]dto.ModelsInfoResponse} "获取成功"
// @Failure 400 {object} handle.Response "获取失败"
// @Router /modelsInfo/getModelsInfoAll [get]
func (m *ModelsInfoApi) GetModelsInfoAll(c *gin.Context) {
	modelsInfos, err := repo.GetModelsInfoAll()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		responses = append(responses, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetModelsInfoPage 分页获取模型信息列表
func (m *ModelsInfoApi) GetModelsInfoPage(c *gin.Context) {
	var req dto.ModelsInfoPageRequest
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

	modelsInfos, total, err := repo.GetModelsInfoPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		responses = append(responses, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	response := dto.ModelsInfoPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetModelsInfoByAlgorithm 根据算法获取模型信息列表
func (m *ModelsInfoApi) GetModelsInfoByAlgorithm(c *gin.Context) {
	algorithm := c.Param("algorithm")
	if algorithm == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "算法参数不能为空"))
		return
	}

	modelsInfos, err := repo.GetModelsInfoByAlgorithm(algorithm)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var models []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		models = append(models, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	response := dto.ModelsInfoByAlgorithmResponse{
		Algorithm: algorithm,
		Models:    models,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetModelsInfoByClassify 根据分类获取模型信息列表
func (m *ModelsInfoApi) GetModelsInfoByClassify(c *gin.Context) {
	classify := c.Param("classify")
	if classify == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "分类参数不能为空"))
		return
	}

	modelsInfos, err := repo.GetModelsInfoByClassify(classify)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var models []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		models = append(models, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	response := dto.ModelsInfoByClassifyResponse{
		Classify: classify,
		Models:   models,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetModelsInfoBySophonType 根据Sophon类型获取模型信息列表
func (m *ModelsInfoApi) GetModelsInfoBySophonType(c *gin.Context) {
	sophonTypeStr := c.Param("sophonType")
	sophonType, err := strconv.Atoi(sophonTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	modelsInfos, err := repo.GetModelsInfoBySophonType(sophonType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		responses = append(responses, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// QueryModelsInfo 根据查询条件获取模型信息列表
func (m *ModelsInfoApi) QueryModelsInfo(c *gin.Context) {
	var req dto.ModelsInfoQueryRequest
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

	modelsInfos, total, err := repo.GetModelsInfoByQuery(
		req.Name,
		req.Algorithm,
		req.Classify,
		req.SophonType,
		req.Specialty,
		req.Page,
		req.PageSize,
	)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		responses = append(responses, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	response := dto.ModelsInfoPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetModelsInfoStatistics 获取模型信息统计
func (m *ModelsInfoApi) GetModelsInfoStatistics(c *gin.Context) {
	statistics, err := repo.GetModelsInfoStatistics()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取统计失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(statistics))
}

// SearchModelsInfo 搜索模型信息
func (m *ModelsInfoApi) SearchModelsInfo(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索关键词不能为空"))
		return
	}

	modelsInfos, err := repo.SearchModelsInfo(keyword)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索失败"))
		return
	}

	var responses []dto.ModelsInfoResponse
	for _, modelsInfo := range modelsInfos {
		responses = append(responses, dto.ModelsInfoResponse{
			ID:          modelsInfo.ID,
			Name:        modelsInfo.Name,
			Algorithm:   modelsInfo.Algorithm,
			DateTime:    modelsInfo.DateTime,
			Description: modelsInfo.Description,
			SophonType:  modelsInfo.SophonType,
			Classify:    modelsInfo.Classify,
			Specialty:   modelsInfo.Specialty,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAllAlgorithms 获取所有算法列表
func (m *ModelsInfoApi) GetAllAlgorithms(c *gin.Context) {
	algorithms, err := repo.GetAllAlgorithms()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取算法列表失败"))
		return
	}

	response := map[string][]string{
		"algorithms": algorithms,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAllClassifies 获取所有分类列表
func (m *ModelsInfoApi) GetAllClassifies(c *gin.Context) {
	classifies, err := repo.GetAllClassifies()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取分类列表失败"))
		return
	}

	response := map[string][]string{
		"classifies": classifies,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}
