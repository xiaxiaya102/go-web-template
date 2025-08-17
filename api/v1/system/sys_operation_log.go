package system

import (
	"go-web-template/global"
	"go-web-template/pkg/dto"
	"go-web-template/pkg/handle"
	"go-web-template/pkg/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetOperationLogList 获取操作日志列表
// @Summary 获取操作日志列表
// @Description 分页获取操作日志列表，支持多条件筛选
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Param username query string false "用户名"
// @Param module query string false "操作模块"
// @Param operation query string false "操作类型"
// @Param ip query string false "操作IP"
// @Param startTime query string false "开始时间 格式:2006-01-02"
// @Param endTime query string false "结束时间 格式:2006-01-02"
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response{data=dto.PageResult{list=[]model.OperationLog}}
// @Router /api/operation-logs [get]
func GetOperationLogList(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 获取筛选参数
	username := c.Query("username")
	module := c.Query("module")
	operation := c.Query("operation")
	ip := c.Query("ip")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	// 构建查询
	db := global.DB.Model(&model.OperationLog{})

	// 添加筛选条件
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	if module != "" {
		db = db.Where("module = ?", module)
	}
	if operation != "" {
		db = db.Where("operation LIKE ?", "%"+operation+"%")
	}
	if ip != "" {
		db = db.Where("ip = ?", ip)
	}
	if startTime != "" {
		db = db.Where("operation_time >= ?", startTime+" 00:00:00")
	}
	if endTime != "" {
		db = db.Where("operation_time <= ?", endTime+" 23:59:59")
	}

	// 获取总数
	var total int64
	db.Count(&total)

	// 分页查询
	var logs []model.OperationLog
	offset := (page - 1) * pageSize
	if err := db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error; err != nil {
		c.JSON(http.StatusOK, handle.Error("查询操作日志失败"))
		return
	}

	// 返回分页结果
	result := dto.PageResult{
		List:     logs,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	c.JSON(http.StatusOK, handle.Success(result))
}

// GetOperationLogDetail 获取操作日志详情
// @Summary 获取操作日志详情
// @Description 根据ID获取操作日志详情
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param id path int true "日志ID"
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response{data=model.OperationLog}
// @Router /api/operation-logs/{id} [get]
func GetOperationLogDetail(c *gin.Context) {
	id := c.Param("id")

	var log model.OperationLog
	if err := global.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.Error("操作日志不存在"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(log))
}

// DeleteOperationLog 删除操作日志
// @Summary 删除操作日志
// @Description 根据ID删除操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param id path int true "日志ID"
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response
// @Router /api/operation-logs/{id} [delete]
func DeleteOperationLog(c *gin.Context) {
	id := c.Param("id")

	if err := global.DB.Delete(&model.OperationLog{}, id).Error; err != nil {
		c.JSON(http.StatusOK, handle.Error("删除操作日志失败"))
		return
	}

	c.JSON(http.StatusOK, handle.OkWithMsg("删除成功"))
}

// BatchDeleteOperationLogs 批量删除操作日志
// @Summary 批量删除操作日志
// @Description 根据ID列表批量删除操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param request body dto.BatchDeleteRequest true "删除请求"
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response
// @Router /api/operation-logs/batch-delete [post]
func BatchDeleteOperationLogs(c *gin.Context) {
	var req dto.BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.Error("参数错误"))
		return
	}

	if len(req.IDs) == 0 {
		c.JSON(http.StatusOK, handle.Error("请选择要删除的日志"))
		return
	}

	if err := global.DB.Delete(&model.OperationLog{}, req.IDs).Error; err != nil {
		c.JSON(http.StatusOK, handle.Error("批量删除操作日志失败"))
		return
	}

	c.JSON(http.StatusOK, handle.OkWithMsg("批量删除成功"))
}

// ClearOperationLogs 清空操作日志
// @Summary 清空操作日志
// @Description 清空指定时间范围的操作日志
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param request body dto.ClearLogsRequest true "清空请求"
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response
// @Router /api/operation-logs/clear [post]
func ClearOperationLogs(c *gin.Context) {
	var req dto.ClearLogsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.Error("参数错误"))
		return
	}

	db := global.DB.Model(&model.OperationLog{})

	// 如果指定了时间范围
	if req.StartTime != "" && req.EndTime != "" {
		db = db.Where("operation_time BETWEEN ? AND ?", req.StartTime+" 00:00:00", req.EndTime+" 23:59:59")
	} else if req.Days > 0 {
		// 清空指定天数之前的日志
		cutoffTime := time.Now().AddDate(0, 0, -req.Days).Format("2006-01-02 15:04:05")
		db = db.Where("operation_time < ?", cutoffTime)
	}

	if err := db.Delete(&model.OperationLog{}).Error; err != nil {
		c.JSON(http.StatusOK, handle.Error("清空操作日志失败"))
		return
	}

	c.JSON(http.StatusOK, handle.OkWithMsg("清空操作日志成功"))
}

// GetOperationLogStats 获取操作日志统计
// @Summary 获取操作日志统计
// @Description 获取操作日志的统计信息
// @Tags 操作日志管理
// @Accept json
// @Produce json
// @Param days query int false "统计天数" default(7)
// @Security ApiKeyAuth
// @Success 200 {object} handle.Response{data=dto.OperationLogStats}
// @Router /api/operation-logs/stats [get]
func GetOperationLogStats(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "7"))
	startTime := time.Now().AddDate(0, 0, -days).Format("2006-01-02 00:00:00")

	// 总操作次数
	var totalCount int64
	global.DB.Model(&model.OperationLog{}).Where("operation_time >= ?", startTime).Count(&totalCount)

	// 按模块统计
	var moduleStats []dto.ModuleStats
	global.DB.Model(&model.OperationLog{}).
		Select("module, COUNT(*) as count").
		Where("operation_time >= ?", startTime).
		Group("module").
		Order("count DESC").
		Scan(&moduleStats)

	// 按操作类型统计
	var operationStats []dto.OperationStats
	global.DB.Model(&model.OperationLog{}).
		Select("operation, COUNT(*) as count").
		Where("operation_time >= ?", startTime).
		Group("operation").
		Order("count DESC").
		Scan(&operationStats)

	// 按用户统计
	var userStats []dto.UserStats
	global.DB.Model(&model.OperationLog{}).
		Select("username, COUNT(*) as count").
		Where("operation_time >= ? AND username != 'anonymous'", startTime).
		Group("username").
		Order("count DESC").
		Limit(10).
		Scan(&userStats)

	// 按日期统计
	var dailyStats []dto.DailyStats
	global.DB.Model(&model.OperationLog{}).
		Select("DATE(operation_time) as date, COUNT(*) as count").
		Where("operation_time >= ?", startTime).
		Group("DATE(operation_time)").
		Order("date").
		Scan(&dailyStats)

	stats := dto.OperationLogStats{
		TotalCount:     totalCount,
		ModuleStats:    moduleStats,
		OperationStats: operationStats,
		UserStats:      userStats,
		DailyStats:     dailyStats,
	}

	c.JSON(http.StatusOK, handle.Success(stats))
}
