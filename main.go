package main

import (
	"fmt"
	"github.com/DuriDuri/parsley/api"
)

func main() {
	fmt.Println("Starting Server")

	a, err := api.Init()
	if err != nil {
		fmt.Print("Error on boot: ", err.Error())
	}

	server := a.GetServer()
	panic(server.Run(":" + a.Port))
}
