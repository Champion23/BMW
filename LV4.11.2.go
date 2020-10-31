package main

import (
	"fmt"
)
type Painter interface {

}
type Heart struct {

}
func Paint (painter Painter){
	for y :=3.0; y >-3.0; y-=0.1 {
		for x :=-2.0; x<2.0; x+=0.1 {
			if (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y<0 {
				fmt.Print("$")
			}else if (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y>0 {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}
func main (){
	var heart Heart
	Paint (heart )
}