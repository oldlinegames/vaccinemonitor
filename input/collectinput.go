package input

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oldlinegames/vaccinemonitor/structs"
)

func CollectInput() structs.MonitorInput {
	inp := structs.MonitorInput{}

	fmt.Println("Do you wish to load settings? (y/n)")
	load := AcceptInput()
	if load == "y" {
		if _, err := os.Stat("./settings.json"); err == nil {
			data, err := ioutil.ReadFile("./settings.json")

			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(data, &inp)

			return inp
		} else if os.IsNotExist(err) {
			fmt.Println("No settings found!")
			fmt.Println("Enter Zip Code: ")
			zip := AcceptInput()

			fmt.Println("Enter Search Radius: ")
			radius := AcceptInput()

			fmt.Println("Enter Discord Webhook (leave blank if you don't want notifications): ")
			webhook := AcceptInput()

			inp.ZipCode = zip
			inp.Radius = radius
			inp.Webhook = webhook

			WriteSettings(inp)

			return inp
		} else {
			fmt.Println("Enter Zip Code: ")
			zip := AcceptInput()

			fmt.Println("Enter Search Radius: ")
			radius := AcceptInput()

			fmt.Println("Enter Discord Webhook (leave blank if you don't want notifications): ")
			webhook := AcceptInput()

			inp.ZipCode = zip
			inp.Radius = radius
			inp.Webhook = webhook

			WriteSettings(inp)

			return inp
		}
	} else {
		fmt.Println("Enter Zip Code: ")
		zip := AcceptInput()

		fmt.Println("Enter Search Radius: ")
		radius := AcceptInput()

		fmt.Println("Enter Discord Webhook (leave blank if you don't want notifications): ")
		webhook := AcceptInput()

		inp.ZipCode = zip
		inp.Radius = radius
		inp.Webhook = webhook

		WriteSettings(inp)

		return inp
	}
}

func WriteSettings(i structs.MonitorInput) {
	payload, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./settings.json", payload, 0644)

	if err != nil {
		panic(err)
	}
}
