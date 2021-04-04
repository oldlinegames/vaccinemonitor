package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// This function is hideous becasue I needed to fit the Discord formatting specs
// which don't play well with Go's JSON compatibility. It gets the job done though!
func SendWebhook(url string, vaxLink string, location string, distance string, brands map[string]bool) {
	payload := make(map[string]interface{})

	payload["username"] = "Covid-19 Vaccine Monitor"
	payload["avatar_url"] = "https://timesofindia.indiatimes.com/thumb/msid-78496670,imgsize-195084,width-400,resizemode-4/78496670.jpg"

	embeds := make([](map[string]interface{}), 1)
	embed := make(map[string]interface{})

	embed["title"] = "Covid-19 Vaccine Appointments Availible!"
	embed["url"] = vaxLink
	embed["description"] = fmt.Sprintf("A vaccine appointment is now availible just %s miles away! Click the link in the title to register for your appointment!", distance)

	fields := make([](map[string]interface{}), 4)

	locField := make(map[string]interface{})
	locField["name"] = "Location"
	locField["value"] = location
	locField["inline"] = false

	jjField := make(map[string]interface{})
	jjField["name"] = "Johnson and Johnson"
	if jjAvail, ok := brands["pfizer"]; ok && jjAvail {
		jjField["value"] = "Availible"
	} else {
		jjField["value"] = "Unavailable"
	}
	jjField["inline"] = true

	pfizerField := make(map[string]interface{})
	pfizerField["name"] = "Pfizer"
	if pfizerAvail, ok := brands["pfizer"]; ok && pfizerAvail {
		pfizerField["value"] = "Availible"
	} else {
		pfizerField["value"] = "Unavailable"
	}
	pfizerField["inline"] = true

	modernaField := make(map[string]interface{})
	modernaField["name"] = "Moderna"
	if modernaAvail, ok := brands["pfizer"]; ok && modernaAvail {
		modernaField["value"] = "Availible"
	} else {
		modernaField["value"] = "Unavailable"
	}
	modernaField["inline"] = true

	fields[0] = locField
	fields[1] = jjField
	fields[2] = pfizerField
	fields[3] = modernaField

	embed["fields"] = fields

	embeds[0] = embed

	payload["embeds"] = embeds

	body, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Error posting to webhook!")
	}
}
