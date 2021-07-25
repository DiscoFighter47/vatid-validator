package main

import (
	"github.com/DiscoFighter47/vatid-validator/api"
	"github.com/DiscoFighter47/vatid-validator/euvies"
)

func main() {
	api.StartServer(euvies.NewClient())
}
