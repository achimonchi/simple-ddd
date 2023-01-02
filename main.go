package main

import "poc-ddd/application/infra"

func main() {
	server, err := infra.NewInfraFactory().CreateInfra(infra.INFRA_Http, ":4444")
	if err != nil {
		panic(err)
	}

	server.Run()
}
