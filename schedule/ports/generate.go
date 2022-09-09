package ports

import (
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=generate-config.yaml ../../api/ScheduleService.yaml
