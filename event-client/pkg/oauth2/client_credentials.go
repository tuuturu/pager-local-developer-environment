package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	GrantTypeClientCredentials = "client_credentials"
)

type DiscoveryDocument struct {
	TokenEndpoint string `json:"token_endpoint"`
}

type Token struct {
	Value string `json:"access_token"`
}

func fetchDiscoveryDocument(wellKnownURL string) (discoveryDocument DiscoveryDocument, err error) {
	response, err := http.Get(wellKnownURL)
	if err != nil {
	    return discoveryDocument, fmt.Errorf("error fetching discovery document: %w", err)
	}

	rawDocument, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    return discoveryDocument, fmt.Errorf("error reading well-known URL body: %w", err)
	}
	
	err = json.Unmarshal(rawDocument, &discoveryDocument)
	if err != nil {
	    return discoveryDocument, fmt.Errorf("error parsing well-known json: %w", err)
	}
	
	return discoveryDocument, nil
}

func generateGrantRequest(tokenURL, clientID, clientSecret string) (request *http.Request, err error) {
	values := url.Values{}
	values.Set("grant_type", GrantTypeClientCredentials)
	values.Set("client_id", clientID)
	values.Set("client_secret", clientSecret)
	
	request, err = http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(values.Encode()))
	if err != nil {
	    return nil, fmt.Errorf("error creating grant request: %w", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	return request, nil
}

func AcquireToken(discoveryUrl url.URL, clientId, clientSecret string) (token string, err error) {
	discoveryDocument, err := fetchDiscoveryDocument(discoveryUrl.String())
	if err != nil {
	    return "", err
	}
	
	request, err := generateGrantRequest(discoveryDocument.TokenEndpoint, clientId, clientSecret)
	if err != nil {
	    return "", err
	}
	
	httpClient := http.Client{}

	response, err := httpClient.Do(request)
	if err != nil {
	    return "", fmt.Errorf("error posting to token endpoint: %w", err)
	}
	
	defer func() {
		_ = response.Body.Close()
	}()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    return "", fmt.Errorf("error extracting body: %w", err)
	}
	
	var tokenResponse Token
	
	err = json.Unmarshal(result, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Value, nil
}
