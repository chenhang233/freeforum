package main

import (
	"fmt"
	"freeforum/config"
)

func main() {
	config.SetupConfig("freeforum.conf")

	fmt.Println("hello freeforum")
}
