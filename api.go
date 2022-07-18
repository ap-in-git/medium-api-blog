package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewFacebookService() *FacebookApiService {
	return &FacebookApiService{
		BaseUrl:     "https://graph.facebook.com/v13.0",
		AccessToken: "yourTokengi",
	}
}

type FacebookApiService struct {
	BaseUrl     string `json:"base_url"`     //The base url for the facebook api endpoint
	AccessToken string `json:"access_token"` //Access token required for the facebook api
}

type FacebookApiError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceId string `json:"fbtrace_id"`
	} `json:"error"`
	StatusCode int `json:"status_code"`
}

type FacebookProfileResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//GetUserDetails makes api calls to get the data from the api
//It returns pointer to  FacebookProfileResponse, FacebookApiError and error
func (service *FacebookApiService) GetUserDetails() (*FacebookProfileResponse, *FacebookApiError, error) {
	//Set api url
	apiUrl := service.BaseUrl + "/me?fields=id,name&access_token=" + service.AccessToken
	//Make api calls
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, nil, err
	}
	//Read the data
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	//For invalid status code
	if resp.StatusCode != http.StatusOK {
		var facebookApiError FacebookApiError
		err = json.Unmarshal(bodyBytes, &facebookApiError)
		//If json marshal fails
		if err != nil {
			return nil, nil, err
		}
		//Return the api error that we got from facebook
		return nil, &facebookApiError, err
	}
	var facebookProfileResponse FacebookProfileResponse
	err = json.Unmarshal(bodyBytes, &facebookProfileResponse)
	if err != nil {
		return nil, nil, err
	}

	//Return success response
	return &facebookProfileResponse, nil, nil
}
