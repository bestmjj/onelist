/*
 * @Time    : 2022年03月23日 19:19:10
 * @Author  : root
 * @Project : kido
 * @File    : cors.go
 * @Software: GoLand
 * @Describe:
 */
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 设置Logger
func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return logger
}

func CORSMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		// 处理请求
		c.Next()

		start := time.Now()
		// 计算请求耗时
		latency := time.Since(start)

		// 获取请求和响应信息
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()

		// 根据状态码设置日志级别
		var level logrus.Level
		switch {
		case statusCode >= 500:
			level = logrus.ErrorLevel
		case statusCode >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		// 记录日志
		logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"method":    reqMethod,
			"uri":       reqUri,
			"latency":   latency,
			"client_ip": c.ClientIP(),
		}).Log(level, "Request processed")

	}
}
