package http

import (
	"context"
	"os"

	"github.com/dodirepository/user-svc/infrastructure/database"
	"github.com/dodirepository/user-svc/infrastructure/http"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

// Start :nodoc:
func Start(ctx context.Context) {
	httpServer := http.NewHTTPServer()
	defer httpServer.Done()
	_, err := database.DBInit()
	if err != nil {
		logrus.Fatal("Failed connected to database")
	}

	logrus.Infof("http server user service start on port %s", os.Getenv("APP_PORT"))
	if err := httpServer.Run(ctx); err != nil {
		logrus.Info("http server stopped")
	}
}
