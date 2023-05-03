package controller

import "go.uber.org/fx"

var Module = fx.Invoke(
	InitTableController,
	InitWorkspaceController,
	InitApiController,
	InitEndpointController,
	InitFunctionController,
)
