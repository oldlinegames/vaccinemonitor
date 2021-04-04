package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/oldlinegames/vaccinemonitor/structs"
)

func StartEngine(inp structs.MonitorInput) {
	apptAvailCache := make(map[string]bool)

	url := fmt.Sprintf("https://api.prod.projectexodus.us/get-locations/v2/%s?radius=%s", inp.ZipCode, inp.Radius)
	method := "GET"

	tr := http.Transport{}

	headerMap := make(map[string][]string)

	// Not sure if this API cares about Chrome headers but I'll copy em anyways :)
	headerMap["authority"] = []string{"api.prod.projectexodus.us"}
	headerMap["sec-ch-ua"] = []string{"\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\""}
	headerMap["accept"] = []string{"application/json, text/plain, */*"}
	headerMap["sec-ch-ua-mobile"] = []string{"?0"}
	headerMap["user-agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"}
	headerMap["origin"] = []string{"https://www.goodrx.com"}
	headerMap["sec-fetch-site"] = []string{"cross-site"}
	headerMap["sec-fetch-mode"] = []string{"cors"}
	headerMap["sec-fetch-dest"] = []string{"https://www.goodrx.com/"}
	headerMap["accept-language"] = []string{"en-US,en;q=0.9"}

	for {
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header = headerMap

		res, err := tr.RoundTrip(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var jsonMap map[string]interface{}

		json.Unmarshal(body, &jsonMap)

		// wanted to cast straight to [](map[string]interace{}) but not allowed :(
		locationArr := jsonMap["results"].([]interface{})

		for _, entry := range locationArr {
			entryMap := entry.(map[string]interface{})
			if entryMap["appointments"] != nil {
				// Check if appts were not previously availible (i.e. have been restocked)
				avail, ok := apptAvailCache[entryMap["address1"].(string)]
				if ok {
					if !avail { // vaccine has restocked, time to notify!
						apptAvailCache[entryMap["address1"].(string)] = true
						// Format information and print to console
						fmt.Println("Vaccine Appointment Availible!")
						loc := fmt.Sprintf("%s at %s %s %s %s", entryMap["locationName"].(string), entryMap["address1"].(string), entryMap["city"].(string), entryMap["stateAbbr"].(string), entryMap["zipcode"].(string))
						fmt.Printf("Location: %s\n", loc)

						dist := entryMap["miles"].(string)
						fmt.Printf("Distance: %s miles\n", dist)

						url := entryMap["webAddress"].(string)
						fmt.Printf("URL: %s\n", url)

						brandAvailibility := entryMap["medNames"].(map[string]bool)
						for brand, availibility := range brandAvailibility {
							if availibility {
								fmt.Printf("%s: available", brand)
							} else {
								fmt.Printf("%s: available", brand)
							}
						}

						fmt.Println()

						// Send to webhook if it was set
						if inp.Webhook != "" {
							SendWebhook(inp.Webhook, url, loc, dist, brandAvailibility)
						}
					}
				} else {
					apptAvailCache[entryMap["address1"].(string)] = true
					// Format information and print to console
					fmt.Println("Vaccine Appointment Availible!")
					loc := fmt.Sprintf("%s at %s %s %s %s", entryMap["locationName"].(string), entryMap["address1"].(string), entryMap["city"].(string), entryMap["stateAbbr"].(string), entryMap["zipcode"].(string))
					fmt.Printf("Location: %s\n", loc)

					dist := entryMap["miles"].(string)
					fmt.Printf("Distance: %s miles\n", dist)

					url := entryMap["webAddress"].(string)
					fmt.Printf("URL: %s\n", url)

					brandAvailibility := entryMap["medNames"].(map[string]interface{})
					boolMap := make(map[string]bool)
					for brand, availibility := range brandAvailibility {
						fmt.Println(brand)
						boolMap[brand] = availibility.(bool)
						if availibility.(bool) {
							fmt.Printf("%s: available\n", brand)
						} else {
							fmt.Printf("%s: unavailable\n", brand)
						}
					}

					fmt.Println()

					// Send to webhook if it was set
					if inp.Webhook != "" {
						SendWebhook(inp.Webhook, url, loc, dist, boolMap)
					}
				}
			} else {
				// Update our info so we know what's in stock
				apptAvailCache[entryMap["address1"].(string)] = false
			}
		}

		time.Sleep(10 * time.Second)
	}

}
