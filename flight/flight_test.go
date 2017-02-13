package flight

import (
	"testing"
	"time"
)

func TestTrottle(t *testing.T) {
	fc := New(17, 27, 22, 18, 2000, 1000, "/dev/pigpio", true)
	for true {
		for i := 0; i <= 100; i++ {
			fc.SetTrottle(i)
			time.Sleep(100 * time.Millisecond)
		}
		for i := 100; i >= 0; i-- {
			fc.SetTrottle(i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
