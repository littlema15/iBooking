package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	port = 9910
)

func main() {
	router := gin.Default()
	routes.RegisterBookingRoutes(router)
	fmt.Printf("Starting listening port:%d\n", port)

	srv := &http.Server{
		Addr:    "localhost:" + strconv.Itoa(port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
