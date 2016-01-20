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
    "os"
    "log"
	"github.com/ChimeraCoder/anaconda"
)

var LookupEnv func(t string) (string, bool) = os.LookupEnv
var TokenSetter func(t string, s string) *anaconda.TwitterApi = anaconda.NewTwitterApi
var ConsumerKeySetter func(k string) = anaconda.SetConsumerKey
var ConsumerSecretSetter func(s string) = anaconda.SetConsumerSecret

func GetTwitterApi() *anaconda.TwitterApi {
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
