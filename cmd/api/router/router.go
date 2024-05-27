package router

import (
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/report"
	"github.com/gin-gonic/gin"
)

// Used for local implementation and testing

type router struct {
	reportUC report.UseCase
}

type HandlerInterface interface {
	RegisterRoutes(r *gin.Engine)
}

func HandlerFactory(config *configuration.Config) HandlerInterface {
	return NewHandler(config)
}

func NewHandler(config *configuration.Config) HandlerInterface {
	r := &router{
		reportUC: report.NewUseCaseFactory(config),
	}
	return r
}

func (rh *router) RegisterRoutes(r *gin.Engine) {
}
