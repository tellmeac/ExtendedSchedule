package main

import (
	"github.com/tellmeac/extended-schedule/userconfig/bootstrap"

	"go.uber.org/fx"
)

// @title        Extended Schedule
// @version      1.0
// @description  Service to work with personal schedule.
func main() {
	fx.New(bootstrap.Module).Run()
}
