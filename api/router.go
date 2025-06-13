package api

import (
	"bytes"
	"io"
	cmn "ng_autocode_back_service/common/model"
	config "ng_autocode_back_service/common/pkg/config"
	moduleGrp "ng_autocode_back_service/internal"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(moduleGrp moduleGrp.ModuleLayer, logger *zap.Logger, appType cmn.APP_TYPE, buildMode int8) *gin.Engine {
	// Initialize the router

	if cmn.BUILD_DEBUG == buildMode {
		gin.SetMode(gin.DebugMode)
	} else if cmn.BUILD_RELEASE == buildMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.TestMode)
	}
	router := gin.New()

	// Set up middleware
	router.Use(cors())
	router.Use(customLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", moduleGrp.GetSystemMatrics)
	//Sample
	router.GET("/version", moduleGrp.GetAppVersion)
	router.GET("/readfile", moduleGrp.GetReadTextData)
	router.GET("/readjson", moduleGrp.GetReadJsonData)
	//Sample

	if appType == cmn.APP_FIXED_DEPOSIT {
		fd := router.Group("/fd")

		v1 := fd.Group("/v1")

		list := v1.Group("/list")
		{
			list.POST("/agent", moduleGrp.GetFDAgentList)
			//list.GET("/agent/:id", moduleGrp.GetFDAgentById)
		}
		book := v1.Group("/book")
		{
			book.POST("/orderbook", moduleGrp.GetFDOrderBook)
			//book.GET("/order/:id", moduleGrp.GetOrderBookById)
		}
		// misc := v1.Group("/misc")
		// {
		// 	book.GET("/TransactionCount", moduleGrp.GetTransactionCount)
		// }
	}
	return router
}

func customLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body string
		)

		start := time.Now()

		contentType := c.Request.Header.Get("Content-Type")
		excludeBody := strings.HasPrefix(contentType, "multipart/form-data")
		if !excludeBody {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				bodyCopy := bytes.NewBuffer(bodyBytes)
				c.Request.Body = io.NopCloser(bodyCopy)
				body = string(bodyBytes)
			}
		}
		c.Next()

		if c.FullPath() == "/health" {

			// Log the request
			logger.Info("Call ended",
				zap.String("path", c.Request.URL.Path),
				zap.String("requestId", c.GetString(config.REQUESTID)),
				zap.String("UCC", c.GetString(config.UCC)),
				zap.String("userId", c.GetString(config.USERID)),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("query", c.Request.URL.RawQuery),
				zap.String("body", body),
				zap.String("remote_addr", c.Request.RemoteAddr),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", time.Since(start).Milliseconds()),
			)
		}
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Link")
		c.Writer.Header().Set("Access-Control-Max-Age", "300")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
