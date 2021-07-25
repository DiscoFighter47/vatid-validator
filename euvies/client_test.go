package euvies_test

import (
	"testing"

	"github.com/DiscoFighter47/vatid-validator/config"
	"github.com/DiscoFighter47/vatid-validator/euvies"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckVat(t *testing.T) {
	client := euvies.NewClient(config.Euvies{Timeout: 5})
	testData := []struct {
		des         string
		countryCode string
		vatNo       string
		err         bool
		valid       bool
	}{
		{
			des:         "valid vat id",
			countryCode: "DE",
			vatNo:       "266182271",
			err:         false,
			valid:       true,
		},
		{
			des:         "invalid vat id",
			countryCode: "DE",
			vatNo:       "000000000",
			err:         false,
			valid:       false,
		},
		{
			des:         "invalid country code",
			countryCode: "BD",
			vatNo:       "266182271",
			err:         true,
			valid:       false,
		},
	}

	for _, td := range testData {
		Convey(td.des, t, func() {
			resp, err := client.CheckVat(&euvies.CheckVatReq{
				CountryCode: td.countryCode,
				VatNumber:   td.vatNo,
			})
			if td.err {
				Convey("should be an error", func() {
					So(err, ShouldBeError)
				})
			} else {
				Convey("error should be nil", func() {
					So(err, ShouldBeNil)
					Convey("validity should match", func() {
						So(resp.Valid, ShouldEqual, td.valid)
					})
				})
			}
		})
	}
}
