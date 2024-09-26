package main

import (
	"SpaceTraders/STRequests"
	"SpaceTraders/constants"
	"SpaceTraders/handlers"
	"SpaceTraders/models"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please choose an option:")
	fmt.Println("1. Register New Client")
	fmt.Println("2. View Contracts for Existing Client")
	fmt.Println("3. View My Agents info")
	fmt.Println("4. View factions")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	switch input {
	case "1":
		userData := handlers.SetUserName()
		handlers.RegisterNewClient(userData)
	case "2":
		viewMy(constants.ApiRoutes["contracts"])
	case "3":
		viewMy(constants.ApiRoutes["agent"])
	case "4":
		viewfaction(constants.ApiRoutes["factions"])
	default:
		fmt.Println("Invalid option. Please choose 1 or 2.")
	}
}

// View contracts for the existing client
func viewMy(route string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TOKEN")

	// Validation info for requests
	var Validation = models.Validation{
		Authorization: token,
	}
	// Set the request
	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	request.Header.Set("Authorization", Validation.Authorization)

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

	prettyjson, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		fmt.Println("error marshal Indent", err)
	}

	fmt.Println(string(prettyjson))
}

// viewing details of selected faction
func viewfaction(route string) {
	response, err := STRequests.GetResponse(route)
	if err != nil {
		fmt.Println("There was an error getting the response set")
	}
	defer response.Body.Close()

	// Handle the response
	var responseBody models.ResponseFactions
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("Pick a faction to see description:")

	factionNames := make(map[string]string)
	for _, faction := range responseBody.Data {
		fmt.Printf("%s\n", faction.Symbol)
		factionNames[faction.Symbol] = faction.Description
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)

	// if the faction matches one in the list return that factions description
	if description, exists := factionNames[input]; exists {
		fmt.Println(description)
	} else {
		fmt.Println("Faction does not exist")
	}

	defer response.Body.Close()
}
