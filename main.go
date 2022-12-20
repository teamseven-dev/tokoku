package main

import (
	"fmt"
	"tokoku/config"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)

	fmt.Println(conn)
}
