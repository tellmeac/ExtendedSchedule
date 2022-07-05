package main

import (
	"tellmeac/extended-schedule/bootstrap"

	"go.uber.org/fx"
)

// @title        Extended Schedule
// @version      1.0
// @description  Service to work with personal schedule.
func main() {
	fx.New(bootstrap.Module).Run()
}
