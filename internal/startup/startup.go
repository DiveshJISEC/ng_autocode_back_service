package startup

import (
	"fmt"
	cmn "ng_autocode_back_service/common/model"
	"os"
)

// Startup is the entry function for initializing the application.
func Startup(releaseMode int8, appType cmn.APP_TYPE) error {
	fmt.Println("Startup called with releaseMode:", releaseMode, "appType:", appType, "name:")
	fmt.Println("Args", os.Args)

	//ctx := context.Background()

	//#region 1.Config
	// Load configuration based on release mode and app type
	fmt.Println("1. Loading Config...")
	envr := os.Args[1]
	if envr == "" {
		envr = "dev"
	}
	//#endcofig

	//#region 2. Logger
	// Initialize logger
	fmt.Println("2. Loading Logger...")
	//#endregion

	//#region 3. Repository
	// Initialize repository
	fmt.Println("3. Loading Repository...")
	//#endregion

	//#region 4. Service
	// Initialize service
	fmt.Println("4. Loading Service modules...")
	//#endregion

	//#region 5. Api Server
	// Initialize API server
	fmt.Println("5. Loading Api Server...")
	//#endregion

	fmt.Println("Startup completed successfully. Server ready to serve requests.")
	return nil
}
