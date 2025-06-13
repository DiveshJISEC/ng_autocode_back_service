//go:build debug

package main

import (
	"fmt"
	cmn "ng_autocode_back_service/common/model"
	startup "ng_autocode_back_service/internal/startup"
	"os"
)

func main() {
	fmt.Println("Hello World - Fixed deposit")
	// Initialize the application
	if err := startup.Startup(cmn.BUILD_DEBUG, cmn.APP_FIXED_DEPOSIT); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %v\n", err)
		os.Exit(1)
	}
}
