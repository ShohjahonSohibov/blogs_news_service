package main

import (
	"context"
	"flag"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"

	"news_blogs_service/api"
	"news_blogs_service/api/handlers"
	"news_blogs_service/config"
	"news_blogs_service/pkg/logger"
	"news_blogs_service/storage"
	"news_blogs_service/storage/postgres"
)
func main() {
	var (
		cfg = config.Load()
		log = initLogger(cfg)
		db  storage.StorageI
	)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	// Assign the ConcreteStorage instance to the db variable
	db = pgStore

	paramHandler := &handlers.ParamHandler{
		Cfg: cfg,
		Log: log,
		Db:  db,
	}
	h := handlers.NewHandler(paramHandler)

	paramApi := &api.ParamSetUpAPI{
		Gin:     gin.New(),
		Cfg:     cfg,
		Handler: h,
	}
	runApi(paramApi)
}

func runApi(param *api.ParamSetUpAPI) {
	var (
		rps = *flag.Int("rps", param.Cfg.RateLimit, "request per second")
	)

	param.Limit = ratelimit.New(rps)

	api.SetUpRouter(param)

	if err := param.Gin.Run(param.Cfg.HTTPPort); err != nil {
		return
	}
}

func initLogger(cfg *config.Config) logger.LoggerI {
	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	return log
}
