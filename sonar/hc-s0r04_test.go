package sonar

import (
	"fmt"
	"testing"
)

func TestSonar(t *testing.T) {
	hc_front := NewHcSr04(21, 13)
	//hc_left := NewHcSr04(21, 19)
	//hc_right := NewHcSr04(21, 26)
	//hc_down := NewHcSr04(21, 5)
	//hc_back := NewHcSr04(21, 6)
	for true {
		dist, err := hc_front.GetDistanceCm()
		fmt.Println("Dist front:", dist, err)
		//fmt.Println("Dist left:", hc_left.GetDistanceCm())
		//fmt.Println("Dist right:", hc_right.GetDistanceCm())
		//dist, err := hc_down.GetDistanceCm()
		//fmt.Println("Dist down:", dist, err)
		//dist, err := hc_back.GetDistanceCm()
		//fmt.Println("Dist back:", dist, err)
		//time.Sleep(100 * time.Millisecond)
	}
}
