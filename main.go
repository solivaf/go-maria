package main

import (
	"fmt"
	"github.com/solivaf/go-maria/internal/app"
	"github.com/solivaf/go-maria/internal/app/config"
	"os"
)

func main() {
	cfg := config.InitConfig()
	if err := app.ExecuteCommand(*cfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
