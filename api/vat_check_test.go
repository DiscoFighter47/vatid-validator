package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DiscoFighter47/vatid-validator/euvies"
	. "github.com/smartystreets/goconvey/convey"
)

type euviesStub struct{}

func (stub *euviesStub) CheckVatContext(ctx context.Context, request *euvies.CheckVatReq) (*euvies.CheckVatResp, error) {
	return nil, nil
}

func (stub *euviesStub) CheckVat(request *euvies.CheckVatReq) (*euvies.CheckVatResp, error) {
	log.Println(request.VatNumber)
	if request.VatNumber == "999999999" {
		log.Println("returning error...")
		return nil, errors.New("euvies stub error")
	}
	name, addr := "name", "addr"
	return &euvies.CheckVatResp{
		CountryCode: "country_code",
		VatNumber:   "vat_no",
		Valid:       true,
		Name:        &name,
		Address:     &addr,
	}, nil
}

func TestVatCheck(t *testing.T) {
	api := apiHandler(&vatCheckAPI{&euviesStub{}})
	testData := []struct {
		des  string
		path string
		code int
		res  string
	}{
		{
			des:  "successful call",
			path: "/api/v1/vatcheck/DE266182271",
			code: http.StatusOK,
			res:  `{"data":{"addr":"addr","country_code":"country_code","name":"name","valid":true,"vat_no":"vat_no"}}`,
		},
		{
			des:  "invalid vat id",
			path: "/api/v1/vatcheck/BD266182271",
			code: http.StatusUnprocessableEntity,
			res:  `{"error":"invalid request: VatID: regular expression mismatch"}`,
		},
		{
			des:  "euvies client error",
			path: "/api/v1/vatcheck/DE999999999",
			code: http.StatusInternalServerError,
			res:  `{"error":"unable to check vat info: euvies stub error"}`,
		},
	}

	for _, td := range testData {
		Convey(td.des, t, func() {
			res := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, td.path, nil)
			api.ServeHTTP(res, req)
			Convey(fmt.Sprintf("status should be %d", td.code), func() {
				So(res.Code, ShouldEqual, td.code)
				Convey("body should have proper json", func() {
					So(string(res.Body.Bytes()), ShouldEqual, td.res)
				})
			})
		})
	}
}
