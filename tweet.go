// post_tweet.go
package main

import (
	"bytes"
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

// postTweet sends a tweet to Twitter
func postTweet(content string) (string, error) {
	// OAuth1 authentication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create tweet content
	tweetData := map[string]interface{}{
		"text": content,
	}
	tweetJSON, err := json.Marshal(tweetData)
	if err != nil {
		return "", fmt.Errorf("error marshaling tweet content: %w", err)
	}

	// Make POST request to v2 tweets endpoint
	response, err := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetJSON))
	if err != nil {
		return "", fmt.Errorf("failed to post tweet: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode == http.StatusCreated {
		var result map[string]interface{}
		body, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %w", err)
		}
		// Extract tweet ID from the response
		return result["data"].(map[string]interface{})["id"].(string), nil
	}

	// Print detailed error information
	body, _ := ioutil.ReadAll(response.Body)
	return "", fmt.Errorf("failed to post tweet: %s, response: %s", response.Status, string(body))
}

func main() {
	// Example of posting a tweet
	tweetID, err := postTweet("Welcome to Jashneek's Tweet!")
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}
	fmt.Printf("Posted tweet with ID: %s\n", tweetID)
}
