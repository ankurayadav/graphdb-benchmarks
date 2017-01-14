/*
 * Copyright 2017 Ankur Yadav (ankurayadav@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cayley

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/cayleygraph/cayley/graph"
)

var (
	handle *graph.Handle
)

var benchmarkQueries = []struct {
	query string
}{
	// Finding movies and genre of movies directed by "Steven Spielberg"?
	{
		query: ` [{"id" : "<m.06pj8>", 
				"<type.object.name>" : null,
				"<film.director.film>" : {
					"<type.object.name>" : null,
					"<film.film.initial_release_date>" : null
				}
			}]
		`,
	},
	{
		query: ` [{"<type.object.name>" : "\"Steven Spielberg\"@en",
				"<film.director.film>" : {
					"<type.object.name>" : null,
					"<film.film.initial_release_date>" : null
				}
			}]
		`,
	},
}

func runBench(n int, b *testing.B) {

	// Http client for getting JSON response.
	hc := &http.Client{Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
	}}

	b.StopTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := http.NewRequest("POST", "http://127.0.0.1:64210/api/v1/query/mql", bytes.NewBufferString(benchmarkQueries[n].query))

		b.StartTimer()
		resp, err := hc.Do(r)
		b.StopTimer()

		if err != nil {
			log.Fatal("Error in query")
		} else {

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Couldn't parse response body. %+v", err)
			}

			// log.Printf("Response body: %s", body)

		}
	}
}
