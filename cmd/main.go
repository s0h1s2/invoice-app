package main

import "github.com/s0h1s2/invoice-app/cmd/api"

func main() {
	eng := api.NewEngine()
	eng.Start()
}
