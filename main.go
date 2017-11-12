package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

var drone *minidrone.Driver

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone = minidrone.NewDriver(bleAdaptor)

	work := func() {
		drone.On(minidrone.Battery, func(data interface{}) {
			fmt.Printf("battery: %d\n", data)
		})

		drone.On(minidrone.FlightStatus, func(data interface{}) {
			fmt.Printf("flight status: %d\n", data)
		})

		drone.On(minidrone.Takeoff, func(data interface{}) {
			fmt.Println("taking off...")
		})

		drone.On(minidrone.Hovering, func(data interface{}) {
			fmt.Println("hovering!")
			/*
				gobot.After(5*time.Second, func() {
					drone.Land()
				})
			*/
		})

		drone.On(minidrone.Landing, func(data interface{}) {
			fmt.Println("landing...")
		})

		drone.On(minidrone.Landed, func(data interface{}) {
			fmt.Println("landed.")
		})

		//	time.Sleep(1000 * time.Millisecond)
		//	drone.TakeOff()
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	setupServer()

	robot.Start()
}

// Message is the json message passed to the server
type Message struct {
	Command string
	Value   string
}

func setupServer() {
	log.Println("Starting Control Server")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)

		if err != nil {
			log.Println(err)
		}

		switch message.Command {
		case "LAUNCH":
			log.Println("Launching")
			drone.TakeOff()

		case "LAND":
			log.Println("Landing")
			drone.Land()
		}
	})

	go http.ListenAndServe(":8080", nil)
}
