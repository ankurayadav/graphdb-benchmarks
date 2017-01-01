# graphdb-benchmark
Benchmark tests for various graph databases.

### Results of first LoadData benchmark tests (30kmoviedata provided in cayley examples folder) on cayleydb and dgraphdb
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

**Note :** Please ensure you have [Git LFS](https://git-lfs.github.com/) installed, before you clone this repository.