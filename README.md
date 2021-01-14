[![Go Reference](https://pkg.go.dev/badge/github.com/negrel/guds.svg)](https://pkg.go.dev/github.com/negrel/guds)
[![Go Report](https://goreportcard.com/badge/github.com/negrel/guds)](https://goreportcard.com/report/github.com/negrel/guds)
[![License Badge](https://img.shields.io/github/license/negrel/guds)](https://github.com/negrel/guds/raw/master/LICENSE)

# :zap: GUDS

## Benchmarks

### Maps (built-in, Map & CMap)

```shell
go test -bench . -run NO_TEST -benchmem -benchtime 1s ./...
goos: linux
goarch: amd64
pkg: github.com/negrel/guds/pkg/maps/hashmap
BenchmarkBuiltInMap_Set_100-8         	  117738	      9675 ns/op	    3150 B/op	      18 allocs/op
BenchmarkBuiltInMap_Set_1000-8        	   10000	    100894 ns/op	   47880 B/op	      66 allocs/op
BenchmarkBuiltInMap_Set_10000-8       	    1392	    812083 ns/op	  389632 B/op	     256 allocs/op
BenchmarkBuiltInMap_Set_100000-8      	     145	   8191015 ns/op	 3279393 B/op	    3929 allocs/op
BenchmarkBuiltInMap_Get_100-8         	 1345612	       890 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_1000-8        	   68545	     17063 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_10000-8       	    4689	    231811 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Get_100000-8      	     402	   3006477 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_100-8      	  624742	      1821 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_1000-8     	   41936	     30271 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_10000-8    	    3465	    312383 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltInMap_Delete_100000-8   	     306	   3877898 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Put_100-8               	   15492	     78019 ns/op	    5390 B/op	       9 allocs/op
BenchmarkCMap_Put_1000-8              	    1766	    676897 ns/op	   85845 B/op	      38 allocs/op
BenchmarkCMap_Put_10000-8             	     164	   6884391 ns/op	  693847 B/op	     398 allocs/op
BenchmarkCMap_Put_100000-8            	      13	  84206797 ns/op	 6641609 B/op	   11067 allocs/op
BenchmarkCMap_Get_100-8               	   37569	     31951 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Get_1000-8              	    3573	    318496 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Get_10000-8             	     378	   3176270 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Get_100000-8            	      36	  34339466 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Delete_100-8            	   12807	     98609 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap_Delete_1000-8           	    1438	    841091 ns/op	      31 B/op	       0 allocs/op
BenchmarkCMap_Delete_10000-8          	     235	   4860092 ns/op	     679 B/op	       7 allocs/op
BenchmarkCMap_Delete_100000-8         	      19	  54757886 ns/op	   14617 B/op	     152 allocs/op
BenchmarkMap_Set_100-8                	   93834	     11982 ns/op	    5376 B/op	       9 allocs/op
BenchmarkMap_Set_1000-8               	    9820	    134850 ns/op	   85457 B/op	      34 allocs/op
BenchmarkMap_Set_10000-8              	    1093	   1088751 ns/op	  676187 B/op	     214 allocs/op
BenchmarkMap_Set_100000-8             	      99	  10956819 ns/op	 5599152 B/op	    3912 allocs/op
BenchmarkMap_Get_100-8                	  754118	      1418 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_1000-8               	   53481	     22454 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_10000-8              	    3579	    304384 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Get_100000-8             	     313	   3748898 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_100-8             	  202149	      5600 ns/op	       2 B/op	       0 allocs/op
BenchmarkMap_Delete_1000-8            	   23496	     49741 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_10000-8           	    2280	    501518 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_100000-8          	     198	   5989757 ns/op	    2990 B/op	      20 allocs/op
BenchmarkSyncMap_Set_100-8            	   11186	    106734 ns/op	   12101 B/op	     310 allocs/op
BenchmarkSyncMap_Set_1000-8           	    1245	   1030323 ns/op	  168504 B/op	    3787 allocs/op
BenchmarkSyncMap_Set_10000-8          	      98	  11566869 ns/op	 1469320 B/op	   40201 allocs/op
BenchmarkSyncMap_Set_100000-8         	       8	 138577154 ns/op	14092458 B/op	  416296 allocs/op
BenchmarkSyncMap_Get_100-8            	   33577	     35695 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMap_Get_1000-8           	    3495	    340049 ns/op	       3 B/op	       0 allocs/op
BenchmarkSyncMap_Get_10000-8          	     354	   3454866 ns/op	      83 B/op	       0 allocs/op
BenchmarkSyncMap_Get_100000-8         	      32	  35057247 ns/op	    3123 B/op	      32 allocs/op
BenchmarkSyncMap_Delete_100-8         	   10000	    142479 ns/op	      19 B/op	       1 allocs/op
BenchmarkSyncMap_Delete_1000-8        	    2091	    580446 ns/op	      90 B/op	       1 allocs/op
BenchmarkSyncMap_Delete_10000-8       	     202	   6188399 ns/op	    7456 B/op	      78 allocs/op
BenchmarkSyncMap_Delete_100000-8      	      19	  61737213 ns/op	  328750 B/op	    3425 allocs/op
PASS
ok  	github.com/negrel/guds/pkg/maps/hashmap	131.109s
```

### Contributing

If you want to contribute to GUDS to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/guds/issues)
or make a [pull request](https://github.com/negrel/guds/pulls).

## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License

MIT © [Alexandre Negrel](https://www.negrel.dev/)
