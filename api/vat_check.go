package api

import (
	"net/http"

	"github.com/DiscoFighter47/vatid-validator/euvies"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type vatCheckAPI struct {
	euvies euvies.Client
}

type checkReq struct {
	VatID string `uri:"vatid" binding:"required" validate:"regexp=^(DE)?[0-9]{9}$"`
}

func (api *vatCheckAPI) check(c *gin.Context) {
	req := &checkReq{}
	if err := c.ShouldBindUri(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}
	if err := validator.Validate(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	resp, err := api.euvies.CheckVat(&euvies.CheckVatReq{
		CountryCode: req.VatID[:2],
		VatNumber:   req.VatID[2:],
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to check vat info: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"country_code": resp.CountryCode,
			"vat_no":       resp.VatNumber,
			"valid":        resp.Valid,
			"name":         resp.Name,
			"addr":         resp.Address,
		},
	})
}
