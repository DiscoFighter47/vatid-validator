package main

import (
	"github.com/DiscoFighter47/vatid-validator/api"
	"github.com/DiscoFighter47/vatid-validator/config"
	"github.com/DiscoFighter47/vatid-validator/euvies"
)

func main() {
	cnf := config.GetConfig()
	euvies := euvies.NewClient(cnf.Euvies)
	api.StartServer(cnf.App, euvies)
}
