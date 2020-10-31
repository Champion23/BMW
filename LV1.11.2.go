package main
import (
	"fmt"
)
type  Video  struct {
	Theme  string
	Duration  float64
	Viewcounts  int

}
func  (p Video) ProntInfo(){
	fmt.Printf("Theme: %v \n   Duration : %v\n    Viewcounts : %v ", p.Theme  , p.Duration, p.Viewcounts)
}
func main (){
	var p1 = Video{
		Theme: "Los Angeles Lakers",
		Duration: 12.6,
		Viewcounts: 185600,
	}
	p1.ProntInfo()
}