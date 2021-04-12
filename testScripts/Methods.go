package main

import "fmt"

type Camera struct {
	name        string
	focalLength int
}

func (t_cam *Camera) doubleFocal() int {
	return 2 * t_cam.focalLength
}

func main() {
	o_cam := Camera{name: "EVS", focalLength: 8}
	fmt.Println("Double focal: ", o_cam.doubleFocal())
}
