package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Message struct {
	Subject          string `json:"subject"`
	From             Sender `json:"from"`
	ReceivedDateTime string `json:"receivedDateTime"`
	Body             struct {
		Content string `json:"content"`
	} `json:"body"`
}

type Sender struct {
	EmailAddress struct {
		Address string `json:"address"`
	} `json:"emailAddress"`
}

func main() {
	// アプリケーション情報
	clientID := ""
	clientSecret := "e3ce034f-97f7-4708-91c3-43fd2fd35050"
	tenantID := "f8cdef31-a31e-4b4a-93e4-5f571e91255a"

	accessToken, err := GetAccessToken(clientID, clientSecret, tenantID)
	if err != nil {
		log.Fatal(err)
	}

	mailboxMessages, err := GetMailboxMessages(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range mailboxMessages {
		fmt.Println("Subject:", message.Subject)
		fmt.Println("From:", message.From.EmailAddress.Address)
		fmt.Println("Received:", message.ReceivedDateTime)
		fmt.Println("Body:", message.Body.Content)
		fmt.Println("-----------------------------------------")
	}
}

func GetAccessToken(clientID, clientSecret, tenantID string) (string, error) {
	// https://login.microsoftonline.com/{tenant}/oauth2/v2.0/token
	tokenEndpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "https://graph.microsoft.com/.default")

	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to retrieve access token. Status code: %d", resp.StatusCode)
	}

	var tokenResp TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}

func GetMailboxMessages(accessToken string) ([]Message, error) {
	requestURL := "https://graph.microsoft.com/v1.0/me/mailfolders/inbox/messages?$select=subject,from,receivedDateTime,body"
	client := &http.Client{}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to retrieve mailbox messages. Status code: %d", resp.StatusCode)
	}

	var messageResponse struct {
		Value []Message `json:"value"`
	}

	err = json.NewDecoder(resp.Body).Decode(&messageResponse)
	if err != nil {
		return nil, err
	}

	return messageResponse.Value, nil
}
