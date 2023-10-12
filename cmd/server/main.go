package main

import "github.com/danielzinhors/APIS-go/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
