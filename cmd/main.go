package main

import (
	"context"
	"eth-indexer/pkg"
	"eth-indexer/pkg/parser/ethereum"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var initialBlock int
	if initialBlockStr, exists := os.LookupEnv("INITIAL_BLOCK"); exists {
		var err error
		initialBlock, err = strconv.Atoi(initialBlockStr)
		if err != nil {
			fmt.Printf("Invalid INITIAL_BLOCK value: %s; using default 0\n", err)
		}
	}

	ethParser := ethereum.NewParser(initialBlock)
	go ethParser.StartParsing()

	srv := pkg.RunGinServer(ethParser)

	// Set up channel to listen for SIGINT signals (CTRL+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a SIGINT is received.
	<-quit
	fmt.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exiting")
}
