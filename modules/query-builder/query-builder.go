package querybuilder

import "gorm.io/datatypes"

type TestEndpointRequest struct {
	InputData datatypes.JSON `json:"input_data"`
}
