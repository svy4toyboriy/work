package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/svy4toyboriy/work/config"
	v1 "github.com/svy4toyboriy/work/internal/controller/http/v1"
	"github.com/svy4toyboriy/work/internal/usecase"
	"github.com/svy4toyboriy/work/internal/usecase/repo"
	"github.com/svy4toyboriy/work/pkg/logger"
	"github.com/svy4toyboriy/work/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	r := repo.New(pg)
	accountUseCase := usecase.New(r)

	// HTTP Server
	handler := gin.Default()
	v1.NewRouter(handler, l, *accountUseCase)

	handler.Run()
}
