package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ok": "ok",
		})
	})

	r.GET("/rahasia", func(ctx *gin.Context) {
		logrus.Info(os.Getenv("RAHASIA"))
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ok": "ok",
		})
	})

	port := "9090"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	logrus.Infof("server running at :%s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
