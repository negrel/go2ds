[![Go Reference](https://pkg.go.dev/badge/github.com/negrel/guds.svg)](https://pkg.go.dev/github.com/negrel/guds)
[![Go Report](https://goreportcard.com/badge/github.com/negrel/guds)](https://goreportcard.com/report/github.com/negrel/guds)

# :zap: GUDS

## Benchmarks

### Maps (built-in, Map & CMap)

```shell
go test -bench . -run Put -benchmem -benchtime 1s ./...                                                                                                                                                                       [10:44:30]
goos: linux
goarch: amd64
pkg: github.com/negrel/guds/pkg/maps/hashmap
BenchmarkBuiltInMap_Put_100-8         	   59280	     24882 ns/op	    3151 B/op	      18 allocs/op
BenchmarkBuiltInMap_Put_1000-8        	    8503	    132849 ns/op	   47875 B/op	      66 allocs/op
BenchmarkBuiltInMap_Put_10000-8       	     930	   1614126 ns/op	  389591 B/op	     256 allocs/op
BenchmarkBuiltInMap_Put_100000-8      	      48	  23125947 ns/op	 3278434 B/op	    3924 allocs/op
BenchmarkBuiltInMap_Get_100-8         	  954879	      1278 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_1000-8        	   50706	     23470 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_10000-8       	    3500	    323723 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_100000-8      	     285	   4186385 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_100-8      	  501444	      2234 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_1000-8     	   31633	     36799 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_10000-8    	    2673	    431905 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_100000-8   	     217	   5454796 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Put_100-8               	   10000	    141336 ns/op	    5389 B/op	       9 allocs/op
BenchmarkCMap_Put_1000-8              	     610	   1890090 ns/op	   85808 B/op	      38 allocs/op
BenchmarkCMap_Put_10000-8             	      58	  21315866 ns/op	  781526 B/op	     926 allocs/op
BenchmarkCMap_Put_100000-8            	      12	 104236713 ns/op	 6378488 B/op	   12042 allocs/op
BenchmarkCMap_Get_100-8               	   31702	     33397 ns/op	      10 B/op	       0 allocs/op
BenchmarkCMap_Get_1000-8              	    3482	    333095 ns/op	      82 B/op	       0 allocs/op
BenchmarkCMap_Get_10000-8             	     360	   3313226 ns/op	     128 B/op	       1 allocs/op
BenchmarkCMap_Get_100000-8            	      24	  44617084 ns/op	    8516 B/op	      88 allocs/op
BenchmarkCMap_Delete_100-8            	   10000	    174783 ns/op	       1 B/op	       0 allocs/op
BenchmarkCMap_Delete_1000-8           	     842	   1408564 ns/op	      53 B/op	       0 allocs/op
BenchmarkCMap_Delete_10000-8          	      84	  12923413 ns/op	     998 B/op	      10 allocs/op
BenchmarkCMap_Delete_100000-8         	      19	  62722267 ns/op	   20862 B/op	     217 allocs/op
BenchmarkMap_Put_100-8                	   86432	     15405 ns/op	    5377 B/op	       9 allocs/op
BenchmarkMap_Put_1000-8               	    7317	    152574 ns/op	   85448 B/op	      34 allocs/op
BenchmarkMap_Put_10000-8              	     788	   1405530 ns/op	  676194 B/op	     214 allocs/op
BenchmarkMap_Put_100000-8             	      73	  15167051 ns/op	 5597046 B/op	    3897 allocs/op
BenchmarkMap_Get_100-8                	  556927	      1814 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_1000-8               	   38859	     30424 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_10000-8              	    2485	    414707 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_100000-8             	     229	   5493025 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_100-8             	  174800	      6744 ns/op	       2 B/op	       0 allocs/op
BenchmarkMap_Delete_1000-8            	   18980	     62965 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_10000-8           	    1706	    682037 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_100000-8          	     141	   8177089 ns/op	    1556 B/op	      10 allocs/op
PASS
ok  	github.com/negrel/guds/pkg/maps/hashmap	108.708s

```

### Contributing

If you want to contribute to GUDS to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/guds/issues)
or make a [pull request](https://github.com/negrel/guds/pulls).

## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License

MIT © [Alexandre Negrel](https://www.negrel.dev/)
