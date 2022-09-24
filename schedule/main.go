package main

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/bootstrap"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
