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
		query: `
				var director = g.V("<m.06pj8>").Out("<type.object.name>").ToArray()[0]
				var b = g.V("<m.06pj8>").Out("<film.director.film>").ToArray()

				_.each(b, function(s){
				var film_name = g.V(s).Out("<type.object.name>").ToArray()[0]
				var file_release = g.V(s).Out("<film.film.initial_release_date>").ToArray()[0]
				var movie = {}
				movie["director"] = director
				movie["film_name"] = film_name
				movie["file_release"] = file_release
				
				var genre = g.V(s).Out("<film.film.genre>").ToArray()
				var genre_list = ""
				_.each(genre, function(gen){
						var genre_name =  g.V(gen).Out("<type.object.name>").ToArray()[0]
						genre_list += genre_name + ", "
				})
				
				movie["genre_list"] = genre_list
				
				g.Emit(movie)
				})
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
		r, err := http.NewRequest("POST", "http://127.0.0.1:64210/api/v1/query/gremlin", bytes.NewBufferString(benchmarkQueries[n].query))

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

			//log.Printf("Response body: %s", body)

		}
	}
}
