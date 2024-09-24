package STRequests

import (
	"SpaceTraders/models"
	"fmt"
	"net/http"
)

// Validation info for requests
var validation = models.Validation{
	ContentType:   "application/json",
	Authorization: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiVU5ERUFEVE9LRU5BUlQiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMDktMDEiLCJpYXQiOjE3MjY5NjExMjIsInN1YiI6ImFnZW50LXRva2VuIn0.p1VBvHoZYad_DX8qWTlqliqJb0K8e7y_jrJ8iZ0URMFgPrCP40TzP1Ma2qoPzkPrd_7E5u92ofKM9_kuLMAquAnaAnoQ4e4ddEKol8AIhYbzrLVOStL_nVID70Zi-m9DbHVRhyVFWKtOGu0Serw360WnLY14NfQK-WnYx-GO6xPUaICb9HknQHgANTSqLgOH7acphX-qGwc18nAQ71W06r4Qwf4Way-v8zkU4OCEVhqM8e0OMt_U6rG0m_4ofSwdUDKzVl0WtL1NhzywoFx3qRGSSjkROBXfkuHF1MiSHTUSvMKMJafwJIIRYTvbeE_IR-5FSmT5tV3i7RTm5aF5Bw",
}

func GetResponse(route string) (*http.Response, error) {
	// Set the request
	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	request.Header.Set("Authorization", validation.Authorization)

	// Send the request
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	return response, err
}
