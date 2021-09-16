package domain

import "github.com/Code0716/go-vtm/app/gen/api"

// CommonSuccessResponse has a text message.
type CommonSuccessResponse api.CommonSuccessResponse

// Pager struct
type Pager struct {
	// limit params
	Limit int

	// offset param
	Offset int

	// status param
	Status string
}
