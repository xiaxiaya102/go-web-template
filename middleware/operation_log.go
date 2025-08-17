package middleware

import (
	"bytes"
	"encoding/json"
	"go-web-template/global"
	"go-web-template/logger"
	"go-web-template/pkg/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// OperationLogMiddleware 操作日志中间件
func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		startTime := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，以便后续处理器可以读取
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建响应写入器包装器
		responseWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = responseWriter

		// 处理请求
		c.Next()

		// 计算执行时间
		duration := time.Since(startTime)

		// 获取用户信息
		userID := getUserID(c)
		username := getUsername(c)

		// 获取客户端IP
		clientIP := getClientIP(c)

		// 创建操作日志
		operationLog := &model.OperationLog{
			UserID:        userID,
			Username:      username,
			Module:        getModule(c.Request.URL.Path),
			Operation:     getOperation(c.Request.Method, c.Request.URL.Path),
			Method:        c.Request.Method,
			URL:           c.Request.URL.String(),
			IP:            clientIP,
			UserAgent:     c.Request.UserAgent(),
			RequestBody:   filterSensitiveData(string(requestBody)),
			ResponseBody:  filterSensitiveData(responseWriter.body.String()),
			Status:        c.Writer.Status(),
			Duration:      duration.Milliseconds(),
			ErrorMessage:  getErrorMessage(c),
			OperationTime: startTime.Format("2006-01-02 15:04:05"),
			Description:   getOperationDescription(c.Request.Method, c.Request.URL.Path),
		}

		// 异步保存日志，避免影响响应性能
		go saveOperationLog(operationLog)
	}
}

// responseBodyWriter 响应体写入器包装器
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// getUserID 获取用户ID
func getUserID(c *gin.Context) uint {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(uint); ok {
			return id
		}
		// 尝试从字符串转换
		if idStr, ok := userID.(string); ok {
			if id, err := strconv.ParseUint(idStr, 10, 32); err == nil {
				return uint(id)
			}
		}
	}
	return 0 // 未登录用户
}

// getUsername 获取用户名
func getUsername(c *gin.Context) string {
	if username, exists := c.Get("username"); exists {
		if name, ok := username.(string); ok {
			return name
		}
	}
	return "anonymous" // 匿名用户
}

// getClientIP 获取客户端真实IP
func getClientIP(c *gin.Context) string {
	// 优先从X-Forwarded-For获取
	if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	// 从X-Real-IP获取
	if ip := c.GetHeader("X-Real-IP"); ip != "" {
		return ip
	}
	// 从RemoteAddr获取
	return c.ClientIP()
}

// getModule 根据URL路径获取模块名
func getModule(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) >= 2 {
		return parts[1] // 例如 /api/user -> user
	}
	return "system"
}

// getOperation 根据方法和路径获取操作类型
func getOperation(method, path string) string {
	switch method {
	case "GET":
		if strings.Contains(path, "/list") || strings.Contains(path, "/page") {
			return "查询列表"
		}
		return "查询详情"
	case "POST":
		if strings.Contains(path, "/login") {
			return "用户登录"
		}
		return "新增"
	case "PUT", "PATCH":
		return "修改"
	case "DELETE":
		return "删除"
	default:
		return method
	}
}

// getOperationDescription 获取操作描述
func getOperationDescription(method, path string) string {
	module := getModule(path)
	operation := getOperation(method, path)
	return operation + module + "数据"
}

// getErrorMessage 获取错误信息
func getErrorMessage(c *gin.Context) string {
	if c.Writer.Status() >= 400 {
		if err, exists := c.Get("error"); exists {
			if errStr, ok := err.(string); ok {
				return errStr
			}
		}
		return http.StatusText(c.Writer.Status())
	}
	return ""
}

// filterSensitiveData 过滤敏感数据
func filterSensitiveData(data string) string {
	if data == "" {
		return ""
	}

	// 限制长度，避免日志过大
	if len(data) > 2000 {
		data = data[:2000] + "..."
	}

	// 过滤密码等敏感字段
	sensitiveFields := []string{"password", "token", "secret", "key"}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err == nil {
		for _, field := range sensitiveFields {
			if _, exists := result[field]; exists {
				result[field] = "***"
			}
		}
		if filtered, err := json.Marshal(result); err == nil {
			return string(filtered)
		}
	}

	// 如果不是JSON格式，使用字符串替换
	for _, field := range sensitiveFields {
		data = strings.ReplaceAll(data, `"`+field+`":"`, `"`+field+`":"***`)
	}

	return data
}

// saveOperationLog 保存操作日志
func saveOperationLog(log *model.OperationLog) {
	if err := global.DB.Create(log).Error; err != nil {
		logger.Error("保存操作日志失败: %v", err)
	}
}
