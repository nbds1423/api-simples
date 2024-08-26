package src

import (
	"fmt"
	"log"
	"net/http"
	"personal-api/src/middleware"
	"personal-api/src/routes"
	"personal-api/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {

	server := &http.Server{
		Addr:    utils.Env.Port,
		Handler: Router(),
	}

	fmt.Printf("Server online! using port %s \n", strings.Replace(utils.Env.Port, ":", "", 1))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//Middleware - Global
	r.Use(middleware.Cors(), middleware.RateLimit())

	//No Group - Public
	r.NoRoute(routes.Err)
	{
		r.Any("/swagger/*any", middleware.ValidateMethod("GET"), ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Any("/status", middleware.ValidateMethod("GET"), routes.Status)
	}

	//Group - Private
	router := r.Group("/api")
	router.Use(middleware.Headers(), middleware.ValidateUserAgent(), middleware.AuthenticateToken())
	{
		router.Any("/", middleware.ValidateMethod("GET"), routes.Index)
	}

	return r
}
