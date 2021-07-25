package euvies_test

import (
	"testing"

	"github.com/DiscoFighter47/vatid-validator/euvies"
)

func TestCheckVat(t *testing.T) {
	client := euvies.NewClient()

	t.Log(client.CheckVat(&euvies.CheckVatReq{
		CountryCode: "DE",
		VatNumber:   "266182271",
	}))
}
