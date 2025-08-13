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

type TaskManageApi struct{}

// SaveTaskManage 保存任务管理
// @Summary 保存任务管理
// @Description 保存新的任务管理信息
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param request body dto.SaveTaskManageRequest true "任务管理信息"
// @Success 200 {object} handle.Response{data=model.TaskManage} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /taskManage/saveTaskManage [post]
func (t *TaskManageApi) SaveTaskManage(c *gin.Context) {
	var req dto.SaveTaskManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	taskManage := &model.TaskManage{
		Name:             req.Name,
		MediaSource:      req.MediaSource,
		WorkPlan:         req.WorkPlan,
		ArithmeticConfig: req.ArithmeticConfig,
		TaskStatus:       req.TaskStatus,
		MediaID:          req.MediaID,
		PlanID:           req.PlanID,
		SophonTypes:      req.SophonTypes,
		ReportAddr:       req.ReportAddr,
		AlarmInterval:    req.AlarmInterval,
		AritParams:       req.AritParams,
		StatusDesc:       req.StatusDesc,
		UsedIP:           req.UsedIP,
	}

	if err := repo.SaveTaskManage(taskManage); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(taskManage))
}

// UpdateTaskManage 更新任务管理
// @Summary 更新任务管理
// @Description 更新已有的任务管理信息
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param request body dto.UpdateTaskManageRequest true "任务管理更新信息"
// @Success 200 {object} handle.Response{data=model.TaskManage} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /taskManage/updateTaskManage [post]
func (t *TaskManageApi) UpdateTaskManage(c *gin.Context) {
	var req dto.UpdateTaskManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	taskManage := &model.TaskManage{
		ID:               req.ID,
		Name:             req.Name,
		MediaSource:      req.MediaSource,
		WorkPlan:         req.WorkPlan,
		ArithmeticConfig: req.ArithmeticConfig,
		TaskStatus:       req.TaskStatus,
		MediaID:          req.MediaID,
		PlanID:           req.PlanID,
		SophonTypes:      req.SophonTypes,
		ReportAddr:       req.ReportAddr,
		AlarmInterval:    req.AlarmInterval,
		AritParams:       req.AritParams,
		StatusDesc:       req.StatusDesc,
		UsedIP:           req.UsedIP,
	}

	if err := repo.UpdateTaskManage(taskManage); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(taskManage))
}

// UpdateTaskStatus 更新任务状态
func (t *TaskManageApi) UpdateTaskStatus(c *gin.Context) {
	var req dto.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.UpdateTaskStatus(req.ID, req.TaskStatus); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新状态失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelTaskManage 删除任务管理
func (t *TaskManageApi) DelTaskManage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteTaskManage(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetTaskManage 根据ID获取任务管理
func (t *TaskManageApi) GetTaskManage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	taskManage, err := repo.GetTaskManageByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.TaskManageResponse{
		ID:               taskManage.ID,
		Name:             taskManage.Name,
		MediaSource:      taskManage.MediaSource,
		WorkPlan:         taskManage.WorkPlan,
		ArithmeticConfig: taskManage.ArithmeticConfig,
		TaskStatus:       taskManage.TaskStatus,
		MediaID:          taskManage.MediaID,
		PlanID:           taskManage.PlanID,
		SophonTypes:      taskManage.SophonTypes,
		ReportAddr:       taskManage.ReportAddr,
		AlarmInterval:    taskManage.AlarmInterval,
		AritParams:       taskManage.AritParams,
		StatusDesc:       taskManage.StatusDesc,
		UsedIP:           taskManage.UsedIP,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetTaskManageList 获取任务管理列表
func (t *TaskManageApi) GetTaskManageList(c *gin.Context) {
	taskManages, err := repo.GetTaskManageList()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.TaskManageResponse
	for _, taskManage := range taskManages {
		responses = append(responses, dto.TaskManageResponse{
			ID:               taskManage.ID,
			Name:             taskManage.Name,
			MediaSource:      taskManage.MediaSource,
			WorkPlan:         taskManage.WorkPlan,
			ArithmeticConfig: taskManage.ArithmeticConfig,
			TaskStatus:       taskManage.TaskStatus,
			MediaID:          taskManage.MediaID,
			PlanID:           taskManage.PlanID,
			SophonTypes:      taskManage.SophonTypes,
			ReportAddr:       taskManage.ReportAddr,
			AlarmInterval:    taskManage.AlarmInterval,
			AritParams:       taskManage.AritParams,
			StatusDesc:       taskManage.StatusDesc,
			UsedIP:           taskManage.UsedIP,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetTaskManagePage 分页获取任务管理列表
func (t *TaskManageApi) GetTaskManagePage(c *gin.Context) {
	var req dto.TaskManagePageRequest
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

	taskManages, total, err := repo.GetTaskManagePage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.TaskManageResponse
	for _, taskManage := range taskManages {
		responses = append(responses, dto.TaskManageResponse{
			ID:               taskManage.ID,
			Name:             taskManage.Name,
			MediaSource:      taskManage.MediaSource,
			WorkPlan:         taskManage.WorkPlan,
			ArithmeticConfig: taskManage.ArithmeticConfig,
			TaskStatus:       taskManage.TaskStatus,
			MediaID:          taskManage.MediaID,
			PlanID:           taskManage.PlanID,
			SophonTypes:      taskManage.SophonTypes,
			ReportAddr:       taskManage.ReportAddr,
			AlarmInterval:    taskManage.AlarmInterval,
			AritParams:       taskManage.AritParams,
			StatusDesc:       taskManage.StatusDesc,
			UsedIP:           taskManage.UsedIP,
		})
	}

	response := dto.TaskManagePageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetTaskManageByStatus 根据状态获取任务列表
func (t *TaskManageApi) GetTaskManageByStatus(c *gin.Context) {
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	taskManages, err := repo.GetTaskManageByStatus(status)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.TaskManageResponse
	for _, taskManage := range taskManages {
		responses = append(responses, dto.TaskManageResponse{
			ID:               taskManage.ID,
			Name:             taskManage.Name,
			MediaSource:      taskManage.MediaSource,
			WorkPlan:         taskManage.WorkPlan,
			ArithmeticConfig: taskManage.ArithmeticConfig,
			TaskStatus:       taskManage.TaskStatus,
			MediaID:          taskManage.MediaID,
			PlanID:           taskManage.PlanID,
			SophonTypes:      taskManage.SophonTypes,
			ReportAddr:       taskManage.ReportAddr,
			AlarmInterval:    taskManage.AlarmInterval,
			AritParams:       taskManage.AritParams,
			StatusDesc:       taskManage.StatusDesc,
			UsedIP:           taskManage.UsedIP,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetTaskManageByMediaID 根据媒体ID获取任务列表
func (t *TaskManageApi) GetTaskManageByMediaID(c *gin.Context) {
	mediaIDStr := c.Param("mediaId")
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	taskManages, err := repo.GetTaskManageByMediaID(mediaID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.TaskManageResponse
	for _, taskManage := range taskManages {
		responses = append(responses, dto.TaskManageResponse{
			ID:               taskManage.ID,
			Name:             taskManage.Name,
			MediaSource:      taskManage.MediaSource,
			WorkPlan:         taskManage.WorkPlan,
			ArithmeticConfig: taskManage.ArithmeticConfig,
			TaskStatus:       taskManage.TaskStatus,
			MediaID:          taskManage.MediaID,
			PlanID:           taskManage.PlanID,
			SophonTypes:      taskManage.SophonTypes,
			ReportAddr:       taskManage.ReportAddr,
			AlarmInterval:    taskManage.AlarmInterval,
			AritParams:       taskManage.AritParams,
			StatusDesc:       taskManage.StatusDesc,
			UsedIP:           taskManage.UsedIP,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}
