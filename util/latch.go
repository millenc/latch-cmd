package util

import (
	"errors"
	"github.com/millenc/golatch"
	"net/http"
	"net/url"
)

//Initializes the latch object that will be used by all subcommands
func NewLatch(AppID string, SecretKey string, Proxy string, UseCallbacks bool, OnLatchRequestStart func(request *golatch.LatchRequest), OnLatchResponseReceive func(request *golatch.LatchRequest, response *http.Response, responseBody string)) (latch *golatch.Latch, err error) {
	if AppID == "" {
		err = errors.New("You must provide an Application's ID (--app).")
	}
	if err == nil && SecretKey == "" {
		err = errors.New("You must provide the secret key (--secret).")
	}

	if err == nil {
		latch = golatch.NewLatch(AppID, SecretKey)

		if Proxy != "" {
			if proxyUrl, err := url.Parse(Proxy); err == nil {
				latch.SetProxy(proxyUrl)
			}
		}

		if UseCallbacks {
			latch.OnRequestStart = OnLatchRequestStart
			latch.OnResponseReceive = OnLatchResponseReceive
		}
	}

	return latch, err
}

//Initializes the latch object that will be used by all subcommands. This does pretty much the same as the NewLatch() function.
func NewLatchUser(UserID string, SecretKey string, Proxy string, UseCallbacks bool, OnLatchRequestStart func(request *golatch.LatchRequest), OnLatchResponseReceive func(request *golatch.LatchRequest, response *http.Response, responseBody string)) (latch *golatch.LatchUser, err error) {
	if UserID == "" {
		err = errors.New("You must provide the User ID (--user).")
	}
	if err == nil && SecretKey == "" {
		err = errors.New("You must provide the user secret key (--secret).")
	}

	if err == nil {
		latch = golatch.NewLatchUser(UserID, SecretKey)

		if Proxy != "" {
			if proxyUrl, err := url.Parse(Proxy); err == nil {
				latch.SetProxy(proxyUrl)
			}
		}

		if UseCallbacks {
			latch.OnRequestStart = OnLatchRequestStart
			latch.OnResponseReceive = OnLatchResponseReceive
		}
	}

	return latch, err
}
