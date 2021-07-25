package api

import (
	"fmt"
	"log"

	"github.com/DiscoFighter47/vatid-validator/config"
	"github.com/DiscoFighter47/vatid-validator/euvies"
	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
)

func StartServer(cnf config.App, euvies euvies.Client) {
	vcAPI := &vatCheckAPI{euvies: euvies}

	log.Println("server starting on port:", cnf.Port)
	graceful.Run(fmt.Sprintf(":%d", cnf.Port), 0, apiHandler(vcAPI))
}

func apiHandler(vatCheckAPI *vatCheckAPI) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	v1 := r.Group("/api/v1")
	{
		vcHandler := v1.Group("/vatcheck")
		{
			vcHandler.GET("/:vatid", vatCheckAPI.check)
		}
	}
	return r
}
