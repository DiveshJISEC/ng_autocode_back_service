//go:build debug

package main

import (
	"fmt"
	cmn "ng_autocode_back_service/common/model"
	"os"
)

func main() {
	fmt.Println("Hello World - Fixed deposit")
	// Initialize the application
	if err := Startup(cmn.BUILD_DEBUG, cmn.APP_FIXED_DEPOSIT); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %v\n", err)
		os.Exit(1)
	}
}
