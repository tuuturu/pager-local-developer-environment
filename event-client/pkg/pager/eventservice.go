package pager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Event struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func CreateEvent(baseURL, token, title, description string) error {
	url := fmt.Sprintf("%s/events", baseURL)
	event := Event{
		Title:       title,
		Description: description,
	}
	
	payload, err := json.Marshal(event)
	if err != nil {
	    return fmt.Errorf("error marshalling payload: %w", err)
	}
	
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
	    return fmt.Errorf("error creating event request: %w", err)
	}
	
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	
	client := http.Client{}
	
	response, err := client.Do(request)
	if err != nil {
	    return fmt.Errorf("error posting event: %w", err)
	}
	
	if response.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("response returned %d", response.StatusCode))
	}
	
	return nil
}

func GetEvents(baseURL, token string) error {
	url := fmt.Sprintf("%s/events", baseURL)
	
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating get events request: %w", err)
	}
	
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	
	client := http.Client{}
	
	response, err := client.Do(request)
	if err != nil {
	    return fmt.Errorf("error doing GET events request: %w", err)
	}
	
	if response.StatusCode != http.StatusOK {
		return errors.New("bad response code")
	}
	
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    return fmt.Errorf("error reading response body: %w", err)
	}
	
	println(string(payload))
	
	return nil
}
