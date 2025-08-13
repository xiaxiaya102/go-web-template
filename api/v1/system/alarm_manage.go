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

type AlarmManageApi struct{}

// SaveAlarmManage 保存报警管理
// @Summary 保存报警管理
// @Description 保存新的报警管理信息
// @Tags 报警管理
// @Accept json
// @Produce json
// @Param request body dto.SaveAlarmManageRequest true "报警管理信息"
// @Success 200 {object} handle.Response{data=model.AlarmManage} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /alarmManage/saveAlarmManage [post]
func (a *AlarmManageApi) SaveAlarmManage(c *gin.Context) {
	var req dto.SaveAlarmManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManage := &model.AlarmManage{
		AlarmTask:    req.AlarmTask,
		VideoMedia:   req.VideoMedia,
		MediaURL:     req.MediaURL,
		Img:          req.Img,
		InitImg:      req.InitImg,
		AlarmType:    req.AlarmType,
		ReportStatus: req.ReportStatus,
		ReportAddr:   req.ReportAddr,
		AlarmVideo:   req.AlarmVideo,
		AlarmDetail:  req.AlarmDetail,
		Tm:           req.Tm,
		MediaID:      req.MediaID,
	}

	if err := repo.SaveAlarmManage(alarmManage); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(alarmManage))
}

// UpdateAlarmManage 更新报警管理
// @Summary 更新报警管理
// @Description 更新已有的报警管理信息
// @Tags 报警管理
// @Accept json
// @Produce json
// @Param request body dto.UpdateAlarmManageRequest true "报警管理更新信息"
// @Success 200 {object} handle.Response{data=model.AlarmManage} "更新成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /alarmManage/updateAlarmManage [post]
func (a *AlarmManageApi) UpdateAlarmManage(c *gin.Context) {
	var req dto.UpdateAlarmManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManage := &model.AlarmManage{
		ID:           req.ID,
		AlarmTask:    req.AlarmTask,
		VideoMedia:   req.VideoMedia,
		MediaURL:     req.MediaURL,
		Img:          req.Img,
		InitImg:      req.InitImg,
		AlarmType:    req.AlarmType,
		ReportStatus: req.ReportStatus,
		ReportAddr:   req.ReportAddr,
		AlarmVideo:   req.AlarmVideo,
		AlarmDetail:  req.AlarmDetail,
		Tm:           req.Tm,
		MediaID:      req.MediaID,
	}

	if err := repo.UpdateAlarmManage(alarmManage); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(alarmManage))
}

// UpdateReportStatus 更新报告状态
func (a *AlarmManageApi) UpdateReportStatus(c *gin.Context) {
	var req dto.UpdateReportStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.UpdateReportStatus(req.ID, req.ReportStatus); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新状态失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelAlarmManage 删除报警管理
func (a *AlarmManageApi) DelAlarmManage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteAlarmManage(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetAlarmManage 根据ID获取报警管理
func (a *AlarmManageApi) GetAlarmManage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManage, err := repo.GetAlarmManageByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.AlarmManageResponse{
		ID:           alarmManage.ID,
		AlarmTask:    alarmManage.AlarmTask,
		VideoMedia:   alarmManage.VideoMedia,
		MediaURL:     alarmManage.MediaURL,
		Img:          alarmManage.Img,
		InitImg:      alarmManage.InitImg,
		AlarmType:    alarmManage.AlarmType,
		ReportStatus: alarmManage.ReportStatus,
		ReportAddr:   alarmManage.ReportAddr,
		AlarmVideo:   alarmManage.AlarmVideo,
		AlarmDetail:  alarmManage.AlarmDetail,
		Tm:           alarmManage.Tm,
		MediaID:      alarmManage.MediaID,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmManageList 获取报警管理列表
func (a *AlarmManageApi) GetAlarmManageList(c *gin.Context) {
	alarmManages, err := repo.GetAlarmManageList()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmManagePage 分页获取报警管理列表
func (a *AlarmManageApi) GetAlarmManagePage(c *gin.Context) {
	var req dto.AlarmManagePageRequest
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

	alarmManages, total, err := repo.GetAlarmManagePage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	response := dto.AlarmManagePageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetAlarmManageByReportStatus 根据报告状态获取报警列表
func (a *AlarmManageApi) GetAlarmManageByReportStatus(c *gin.Context) {
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManages, err := repo.GetAlarmManageByReportStatus(status)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmManageByAlarmType 根据报警类型获取报警列表
func (a *AlarmManageApi) GetAlarmManageByAlarmType(c *gin.Context) {
	alarmType := c.Param("type")
	if alarmType == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManages, err := repo.GetAlarmManageByAlarmType(alarmType)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetAlarmManageByMediaID 根据媒体ID获取报警列表
func (a *AlarmManageApi) GetAlarmManageByMediaID(c *gin.Context) {
	mediaIDStr := c.Param("mediaId")
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	alarmManages, err := repo.GetAlarmManageByMediaID(mediaID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// QueryAlarmManage 根据查询条件获取报警列表
func (a *AlarmManageApi) QueryAlarmManage(c *gin.Context) {
	var req dto.AlarmManageQueryRequest
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

	alarmManages, total, err := repo.GetAlarmManageByQuery(
		req.AlarmType,
		req.ReportStatus,
		req.MediaID,
		req.StartTime,
		req.EndTime,
		req.Page,
		req.PageSize,
	)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.AlarmManageResponse
	for _, alarmManage := range alarmManages {
		responses = append(responses, dto.AlarmManageResponse{
			ID:           alarmManage.ID,
			AlarmTask:    alarmManage.AlarmTask,
			VideoMedia:   alarmManage.VideoMedia,
			MediaURL:     alarmManage.MediaURL,
			Img:          alarmManage.Img,
			InitImg:      alarmManage.InitImg,
			AlarmType:    alarmManage.AlarmType,
			ReportStatus: alarmManage.ReportStatus,
			ReportAddr:   alarmManage.ReportAddr,
			AlarmVideo:   alarmManage.AlarmVideo,
			AlarmDetail:  alarmManage.AlarmDetail,
			Tm:           alarmManage.Tm,
			MediaID:      alarmManage.MediaID,
		})
	}

	response := dto.AlarmManagePageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}
