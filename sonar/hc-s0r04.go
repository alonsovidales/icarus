package sonar

import (
	"errors"
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const samplePeriod = 100 * time.Millisecond

const MAX_WAIT_ITERS = 10000

type HcSr04 struct {
	trigger rpio.Pin
	echo    rpio.Pin
}

func NewHcSr04(triggerPort, echoPort int) (hc *HcSr04) {
	fmt.Println("Trigger:", triggerPort, "Echo:", echoPort)
	if err := rpio.Open(); err != nil {
		panic("Error while opening GPIO port")
	}
	hc = &HcSr04{
		trigger: rpio.Pin(triggerPort),
		echo:    rpio.Pin(echoPort),
	}

	hc.trigger.Output()
	hc.trigger.Low()
	hc.echo.Input()

	return hc
}

func (hc *HcSr04) GetDistanceCm() (int, error) {
	hc.trigger.High()
	time.Sleep(10 * time.Microsecond)
	hc.trigger.Low()
	pulseStart := time.Now()
	for hc.echo.Read() == rpio.Low {
		pulseStart = time.Now()
	}

	pulseEnd := time.Now()
	for hc.echo.Read() == rpio.High {
		pulseEnd = time.Now()
	}

	distCm := float64(pulseEnd.UnixNano()-pulseStart.UnixNano()) / 1000000000 * 17150

	if distCm > 1000 {
		return -1, errors.New("InvalidRead")
	}

	return int(distCm), nil
}

func (hc *HcSr04) AddTrigger(minCm int, trigger func(dist int)) {
	c := time.Tick(samplePeriod)
	for _ = range c {
		dist, err := hc.GetDistanceCm()
		if err == nil && dist < minCm {
			trigger(dist)
		}
	}
}
