package main

import (
	"syscall/js"

	"github.com/hellgrenj/bf-tsm/pkg/routes"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("route", js.FuncOf(route))
	<-c
}

func route(this js.Value, inputs []js.Value) interface{} {
	points := []routes.Point{}
	for i := range inputs {
		points = append(points, routes.Point{X: inputs[i].Get("x").Float(), Y: inputs[i].Get("y").Float(), Label: inputs[i].Get("label").String()})
	}
	optimal := routes.OptimalPath(points)
	jsObj := map[string]interface{}{}
	for _, p := range optimal.Points {
		jsObj[p.Label] = map[string]interface{}{"x": p.X, "y": p.Y}
	}
	return js.ValueOf(jsObj)
}
