package http

import (
	"context"
	"os"

	"github.com/dodirepository/user-svc/infrastructure/http"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

// Start :nodoc:

//	func Start(ctx context.Context) {
//		services := bootstrap.NewService(ctx)
//		defer services.DisconnectAllConnection(ctx)
//		startHttpServer(ctx, services)
//	}
func Start(ctx context.Context) {
	httpServer := http.NewHTTPServer()
	defer httpServer.Done()

	logrus.Infof("http server start on port %s", os.Getenv("APP_PORT"))
	if err := httpServer.Run(ctx); err != nil {
		logrus.Info("http server stopped")
	}
}
