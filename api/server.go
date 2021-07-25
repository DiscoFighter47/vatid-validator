package api

import (
	"github.com/DiscoFighter47/vatid-validator/euvies"
	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
)

func StartServer(euvies euvies.Client) {
	vcAPI := &vatCheckAPI{euvies: euvies}
	graceful.Run(":8080", 0, apiHandler(vcAPI))
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
