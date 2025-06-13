package main

import (
	"context"
	"fmt"
	"log"
	"ng_autocode_back_service/api"
	repo "ng_autocode_back_service/common/infra/repository"
	cmn "ng_autocode_back_service/common/model"
	cfg "ng_autocode_back_service/common/pkg/config"
	logger "ng_autocode_back_service/common/pkg/logger"
	cmnUtils "ng_autocode_back_service/common/util"
	moduleGrp "ng_autocode_back_service/internal"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go.uber.org/zap"
)

// Startup is the entry function for initializing the application.
func Startup(releaseMode int8, appType cmn.APP_TYPE) error {
	fmt.Println("Startup called with releaseMode:", releaseMode, "appType:", appType, "name:")
	fmt.Println("Args", os.Args)

	ctx := context.Background()

	//#region 1.Config - Load configuration based on release mode and app type
	fmt.Println("1. Loading Config...")
	envr := os.Args[1]
	if envr == "" {
		envr = "dev"
	}
	config := cfg.LoadConfig(appType, envr)
	if config == nil {
		log.Fatal("Failed to load config")
		return fmt.Errorf("failed to load config")
	}
	//#endcofig

	//#region 2. Logger - Initialize logger
	fmt.Println("2. Loading Logger...")
	logLevel, err := strconv.Atoi(config.GetString("log.level"))
	if err != nil {
		log.Fatal("Error converting log level:", err)
	}
	sLogFilePath := config.GetString("log.path")
	sLogFilePath = cmnUtils.UpdatePathSeparator(sLogFilePath)
	cmnUtils.CraeteFileWithPath(sLogFilePath)
	logger.LoggerInit(sLogFilePath, logger.GetZapLevel(logLevel))
	//#endregion

	//#region 3. Repository - Initialize repository
	fmt.Println("3. Loading Repository...")
	repoOb, err := repo.CreateNewAppRepo(ctx)
	if err != nil {
		logger.Log(ctx).Error("Failed to create new app repository:", zap.Error(err))
		return fmt.Errorf("failed to create new app repository: %w", err)
	}
	//#endregion

	//#region 4. Service - Initialize service
	fmt.Println("4. Loading Service modules...")
	modulesLayer := moduleGrp.NewAppModuleLayer(repoOb)
	//#endregion

	//#region 5. Api Server - Initialize API server
	fmt.Println("5. Loading Api Server...")
	api.SetAPIServer(&ctx, modulesLayer, appType, releaseMode)
	//#endregion

	sAddr := fmt.Sprintf("%s:%s", config.GetString("server.host"), config.GetString("server.port"))
	fmt.Println("Server is starting at:", sAddr)

	fmt.Println("Startup completed successfully. Server ready to serve requests.")
	addShutdownHook()
	return nil
}

func addShutdownHook() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	logger.Log().Info("Quit signal received. Shutting down server...")
	api.ShutdownRouter()
	//api.CloseDBConnections()
	logger.Log().Info("Server shutdown completed.")
}
