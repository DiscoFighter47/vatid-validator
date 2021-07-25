// Code generated by gowsdl DO NOT EDIT.

package euvies

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/DiscoFighter47/vatid-validator/config"
	"github.com/hooklift/gowsdl/soap"
)

type CheckVatReq struct {
	XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
	CountryCode string   `xml:"countryCode,omitempty" json:"countryCode,omitempty"`
	VatNumber   string   `xml:"vatNumber,omitempty" json:"vatNumber,omitempty"`
}

type CheckVatResp struct {
	XMLName     xml.Name     `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`
	CountryCode string       `xml:"countryCode,omitempty" json:"countryCode,omitempty"`
	VatNumber   string       `xml:"vatNumber,omitempty" json:"vatNumber,omitempty"`
	RequestDate soap.XSDDate `xml:"requestDate,omitempty" json:"requestDate,omitempty"`
	Valid       bool         `xml:"valid,omitempty" json:"valid,omitempty"`
	Name        *string      `xml:"name,omitempty" json:"name,omitempty"`
	Address     *string      `xml:"address,omitempty" json:"address,omitempty"`
}

type Client interface {
	CheckVat(request *CheckVatReq) (*CheckVatResp, error)
	CheckVatContext(ctx context.Context, request *CheckVatReq) (*CheckVatResp, error)
}

type client struct {
	soapClient *soap.Client
}

func NewClient(cnf config.Euvies) Client {
	return &client{soap.NewClient("http://ec.europa.eu/taxation_customs/vies/services/checkVatService",
		soap.WithRequestTimeout(time.Second*time.Duration(cnf.Timeout))),
	}
}

func (client *client) CheckVatContext(ctx context.Context, request *CheckVatReq) (*CheckVatResp, error) {
	response := new(CheckVatResp)
	err := client.soapClient.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *client) CheckVat(request *CheckVatReq) (*CheckVatResp, error) {
	return client.CheckVatContext(
		context.Background(),
		request,
	)
}
