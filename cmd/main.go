package main

import (
	"os"
	"strconv"

	"github.com/s0h1s2/invoice-app/cmd/api"
)

func main() {
	var addr int = 8080
	if port, ok := os.LookupEnv("PORT"); ok {
		addr, _ = strconv.Atoi(port)
	}
	eng := api.NewEngine(addr)
	eng.Start()
}
