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

type PlanTemplateApi struct{}

// SavePlanTemplate 保存计划模板
// @Summary 保存计划模板
// @Description 保存新的计划模板信息
// @Tags 计划模板
// @Accept json
// @Produce json
// @Param request body dto.SavePlanTemplateRequest true "计划模板信息"
// @Success 200 {object} handle.Response{data=model.PlanTemplate} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /planTemplate/savePlanTemplate [post]
func (p *PlanTemplateApi) SavePlanTemplate(c *gin.Context) {
	var req dto.SavePlanTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	planTemplate := &model.PlanTemplate{
		Name: req.Name,
		Plan: req.Plan,
	}

	if err := repo.SavePlanTemplate(planTemplate); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(planTemplate))
}

// UpdatePlanTemplate 更新计划模板
// @Summary 更新计划模板
// @Description 更新已有的计划模板信息
// @Tags 计划模板
// @Accept json
// @Produce json
// @Param request body dto.UpdatePlanTemplateRequest true "计划模板更新信息"
// @Success 200 {object} handle.Response{data=model.PlanTemplate} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /planTemplate/updatePlanTemplate [post]
func (p *PlanTemplateApi) UpdatePlanTemplate(c *gin.Context) {
	var req dto.UpdatePlanTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	planTemplate := &model.PlanTemplate{
		ID:   req.ID,
		Name: req.Name,
		Plan: req.Plan,
	}

	if err := repo.UpdatePlanTemplate(planTemplate); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(planTemplate))
}

// DelPlanTemplate 删除计划模板
func (p *PlanTemplateApi) DelPlanTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeletePlanTemplate(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetPlanTemplate 根据ID获取计划模板
func (p *PlanTemplateApi) GetPlanTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	planTemplate, err := repo.GetPlanTemplateByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.PlanTemplateResponse{
		ID:   planTemplate.ID,
		Name: planTemplate.Name,
		Plan: planTemplate.Plan,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetPlanTemplateList 获取计划模板列表
func (p *PlanTemplateApi) GetPlanTemplateList(c *gin.Context) {
	planTemplates, err := repo.GetPlanTemplateList()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.PlanTemplateResponse
	for _, planTemplate := range planTemplates {
		responses = append(responses, dto.PlanTemplateResponse{
			ID:   planTemplate.ID,
			Name: planTemplate.Name,
			Plan: planTemplate.Plan,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetPlanTemplatePage 分页获取计划模板列表
func (p *PlanTemplateApi) GetPlanTemplatePage(c *gin.Context) {
	var req dto.PlanTemplatePageRequest
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

	planTemplates, total, err := repo.GetPlanTemplatePage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.PlanTemplateResponse
	for _, planTemplate := range planTemplates {
		responses = append(responses, dto.PlanTemplateResponse{
			ID:   planTemplate.ID,
			Name: planTemplate.Name,
			Plan: planTemplate.Plan,
		})
	}

	response := dto.PlanTemplatePageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}
