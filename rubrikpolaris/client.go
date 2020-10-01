// Copyright 2018 Rubrik, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License prop
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rubrikcdm transforms the Rubrik API functionality into easy to consume functions. This eliminates the need to understand
// how to consume raw Rubrik APIs with Go and extends upon one of Rubrik’s main design centers - simplicity. Rubrik’s API first architecture enables
// organizations to embrace and integrate Rubrik functionality into their existing automation processes.
package rubrikpolaris

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"sort"
	"time"

	"github.com/gobuffalo/packr/v2"
)

// Type and Constants are used for escaping Get requests
type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeQueryComponent
)

// Credentials contains the parameters used to authenticate against the Rubrik cluster and can be consumed
// through ConnectEnv() or Connect().
type Credentials struct {
	PolarisDomain string
	Username      string
	Password      string
}

var polarisAuthentication apiToken

type apiToken struct {
	Token   string
	Created time.Time
}

// Connect initializes a new API client based on manually provided Rubrik cluster credentials. When possible,
// the Rubrik credentials should not be stored as plain text in your .go file. ConnectEnv() can be used
// as a safer alternative.
func Connect(nodeIP, username, password string, apiToken ...string) *Credentials {
	client := &Credentials{
		PolarisDomain: nodeIP,
		Username:      username,
		Password:      password,
	}

	return client
}

// ConnectEnv is the preferred method to initialize a new API client by attempting to read the
// following environment variables:
//
//  rubrik_polaris_domain
//
//  rubrik_cdm_token
//
//  rubrik_polaris_username
//
//  rubrik_polaris_password
//
// rubrik_cdm_token will always take precedence over rubrik_polaris_username and rubrik_polaris_password
func ConnectEnv() (*Credentials, error) {

	polarisDomain, ok := os.LookupEnv("rubrik_polaris_domain")
	if ok != true {
		return nil, errors.New("The `rubrik_polaris_domain` environment variable is not present")
	}

	var client *Credentials

	username, ok := os.LookupEnv("rubrik_polaris_username")
	if ok != true {
		return nil, errors.New("The `rubrik_polaris_username` or `rubrik_cdm_token` environment variable is not present")
	}
	password, ok := os.LookupEnv("rubrik_polaris_password")
	if ok != true {
		return nil, errors.New("The `rubrik_polaris_password` or `rubrik_cdm_token` environment variable is not present")
	}

	client = &Credentials{
		PolarisDomain: polarisDomain,
		Username:      username,
		Password:      password,
	}

	return client, nil

}

// Consolidate the base API functions.
func (c *Credentials) commonAPI(callType string, config map[string]interface{}, timeout int) (interface{}, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(timeout),
	}

	var requestURL string
	if callType == "graphql" {
		requestURL = fmt.Sprintf("https://%s.my.rubrik.com/api/graphql", c.PolarisDomain)

	} else {
		requestURL = fmt.Sprintf("https://%s.my.rubrik.com/api/session", c.PolarisDomain)

	}

	var request *http.Request

	convertedConfig, _ := json.Marshal(config)
	request, _ = http.NewRequest("POST", requestURL, bytes.NewBuffer(convertedConfig))

	if callType == "graphql" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", polarisAuthentication.Token))

	} else {
		request.SetBasicAuth(c.Username, c.Password)

	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Rubrik Polaris Go SDK v1.0.0")

	apiRequest, err := client.Do(request)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return nil, errors.New("Unable to establish a connection to the Rubrik cluster")
	} else if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(apiRequest.Body)

	apiResponse := []byte(body)

	var convertedAPIResponse interface{}
	if err := json.Unmarshal(apiResponse, &convertedAPIResponse); err != nil {

		// DELETE request will return a 204 No Content status
		if apiRequest.StatusCode == 204 {
			convertedAPIResponse = map[string]interface{}{}
			convertedAPIResponse.(map[string]interface{})["statusCode"] = apiRequest.StatusCode
		} else if apiRequest.StatusCode != 200 {
			return nil, fmt.Errorf("%s", apiRequest.Status)
		}

	}

	//
	if reflect.TypeOf(convertedAPIResponse).Kind() == reflect.Slice {
		return convertedAPIResponse, nil
	}

	if _, ok := convertedAPIResponse.(map[string]interface{})["errorType"]; ok {
		return nil, fmt.Errorf("%s", convertedAPIResponse.(map[string]interface{})["message"])
	}

	if _, ok := convertedAPIResponse.(map[string]interface{})["message"]; ok {
		// Add exception for bootstrap
		if _, ok := convertedAPIResponse.(map[string]interface{})["setupEncryptionAtRest"]; ok {
			return convertedAPIResponse, nil

		}

		return nil, fmt.Errorf("%s", convertedAPIResponse.(map[string]interface{})["message"])
	}

	return convertedAPIResponse, nil

}

// httpTimeout returns a default timeout value of 15 or use the value provided by the end user
func httpTimeout(timeout []int) int {
	if len(timeout) == 0 {
		return int(15) // if not timeout value is provided, set the default to 15
	}
	return int(timeout[0]) // set the timeout value to the first value in the timeout slice

}

// Post sends a POST request to the provided Rubrik API endpoint and returns the full API response. Supported "apiVersions" are v1, v2, and internal.
// The optional timeout value corresponds to the number of seconds to wait to establish a connection to the Rubrik cluster before returning a
// timeout error. If no value is provided, a default of 15 seconds will be used.
func (c *Credentials) Query(query string, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	c.generateAPIToken(httpTimeout)

	config := map[string]interface{}{}
	config["query"] = query

	apiRequest, err := c.commonAPI("graphql", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil

}

func (c *Credentials) QueryWithVariables(query string, variables map[string]interface{}, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	c.generateAPIToken(httpTimeout)

	config := map[string]interface{}{}
	config["query"] = query
	config["variables"] = variables

	apiRequest, err := c.commonAPI("graphql", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil

}

func (c *Credentials) generateAPIToken(timeout ...int) (string, error) {

	httpTimeout := httpTimeout(timeout)

	// // The Polaris API Tokens expire after 24 hours. To allow for wiggle room
	// // tokenHasExpired will return true if it has been 23 hours since the
	// // token was created.
	tokenExpiresAt := time.Now().Add(-23 * time.Hour) // PROD
	// tokenExpiresAt := time.Now().Add(-5 * time.Second) // TEST

	tokenHasExpired := tokenExpiresAt.After(polarisAuthentication.Created)

	if polarisAuthentication.Token == "" || tokenHasExpired {

		config := map[string]interface{}{}
		config["username"] = c.Username
		config["password"] = c.Password
		apiRequest, err := c.commonAPI("apiToken", config, httpTimeout)
		if err != nil {
			return "", err
		}

		polarisAuthentication.Token = apiRequest.(map[string]interface{})["access_token"].(string)
		polarisAuthentication.Created = time.Now()
		return polarisAuthentication.Token, nil

	} else {
		return polarisAuthentication.Token, nil

	}

}

func (c *Credentials) readQueryFile(filePath string, timeout ...int) (string, error) {

	// set up a new box by giving it a name and an optional (relative) path to a folder on disk:
	box := packr.New("Static GQL Files", "./query")

	// Get the string representation of a file, or an error if it doesn't exist:
	file, err := box.FindString(filePath)
	if err != nil {
		return "", err
	}

	return file, nil

}

// stringEq converts b to []string, sorts the two []string, and checks for equality
func stringEq(a []string, b []interface{}) bool {

	// Convert []interface {} to []string
	c := make([]string, len(b))
	for i, v := range b {
		c[i] = fmt.Sprint(v)
	}

	sort.Strings(a)
	sort.Strings(c)

	// If one is nil, the other must also be nil.
	if (a == nil) != (c == nil) {
		return false
	}

	if len(a) != len(c) {
		return false
	}

	for i := range a {
		if a[i] != c[i] {
			return false
		}
	}

	return true
}
