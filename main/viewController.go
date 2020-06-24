package main

type View struct {
	data []byte
}
type Controller struct {
	input View
	processor func(View) View
	run func(Controller) Controller
	output View
}

func NewController(input View,processor func(View) View,run func(Controller) Controller, output View) *Controller {

	return &Controller{
		input:input,
		processor: processor,
		run: run,
		output: output,
	}
}


