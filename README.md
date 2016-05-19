# BxogTest

Test router Bxog and its benchmark (and other popular routers - multiplexers, written in the Go)

'const ADD_PATH_COUNT = 150'

- BenchmarkBxogMux-4      	 5000000	       330 ns/op
- BenchmarkHttpRouterMux-4	 3000000	       395 ns/op
- BenchmarkZeusMux-4      	  100000	     23772 ns/op
- BenchmarkGorillaMux-4   	   50000	     30223 ns/op
- BenchmarkGorillaPatMux-4	 1000000	      1253 ns/op
- BenchmarkBoneMux2-4     	   20000	     63656 ns/op
