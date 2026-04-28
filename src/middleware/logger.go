package middleware

import (
	"bytes"
	"geji/util"
	"time"

	"github.com/gin-gonic/gin"
)

// responseWriter 自定义响应写入器，用于捕获响应内容
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 包装响应写入器以捕获响应内容
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Next()

		duration := time.Since(start)

		// 记录请求信息
		util.Logi("[%s] %s %s %v", c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			duration)

		// 记录响应内容（限制长度避免日志过大）
		responseBody := writer.body.String()
		if len(responseBody) > 0 {
			// 限制响应日志长度，最多记录2000字符
			maxLen := 1024
			if len(responseBody) > maxLen {
				responseBody = responseBody[:maxLen] + "... [truncated]"
			}
			util.Logi("[RESPONSE] %s %s | Status: %d | Body: %s",
				c.Request.Method,
				c.Request.URL.Path,
				c.Writer.Status(),
				responseBody)
		}
	}
}
