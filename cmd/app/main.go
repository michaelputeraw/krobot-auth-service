package main

import (
	"github.com/michaelputeraw/krobot-auth-service/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.AppModule).Run()
}
