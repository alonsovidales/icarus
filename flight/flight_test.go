package flight

import (
	"fmt"
	"testing"
	"time"
)

func TestTrottle(t *testing.T) {
	fc := New(17, 27, 22, 18, 23, 2000, 1000, "/dev/pigpio", true)

	fmt.Println("Increasing trottle")
	for i := 0; i <= 20; i++ {
		fc.SetTrottle(i)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Decreasing trottle to 5%")
	for i := 20; i >= 5; i-- {
		fc.SetTrottle(i)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("Moving front 50%")
	fc.Front(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Moving back 50%")
	fc.Back(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Turning left 50%")
	fc.TurnLeft(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Turning right 50%")
	fc.TurnRight(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Moving right 50%")
	fc.MoveRight(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Moving left 50%")
	fc.MoveLeft(50)
	time.Sleep(2 * time.Second)
	fc.Still()

	fmt.Println("Off")
	fc.SetTrottle(0)
}
