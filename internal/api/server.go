package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	cmn "ng_autocode_back_service/common/model"
	moduleGrp "ng_autocode_back_service/internal/core"
	config "ng_autocode_back_service/pkg/config"
	"ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

var (
	srv *http.Server
	ctx context.Context
)

func SetAPIServer(a_ctx *context.Context, modulesGrp moduleGrp.ModuleLayer, appType cmn.APP_TYPE, buildMode int8) {
	ctx = *a_ctx
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().GetInt("server.port")),
		Handler: SetupRoutes(modulesGrp, logger.Log(), appType, buildMode),
	}
	logger.Log().Info("Server starting on port")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log().Error("Failed to start API server", zap.Error(err))
		}
	}()
}

func ShutdownRouter() {
	timeout := config.GetConfig().GetDuration("server.timeout")
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	logger.Log().Info("Shutting down server...")
	defer logger.Log().Info("Server shutdown complete")

	if err := srv.Shutdown(timeoutCtx); err != nil {
		logger.Log().Error("Server forced to shutdown", zap.Error(err))
	} else {
		logger.Log().Info("Server exited gracefully")
	}
	select {
	case <-timeoutCtx.Done():
		log.Println("timeout of ", timeout, " exceeded")
	}
}
