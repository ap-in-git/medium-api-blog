package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const facebookInvalidResponse = `{
  "error": {
    "message": "Invalid OAuth access token - Cannot parse access token",
    "type": "OAuthException",
    "code": 190,
    "fbtrace_id": "ACW2fFAhOy3xDgp1YVf3qG9"
  }
}`

const facebookProfileSuccessResponse = `{
  "id": "1937012216362664",
  "name": "Ashok Poudel Ap"
}`

func TestFacebookApiService_GetUserDetails_InvalidResponse(t *testing.T) {
	service := NewFacebookService()
	//Mock the server
	invalidServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(facebookInvalidResponse))
	}))
	//Change the base url from facebook api to the server url
	service.BaseUrl = invalidServer.URL + "/"

	_, clientError, _ := service.GetUserDetails()
	assert.Equal(t, http.StatusUnauthorized, clientError.StatusCode)
	assert.Equal(t, clientError.Error.Message, "Invalid OAuth access token - Cannot parse access token")

}

func TestFacebookApiService_GetUserDetails_SuccessResponse(t *testing.T) {
	service := NewFacebookService()

	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(facebookProfileSuccessResponse))
	}))
	//Change the base url from facebook api to the server url
	service.BaseUrl = successServer.URL + "/"

	profileData, clientError, err := service.GetUserDetails()
	assert.Nil(t, clientError)
	assert.Nil(t, err)
	assert.Equal(t, profileData.Name, "Ashok Poudel Ap")

}
