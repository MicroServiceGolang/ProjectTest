package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"practise/test_task_post/api/docs"
	"practise/test_task_post/config"
	"practise/test_task_post/grpc/server"
	"practise/test_task_post/pkg/logger"
	"practise/test_task_post/rest"
	"practise/test_task_post/services/srv"
	"practise/test_task_post/storage/store"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	_ = godotenv.Load()

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	var cfg config.Config
	if err := envconfig.ProcessWith(context.TODO(), &cfg, envconfig.OsLookuper()); err != nil {
		panic(err)
	}
	// Set Host for swagger. In case we need to add swagger to WSO2API Manager
	docs.SwaggerInfo.Host = cfg.ServerIP + cfg.HTTPPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	log := logger.New(cfg.LogLevel, cfg.Environment)

	db, err := sql.Open("postgres", cfg.PostgresURI())
	if err != nil {
		panic(err)
	}
	log.Info("Database connection established")

	service := srv.New(cfg, log)
	log.Info("Connected to service...")

	store := store.New(db, log)
	log.Info("Connected to storage...")

	r := gin.New()

	// *** gRPC Initialization
	listener, errOnListen := net.Listen("tcp", cfg.GRPCPort)
	if errOnListen != nil {
		log.Error(errOnListen.Error())
	}

	rpcServer := server.StartServer(cfg, log, service, store)

	go func() {
		if errOnServer := rpcServer.Serve(listener); errOnServer != nil {
			log.Error(errOnServer.Error())
		}
	}()
	log.Info("main: gRPC Server started at port " + cfg.GRPCPort)

	// *** REST Initialization
	server := rest.New(
		store,
		service,
		log, cfg,
		r,
	)

	restserver := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: server,
	}

	// Graceful HTTP server shutdown context
	go func() {
		if err := restserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	log.Info(fmt.Sprintf("Server started at %v", cfg.HTTPPort))

	const valtime = 5
	OSCall := <-quitSignal
	ctx, cancel := context.WithTimeout(context.Background(), valtime*time.Second)
	defer cancel()
	fmt.Printf("system call:%+v", OSCall)
	if err := restserver.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("REST Server Graceful Shutdown Failed: %s\n", err))
	}
	log.Info("REST Server Gracefully Shut Down")

}
