package main

import (
	"fmt"
	"github.com/tuuturu/event-client/pkg/oauth2"
	"github.com/tuuturu/event-client/pkg/pager"
	"log"
	"net/url"
	"os"
)

func getToken() string {
	discoveryURL, err := url.Parse(os.Getenv("DISCOVERY_URL"))
	if err != nil {
		log.Fatal(err)
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	token, err := oauth2.AcquireToken(*discoveryURL, clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func main() {
	var err error
	
	token := getToken()
	if token == "" {
		log.Fatal("no token found")
	}

	baseURL := "http://localhost:3000"
	
	switch "multiple" {
	case "add":
		err := pager.CreateEvent(baseURL, token, "Slack: new mention", "@John Doe mentioned you in \"@Julius, kan du ta denne?\"")
		if err != nil {
			panic(err)
		}
	case "getAll":
		err = pager.GetEvents(baseURL, token)
		if err != nil {
			panic(err)
		}
	case "multiple":
		err = createTestEvents(baseURL, token)
		if err != nil {
			panic(err)
		}
	}
}

func createTestEvents(baseURL, token string) error {
	events := []struct {
		Title string
		Description string
	}{
	    {
	    	Title: "New email",
	    	Description: "Message received from person@hotmail.com",
	    },
	    {
	    	Title: "Github: new issue",
	    	Description: "@johnd created issue: terraform should know what env to use after using CURRENT_ENV",
		},
		{
	    	Title: "Slack: new mention",
	    	Description: "@John Doe mentioned you in \"@Julius, kan du ta denne?\"",
		},
		{
	    	Title: "Bank: new deposit",
	    	Description: "You just received 400kr with the message: \"Utlegg for middag\"",
		},
		{
	    	Title: "Calendar: upcoming meeting",
	    	Description: "Meeting about \"Fredagslunsj\" in 15 minutes",
		},
	}
	
	for _, event := range events {
		event := event
		
		err := pager.CreateEvent(baseURL, token, event.Title, event.Description)
		if err != nil {
		    return fmt.Errorf("error creating event: %w", err)
		}
	}
	
	return nil
}
