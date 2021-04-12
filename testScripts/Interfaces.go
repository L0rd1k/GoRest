package main

import "fmt"

type Camera interface {
	resolution() int
	speed() int
}

type EVS struct {
	width   int
	height  int
	pixSize int
}

type IRZ struct {
	width   int
	height  int
	pixSize int
}

func (c_evs EVS) resolution() int {
	return c_evs.height * c_evs.width
}

func (c_evs EVS) speed() int {
	return c_evs.pixSize * 20
}

func (c_irz IRZ) resolution() int {
	return c_irz.height * c_irz.width
}

func (c_irz IRZ) speed() int {
	return c_irz.pixSize * 15
}

func startMeasure(camera Camera) {
	fmt.Println(camera)
	fmt.Println(camera.resolution())
	fmt.Println(camera.speed())
}

func main() {
	evs_cam := EVS{width: 1024, height: 768, pixSize: 120}
	irz_cam := IRZ{width: 768, height: 576, pixSize: 100}
	startMeasure(evs_cam)
	fmt.Println()
	startMeasure(irz_cam)
}
