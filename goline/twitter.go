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
	"github.com/ChimeraCoder/anaconda"
	"log"
	"os"
)

// LookupEnv contains the function used to lookup environment variables
var LookupEnv = os.LookupEnv

// TokenSetter contains the function used to login on Twitter as the user, and get the API interface
var TokenSetter = anaconda.NewTwitterApi

// ConsumerKeySetter contains the function used to provide the Consumer key to Twitter
var ConsumerKeySetter = anaconda.SetConsumerKey

// ConsumerSecretSetter contains the function used to provide the Consumer secret to Twitter
var ConsumerSecretSetter = anaconda.SetConsumerSecret

// GetTwitterAPI logins in Twitter and return the authenticated API
func GetTwitterAPI() *anaconda.TwitterApi {
	consumerKey, err1 := LookupEnv("GOLINE_CONSUMER_KEY")
	consumerSecret, err2 := LookupEnv("GOLINE_CONSUMER_SECRET")
	token, err3 := LookupEnv("GOLINE_ACCESS_TOKEN")
	secret, err4 := LookupEnv("GOLINE_ACCESS_TOKEN_SECRET")
	if err1 || err2 || err3 || err4 {
		log.Fatal("Some error happened when accessing GOLINE_* environment variables")
	}
	ConsumerKeySetter(consumerKey)
	ConsumerSecretSetter(consumerSecret)
	return TokenSetter(token, secret)
}
