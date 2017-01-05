# graphdb-benchmarks
Benchmark tests for various graph databases.

### Results of first LoadData benchmark tests.
#### Loading 30kmoviedata provided in cayley examples folder
1. **Loading data in cayley graph with bolt backend.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=cayley
Starting benchmark tests for cayley.
==> Executing: go test ./cayley/ -bench=.
==> Output: testing: warning: no tests to run
==> Executing: cayley init -db=bolt -dbpath=cayleydb -quads=../data/30kmoviedata.nq.gz
BenchmarkImportDataToDB-4   	       1	30028388985 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/cayley	30.035s
</pre>
2. **Loading data in dgraph.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=dgraph
Starting benchmark tests for dgraph.
==> Executing: go test ./dgraph/ -bench=.
==> Output: testing: warning: no tests to run
==> Executing: dgraphloader -r=../data/30kmoviedata.nq.gz
==> Output:
Processing ../data/30kmoviedata.nq.gz
Number of mutations run   : 472ne:   471705 RDFs per second:  124927
Number of RDFs processed  : 471705
Time spent                : 4.715085371s
RDFs processed per second : 117926
BenchmarkImportDataToDB-4   	       1	4738905881 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/dgraph	4.745s
</pre>

#### Loading 21million.rdf.gz data
1. **Loading data in cayley graph with bolt backend.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=cayley
Starting benchmark tests for cayley.
==> Executing: go test -timeout=5h ./cayley/ -bench=.
==> Output: testing: warning: no tests to run
==> Executing: cayley init -db=bolt -dbpath=cayleydb -quads=../data/21million.rdf.gz
BenchmarkImportDataToDB-4   	       1	8527419635659 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/cayley	8527.427s
</pre>
2. **Loading data in dgraph.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=dgraph
Starting benchmark tests for dgraph.
==> Executing: go test -timeout=5h ./dgraph/ -bench=.
==> Output: testing: warning: no tests to run
==> Executing: dgraphloader -r=../data/21million.rdf.gz
==> Output:
Processing ../data/21million.rdf.gz
Number of mutations run   : 21240: 21239872 RDFs per second:   24229
Number of RDFs processed  : 21239872
Time spent                : 14m39.848819629s
RDFs processed per second : 24163
BenchmarkImportDataToDB-4   	       1	879903054838 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/dgraph	879.913s
</pre>

### Results of Queries benchmark.
#### Query to find all movies and their genre which are directed by director "Steven Spielberg".
1. **Querying in cayley graph with bolt backend.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=cayley -bench=queries
Starting benchmark tests for cayley.
Starting queries benchmarking tests.
==> Executing: go test -timeout 4h ./cayley/ -bench BenchmarkQuery*
==> Output: testing: warning: no tests to run
BenchmarkQueryFilmByDirector-4   	      20	  64747834 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/cayley	1.376s
</pre>
2. **Querying in dgraph.**
<pre>
Ankurs-MacBook-Pro:graphdb-benchmark ankur$ go run main.go -db=dgraph -bench=queries
Starting benchmark tests for dgraph.
Starting queries benchmarking tests.
==> Executing: go test -timeout 4h ./dgraph/ -bench BenchmarkQuery*
==> Output: testing: warning: no tests to run
BenchmarkQueryFilmByDirector-4   	    1000	   1764900 ns/op
PASS
ok  	_/Users/ankuryadav/dev/benchmark/graphdb-benchmark/dgraph	2.067s
</pre>

**Note :** Please ensure you have [Git LFS](https://git-lfs.github.com/) installed, before you clone this repository.