package main

import (
	"time"

	"infraction-tracker/db_impl"
	"infraction-tracker/endpoints"
	"infraction-tracker/env"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func bindApiRoutes(router *gin.Engine, sharedContext endpoints.SharedContext) {
	api := router.Group("/api")
	{
		api.GET("/list", sharedContext.ApiGetList)
	}

	authRoutes := api.Group("/")
	authRoutes.Use(sharedContext.Auth)
	{
		authRoutes.POST("/add/:name", sharedContext.ApiPostAddPunishee)
		authRoutes.POST("/update/:name/:count", sharedContext.ApiPostUpdatePunishee)
		authRoutes.POST("/remove/:name", sharedContext.ApiPostRemovePunishee)
	}
}

func bindPageRoutes(router *gin.Engine, sharedContext endpoints.SharedContext) {
	router.LoadHTMLGlob("templates/*")

	router.GET("/", sharedContext.PageGetIndex)
	router.GET("/admin", sharedContext.PageGetAdmin)
}

func bindStaticFiles(router *gin.Engine) {
	router.Static("/static", "./public")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.Domain()},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-API-KEY"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

func main() {
	env.EnvironmentVariablesSpecified()

	db := &db_impl.MemoryDb{}
	db.Init()

	sharedContext := endpoints.SharedContext{
		Db: db,
	}

	router := initRouter()
	bindApiRoutes(router, sharedContext)
	bindPageRoutes(router, sharedContext)
	bindStaticFiles(router)

	router.Run()
}
