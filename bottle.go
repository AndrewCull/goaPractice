package main

import (
	"cellar/app"
	"fmt"
	"github.com/goadesign/goa"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show implements the "show" action of the "bottles" controller.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	if ctx.BottleID == 0 {
		// Emulate a missing record with ID 0
		return ctx.NotFound()
	}
	// Build the resource using the generated data structure
	bottle := app.GoaExampleBottle{
		ID:   ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}

	// Let the generated code produce the HTTP response using the
	// media type described in the design (BottleMedia).
	return ctx.OK(&bottle)
}

// BottleController_Show: start_implement

// Put your logic here

//BottleController_Show: end_implement
//	res := &app.GoaExampleBottle{}
//	return ctx.OK(res)
//}
