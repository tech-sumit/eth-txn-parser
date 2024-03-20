package pkg

import (
	"errors"
	"eth-indexer/pkg/controllers"
	"eth-indexer/pkg/parser"
	"github.com/gin-gonic/gin"

	_ "eth-indexer/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)

func RunGinServer(parser parser.Parser) *http.Server {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ct := controllers.Setup(parser)
	// Endpoint to get the current block
	r.GET("/currentBlock", ct.GetCurrentBlock)

	// Endpoint to subscribe to an address
	r.POST("/subscribe", ct.Subscribe)

	// Endpoint to get transactions for an address
	r.GET("/transactions/:address", ct.GetTransactions)

	srv := &http.Server{
		Addr: func() string {
			if port, exists := os.LookupEnv("PORT"); exists {
				return "0.0.0.0:" + port
			} else {
				return "0.0.0.0:8080"
			}
		}(),
		Handler: r,
	}

	go func() {
		// Serve with the HTTP server
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
