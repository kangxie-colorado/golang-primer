package trafficlight

import (
	"strconv"
	"time"

	"github.com/enriquebris/goconcurrentqueue"
	"github.com/kangxie-colorado/golang-primer/messaging/lib"
	log "github.com/sirupsen/logrus"
)

type TrafficLightColor int

const (
	Undefined = iota
	Green
	Yellow
	Red
)

const YellowTimer int = 5

type TrafficLight struct {
	whichIsGreen int
	lightColors  [2]TrafficLightColor
	Timer        int    // how much time has elapsed in this state(color is just state)
	greenSecs    [2]int // how long to stay green

	InputQueue *goconcurrentqueue.FIFO
	Controller LightControl

	MySock         lib.SocketDescriptor
	ButtonListener func(*TrafficLight) // use a callback is good enough
}

func CreateATrafficLight(greenSecs [2]int, lightSocks [2]lib.SocketDescriptor, mySock lib.SocketDescriptor) *TrafficLight {
	var tfl = TrafficLight{}
	// some light has to start green, otherwise, it is deadlock...
	// or use a red timer and let the short red timer to fire off and become green... not a big deal
	// only implementation details
	tfl.whichIsGreen = 0 // let the first light to be green, first light? North-South?
	tfl.lightColors = [2]TrafficLightColor{Green, Red}
	tfl.Timer = 0
	tfl.greenSecs = greenSecs
	tfl.InputQueue = goconcurrentqueue.NewFIFO()
	tfl.Controller = CreateLightControl(lightSocks)

	tfl.ButtonListener = StartButtonListener
	tfl.MySock = mySock

	return &tfl
}

func (tfl *TrafficLight) StartTimer() {
	for {
		time.Sleep(200 * time.Millisecond)
		tfl.InputQueue.Enqueue("One Second Passed")
	}
}

func (tfl *TrafficLight) Start() {

	go tfl.ButtonListener(tfl)
	go tfl.StartTimer()

	// main loop of traffic ligh state machine
	for {
		theOtherLightNo := 0
		if tfl.whichIsGreen == 0 {
			theOtherLightNo = 1
		}

		msg, err := tfl.InputQueue.DequeueOrWaitForNextElement()
		if err != nil {
			log.Errorln("Error dequeue inbox", err.Error())
		}

		switch msg {
		case "Button Pressed":
			if tfl.lightColors[tfl.whichIsGreen] == Green && tfl.greenSecs[tfl.whichIsGreen] == 60 {
				if tfl.Timer > 30 {
					tfl.Controller.changeColor(tfl.whichIsGreen, Yellow, "5")
					tfl.lightColors[tfl.whichIsGreen] = Yellow
					tfl.Timer = 0
				} else {
					tfl.Timer += 30
				}

			}
		case "One Second Passed":
			// one more second spent in this state
			tfl.Timer += 1

			if tfl.lightColors[tfl.whichIsGreen] == Green && tfl.Timer > tfl.greenSecs[tfl.whichIsGreen] {
				log.Infof("Lights are %v, turning light%v from Green to Yellow", tfl.lightColors, tfl.whichIsGreen)
				tfl.Controller.changeColor(tfl.whichIsGreen, Yellow, "5")
				tfl.lightColors[tfl.whichIsGreen] = Yellow
				tfl.Timer = 0
			} else if tfl.lightColors[tfl.whichIsGreen] == Yellow && tfl.Timer > YellowTimer {

				log.Infof("Lights are %v, turning light%v from Yellow to Red", tfl.lightColors, tfl.whichIsGreen)
				tfl.Controller.changeColor(tfl.whichIsGreen, Red, "")
				tfl.lightColors[tfl.whichIsGreen] = Red

				// this should change another light to green, how to do that...
				// huh, actually there is not another state machine, one state machine controls two lights...
				// yeah, and shit, I entered a dead place when thinking I need two timers for green/red light length

				tfl.Controller.changeColor(theOtherLightNo, Green, strconv.Itoa(tfl.greenSecs[theOtherLightNo]))
				tfl.lightColors[theOtherLightNo] = Green
				tfl.whichIsGreen = theOtherLightNo

				tfl.Timer = 0
			} else {
				if tfl.lightColors[tfl.whichIsGreen] == Green {
					tfl.Controller.changeColor(tfl.whichIsGreen, Green, strconv.Itoa(tfl.greenSecs[tfl.whichIsGreen]-tfl.Timer))

				} else if tfl.lightColors[tfl.whichIsGreen] == Yellow {
					tfl.Controller.changeColor(tfl.whichIsGreen, Yellow, strconv.Itoa(YellowTimer-tfl.Timer))
				}
			}

		default:
			log.Errorln("Unknown input events:", msg)
		}
	}

}
