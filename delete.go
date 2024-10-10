package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

// Twitter API credentials

const (
        consumerKey    = "LzJukk7dZkkbur41a6IG0uiGz"
        consumerSecret = "CIfSlRBVovdIuxrbWo6x4c9qsbSK396dth53oF2kV1e2aZgF6t"
        accessToken    = "1844361993286975489-mYTFCl3GD0MXpG8D1Md6EMZBap1YKx"
        accessSecret   = "AKKbvX8ZuU1pfU4MUhSpyLWDoANGOuvVRTlQyAU3aWkTW"
)

// deleteTweet deletes a tweet by its ID

func deleteTweet(tweetID string) error {
	// OAuth1 authentication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	// Execute the DELETE request
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute delete request: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode == http.StatusOK {
		// Read and unmarshal the response body
		var result map[string]interface{}
		body, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			return fmt.Errorf("error unmarshaling response: %w", err)
		}
		// Print success message
		if deleted, ok := result["data"].(map[string]interface{})["deleted"]; ok && deleted == true {
			fmt.Printf("Successfully deleted tweet with ID: %s\n", tweetID)
		} else {
			return fmt.Errorf("failed to confirm tweet deletion")
		}
		return nil
	}

	// Print detailed error information
	body, _ := ioutil.ReadAll(response.Body)
	return fmt.Errorf("failed to delete tweet: %s, response: %s", response.Status, string(body))
}

func main() {
	// Replace with the tweet ID you want to delete
	tweetID := "1844475032426172437" // Use the tweet ID from the postTweet function

	// Example of deleting the tweet
	err := deleteTweet(tweetID)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)
	}
}
