// Copyright © 2016 Ignasi Fosch <ignasi.fosch@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goline

import (
	goline "."
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"reflect"
	"testing"
)

type MockTwitterApi struct {
	Credentials *oauth.Credentials
}

var rToken string
var rSecret string

func newMockTwitterApi(token string, secret string) *anaconda.TwitterApi {
	rToken = token
	rSecret = secret
	return &anaconda.TwitterApi{
		Credentials: &oauth.Credentials{
			Token:  token,
			Secret: secret,
		},
	}
}

var rConsumerKey string

func mockKeySetter(key string) {
	rConsumerKey = key
}

var rConsumerSecret string

func mockSecretSetter(secret string) {
	rConsumerSecret = secret
}

var xConsumerKey = "ConsumerKey"
var xConsumerSecret = "ConsumerSecret"
var xToken = "ConsumerKey"
var xSecret = "ConsumerSecret"

func mockLookupEnv(k string) (string, bool) {
	switch k {
	case "GOLINE_CONSUMER_KEY":
		return xConsumerKey, false
	case "GOLINE_CONSUMER_SECRET":
		return xConsumerSecret, false
	case "GOLINE_ACCESS_TOKEN":
		return xToken, false
	case "GOLINE_ACCESS_TOKEN_SECRET":
		return xSecret, false
	}
	return "", true
}

func TestGetTwitterApi(t *testing.T) {
	goline.TokenSetter = newMockTwitterApi
	goline.ConsumerKeySetter = mockKeySetter
	goline.ConsumerSecretSetter = mockSecretSetter
	goline.LookupEnv = mockLookupEnv

	var result interface{} = goline.GetTwitterApi()

	v, err := result.(MockTwitterApi)
	if err {
		t.Fatalf("GetTwitterApi returned %s instead of anaconda.TwitterApi", reflect.TypeOf(v))
	}
	if rConsumerKey != xConsumerKey {
		t.Fatalf("Received consumer key %s instead of %s", rConsumerKey, xConsumerKey)
	}
	if rConsumerSecret != xConsumerSecret {
		t.Fatalf("Received consumer secret %s instead of %s", rConsumerSecret, xConsumerSecret)
	}
	if rToken != xToken {
		t.Fatalf("Received token %s instead of %s", rToken, xToken)
	}
	if rSecret != xSecret {
		t.Fatalf("Received token secret %s instead of %s", rSecret, xSecret)
	}
}
