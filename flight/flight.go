package flight

import (
	"fmt"
	"os"
)

type FlightControl struct {
	pigpioPipe  *os.File
	maxPulseLen int
	minPulseLen int
	trottlePort int
	rollPort    int
	pitchPort   int
	yawPort     int
	debug       bool
}

func New(trottlePort, rollPort, pitchPort, yawPort, maxPulseLen, minPulseLen int, pigpioPipePath string, debug bool) *FlightControl {
	pigpioPipe, err := os.OpenFile(pigpioPipePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(`Error trying to open the pigpio pipe,
			please install it: http://abyz.co.uk/rpi/pigpio/
			and make sure that it is running. Path:`, pigpioPipePath)
		panic(err)
	}

	fc := &FlightControl{
		pigpioPipe:  pigpioPipe,
		debug:       debug,
		trottlePort: trottlePort,
		rollPort:    rollPort,
		pitchPort:   pitchPort,
		yawPort:     yawPort,
		maxPulseLen: maxPulseLen,
		minPulseLen: minPulseLen,
	}

	return fc
}

func (fc *FlightControl) sendPercentage(port, perc int) {
	toSend := fmt.Sprintf(
		"s %d %d\n",
		port,
		(((fc.maxPulseLen-fc.minPulseLen)/100)*perc)+fc.minPulseLen)
	if fc.debug {
		fmt.Println("Sending perc:", perc, "Str:", toSend)
	}
	_, err := fc.pigpioPipe.Write([]byte(fmt.Sprintf(toSend)))
	if err != nil {
		fmt.Println("Error sending PWM signal:", err)
	}
}

func (fc *FlightControl) SetTrottle(perc int) {
	fc.sendPercentage(fc.trottlePort, perc)
}
