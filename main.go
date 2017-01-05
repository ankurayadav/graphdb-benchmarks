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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	db    = flag.String("db", "", "Database on which to run benchmark tests. (dgraph or cayley)")
	bench = flag.String("bench", "", "Which type of functions to benchmark. (dataimport or queries)")
)

func main() {

	flag.Parse()

	var pckgDir string
	var benchregex string

	if *db != "" {
		if *db == "dgraph" {
			fmt.Println("Starting benchmark tests for dgraph.")
			pckgDir = "./dgraph/"

		} else if *db == "cayley" {
			fmt.Println("Starting benchmark tests for cayley.")
			pckgDir = "./cayley/"

		} else {
			log.Fatal("Given Database name not supported for benchmarking!")
		}
	} else {
		log.Fatal("No Database name selected for benchmarking!")
	}

	if *bench != "" {
		if *bench == "dataimport" {
			fmt.Println("Starting dataimport benchmarking tests.")
			benchregex = "BenchmarkImportDataToDB"

		} else if *bench == "queries" {
			fmt.Println("Starting queries benchmarking tests.")
			benchregex = "BenchmarkQuery*"

		} else {
			log.Fatal("Given type of benchmarking tests not supported.")
		}
	} else {
		log.Fatal("No benchmarking tests selected!")
	}

	cmd := exec.Command("go", "test", "-timeout", "4h", pckgDir, "-bench", benchregex)

	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)

}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
