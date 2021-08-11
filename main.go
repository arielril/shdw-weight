package main

import "github.com/arielril/attack-graph-flow-weight-api/api/http/router"

func main() {
	rt := router.GetRouter()
	rt.Run(":6000")
}
