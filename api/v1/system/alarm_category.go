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

type AlarmCategoryApi struct{}

// SaveAlarmCategory 保存告警分类
// @Summary 保存告警分类
// @Description 保存新的告警分类信息
// @Tags 告警分类
// @Accept json
// @Produce json
// @Param request body dto.SaveAlarmCategoryRequest true "告警分类信息"
// @Success 200 {object} handle.Response{data=model.AlarmCategory} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /alarmCategory/saveAlarmCategory [post]
func (a *AlarmCategoryApi) SaveAlarmCategory(c *gin.Context) {
	var req dto.SaveAlarmCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查告警分类是否已存在
	exists, err := repo.CheckAlarmCategoryExists(req.Arithmetic, req.AlarmType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查告警分类失败"))
		return
	}
	if exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该算法下的告警类型已存在"))
		return
	}

	alarmCategory := &model.AlarmCategory{
		Arithmetic: req.Arithmetic,
		AlarmType:  req.AlarmType,
		AlarmDesc:  req.AlarmDesc,
		AudioFile:  req.AudioFile,
		SophonType: req.SophonType,
	}

	if err := repo.SaveAlarmCategory(alarmCategory); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(alarmCategory))
}

// UpdateAlarmCategory 更新告警分类
func (a *AlarmCategoryApi) UpdateAlarmCategory(c *gin.Context) {
	var req dto.UpdateAlarmCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查告警分类是否存在
	existingCategory, err := repo.GetAlarmCategoryByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "告警分类不存在"))
		return
	}

	// 如果算法或告警类型发生变化，检查新的组合是否已存在
	if existingCategory.Arithmetic != req.Arithmetic || existingCategory.AlarmType != req.AlarmType {
		exists, err := repo.CheckAlarmCategoryExists(req.Arithmetic, req.AlarmType)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查告警分类失败"))
			return
		}
		if exists {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "该算法下的告警类型已存在"))
			return
		}
	}

	alarmCategory := &model.AlarmCategory{
		ID:         req.ID,
		Arithmetic: req.Arithmetic,
		AlarmType:  req.AlarmType,
		AlarmDesc:  req.AlarmDesc,
		AudioFile:  req.AudioFile,
		SophonType: req.SophonType,
	}

	if err := repo.UpdateAlarmCategory(alarmCategory); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(alarmCategory))
}

// BatchUpdateAlarmCategory 批量更新告警分类
func (a *AlarmCategoryApi) BatchUpdateAlarmCategory(c *gin.Context) {
	var req dto.BatchUpdateAlarmCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if len(req.Categories) == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "告警分类列表不能为空"))
		return
	}

	var alarmCategories []model.AlarmCategory
	for _, categoryReq := range req.Categories {
		alarmCategories = append(alarmCategories, model.AlarmCategory{
			ID:         categoryReq.ID,
			Arithmetic: categoryReq.Arithmetic,
			AlarmType:  categoryReq.AlarmType,
			AlarmDesc:  categoryReq.AlarmDesc,
			AudioFile:  categoryReq.AudioFile,
			SophonType: categoryReq.SophonType,
		})
	}

	if err := repo.BatchUpdateAlarmCategory(alarmCategories); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "批量更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelAlarmCategory 删除告警分类
func (a *AlarmCategoryApi) DelAlarmCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteAlarmCategory(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetAlarmCategory 根据ID获取告警分类
func (a *AlarmCategoryApi) GetAlarmCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmCategory, err := repo.GetAlarmCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.AlarmCategoryResponse{
		ID:         alarmCategory.ID,
		Arithmetic: alarmCategory.Arithmetic,
		AlarmType:  alarmCategory.AlarmType,
		AlarmDesc:  alarmCategory.AlarmDesc,
		AudioFile:  alarmCategory.AudioFile,
		SophonType: alarmCategory.SophonType,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoryAll 获取所有告警分类
func (a *AlarmCategoryApi) GetAlarmCategoryAll(c *gin.Context) {
	alarmCategories, err := repo.GetAlarmCategoryAll()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmCategoryPage 分页获取告警分类列表
func (a *AlarmCategoryApi) GetAlarmCategoryPage(c *gin.Context) {
	var req dto.AlarmCategoryPageRequest
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

	alarmCategories, total, err := repo.GetAlarmCategoryPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	response := dto.AlarmCategoryPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoryByArithmetic 根据算法获取告警分类列表
func (a *AlarmCategoryApi) GetAlarmCategoryByArithmetic(c *gin.Context) {
	arithmetic := c.Param("arithmetic")
	if arithmetic == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "算法参数不能为空"))
		return
	}

	alarmCategories, err := repo.GetAlarmCategoryByArithmetic(arithmetic)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var categories []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		categories = append(categories, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	response := dto.AlarmCategoryByArithmeticResponse{
		Arithmetic: arithmetic,
		Categories: categories,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoryByAlarmType 根据告警类型获取告警分类列表
func (a *AlarmCategoryApi) GetAlarmCategoryByAlarmType(c *gin.Context) {
	alarmType := c.Param("alarmType")
	if alarmType == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "告警类型参数不能为空"))
		return
	}

	alarmCategories, err := repo.GetAlarmCategoryByAlarmType(alarmType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var categories []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		categories = append(categories, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	response := dto.AlarmCategoryByTypeResponse{
		AlarmType:  alarmType,
		Categories: categories,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoryBySophonType 根据Sophon类型获取告警分类列表
func (a *AlarmCategoryApi) GetAlarmCategoryBySophonType(c *gin.Context) {
	sophonTypeStr := c.Param("sophonType")
	sophonType, err := strconv.Atoi(sophonTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmCategories, err := repo.GetAlarmCategoryBySophonType(sophonType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// QueryAlarmCategory 根据查询条件获取告警分类列表
func (a *AlarmCategoryApi) QueryAlarmCategory(c *gin.Context) {
	var req dto.AlarmCategoryQueryRequest
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

	alarmCategories, total, err := repo.GetAlarmCategoryByQuery(
		req.Arithmetic,
		req.AlarmType,
		req.SophonType,
		req.Page,
		req.PageSize,
	)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	response := dto.AlarmCategoryPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoryStatistics 获取告警分类统计
func (a *AlarmCategoryApi) GetAlarmCategoryStatistics(c *gin.Context) {
	statistics, err := repo.GetAlarmCategoryStatistics()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取统计失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(statistics))
}

// SearchAlarmCategory 搜索告警分类
func (a *AlarmCategoryApi) SearchAlarmCategory(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索关键词不能为空"))
		return
	}

	alarmCategories, err := repo.SearchAlarmCategory(keyword)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmCategoryLists 获取算法和告警类型列表
func (a *AlarmCategoryApi) GetAlarmCategoryLists(c *gin.Context) {
	arithmetics, err := repo.GetAllArithmetics()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取算法列表失败"))
		return
	}

	alarmTypes, err := repo.GetAllAlarmTypes()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取告警类型列表失败"))
		return
	}

	response := dto.AlarmCategoryListResponse{
		Arithmetics: arithmetics,
		AlarmTypes:  alarmTypes,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmCategoriesWithAudio 获取有语音文件的告警分类
func (a *AlarmCategoryApi) GetAlarmCategoriesWithAudio(c *gin.Context) {
	alarmCategories, err := repo.GetAlarmCategoriesWithAudio()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmCategoriesWithoutAudio 获取无语音文件的告警分类
func (a *AlarmCategoryApi) GetAlarmCategoriesWithoutAudio(c *gin.Context) {
	alarmCategories, err := repo.GetAlarmCategoriesWithoutAudio()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmCategoryResponse
	for _, alarmCategory := range alarmCategories {
		responses = append(responses, dto.AlarmCategoryResponse{
			ID:         alarmCategory.ID,
			Arithmetic: alarmCategory.Arithmetic,
			AlarmType:  alarmCategory.AlarmType,
			AlarmDesc:  alarmCategory.AlarmDesc,
			AudioFile:  alarmCategory.AudioFile,
			SophonType: alarmCategory.SophonType,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}
