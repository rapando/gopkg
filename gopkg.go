package gopkg

import (
	"fmt"

	"github.com/rapando/gopkg/shape"
)

// Describe -  Runs all the functions of the interface
func Describe(s shape.Shape) {
	s.Description()
	fmt.Printf("the area is %v\n", s.Area())
	fmt.Printf("the volume is %v\n", s.Volume())
}
