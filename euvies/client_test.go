package euvies_test

import (
	"testing"

	"github.com/DiscoFighter47/vatid-validator/euvies"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckVat(t *testing.T) {
	client := euvies.NewClient()

	Convey("valid vat id", t, func() {
		resp, err := client.CheckVat(&euvies.CheckVatReq{
			CountryCode: "DE",
			VatNumber:   "266182271",
		})
		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
			Convey("valid should be true", func() {
				So(resp.Valid, ShouldBeTrue)
			})
		})
	})

	Convey("invalid vat id", t, func() {
		resp, err := client.CheckVat(&euvies.CheckVatReq{
			CountryCode: "DE",
			VatNumber:   "000000000",
		})
		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
			Convey("valid should be false", func() {
				So(resp.Valid, ShouldBeFalse)
			})
		})
	})

	Convey("invalid country id", t, func() {
		_, err := client.CheckVat(&euvies.CheckVatReq{
			CountryCode: "BD",
			VatNumber:   "266182271",
		})
		Convey("should be an error", func() {
			So(err, ShouldBeError)
		})
	})
}
