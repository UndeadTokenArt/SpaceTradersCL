package handlers

import (
	"SpaceTraders/constants"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func SetUserName() map[string]string {
	userdetails := make(map[string]string)

	fmt.Println("pick a name between 4-14 characters")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}

	// check if input is viable
	if isValidLength(input) {
		// set the users chosen name
		userdetails["symbol"] = strings.TrimSpace(input)

	} else {
		fmt.Println("The string must be between 4 and 14 characters.")
	}
	// prompt for faction choice
	fmt.Println("What faction woudld you like join?")

	faction, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input", err)
		return nil
	}
	faction = strings.TrimSpace(faction)
	faction = strings.ToUpper(faction)
	userdetails["faction"] = faction
	return userdetails

}

func RegisterNewClient(userDetails map[string]string) {

	payload := map[string]string{
		"symbol":  userDetails["symbol"],
		"faction": userDetails["faction"],
	}

	// Marshal the data into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Set the request
	request, err := http.NewRequest("POST", constants.ApiRoutes["registerNew"], bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	request.Header.Set("Content-Type", constants.Validation.ContentType)

	// Send the request
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Handle the response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println(responseBody)
}
