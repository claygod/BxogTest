package bxogtest

// Bench test
// Compare the speed of the multiplexer Bxog and other popular multiplexers
//
//   10%
//   █▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
//   30%
//   █████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
//   50%
//   ██████████▒▒▒▒▒▒▒▒▒▒
//   100%
//   ████████████████████
//
// 2016 Eduard Sesigin. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/claygod/Bxog"
	"github.com/daryl/zeus"
	"github.com/go-zoo/bone"
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"github.com/julienschmidt/httprouter"
)

const ADD_PATH_COUNT = 150 // 1 10 50 100 150 250 500 1000

// Test bxog ns/op
func BenchmarkBxogMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := bxog.New()
	muxx.Add("/", Bench2)
	muxx.Add("/a", Bench2)
	muxx.Add("/aas", Bench2)
	muxx.Add("/sd", Bench2)
	muxx.Add("/sd7", Bench2)

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.Add("/produkt"+strconv.Itoa(i)+"/:num/are/:year/goo/:price", Bench2)
	}

	muxx.Add("/produkt900/:num/art/:year/goo/:price", Bench2)

	muxx.Test()
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

/**/
// Test httprouter ns/op
func BenchmarkHttpRouterMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := httprouter.New()

	muxx.Handler("GET", "/", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/a", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/aas", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/sd", http.HandlerFunc(Bench))

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.Handler("GET", "/produkt"+strconv.Itoa(i)+"/:num/are/:year/goo/:price", http.HandlerFunc(Bench))
	}

	muxx.Handler("GET", "/produkt900/:num/art/:year/goo/:price", http.HandlerFunc(Bench))

	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test daryl/zeus ns/op

func BenchmarkZeusMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := zeus.New()

	muxx.GET("/", Bench)
	muxx.GET("/a", Bench)
	muxx.GET("/aas", Bench)
	muxx.GET("/sd/:id", Bench)

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.GET("/produkt"+strconv.Itoa(i)+"/:num/art/:year/goo/:price", Bench)
	}
	muxx.GET("/produkt900/:num/art/:year/goo/:price", Bench)

	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test gorilla/mux ns/op

func BenchmarkGorillaMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := mux.NewRouter()

	muxx.Handle("/", http.HandlerFunc(Bench))
	muxx.Handle("/a", http.HandlerFunc(Bench))
	muxx.Handle("/aas", http.HandlerFunc(Bench))
	muxx.Handle("/sd", http.HandlerFunc(Bench))

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.Handle("/produkt"+strconv.Itoa(i)+"/{num}/art/{year}/goo/{price}", http.HandlerFunc(Bench))
	}
	muxx.Handle("/produkt900/{num}/art/{year}/goo/{price}", http.HandlerFunc(Bench))

	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test gorilla/pat ns/op
func BenchmarkGorillaPatMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := pat.New()

	muxx.Get("/", Bench)
	muxx.Get("/a", Bench)
	muxx.Get("/aas", Bench)
	muxx.Get("/sd", Bench)

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.Get("/produkt"+strconv.Itoa(i)+"/{num}/art/{year}/goo/{price}", http.HandlerFunc(Bench))
	}
	muxx.Get("/produkt900/{num}/art/{year}/goo/{price}", http.HandlerFunc(Bench))

	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test bone ns/op
func BenchmarkBoneMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/produkt900/92f23r44df/art/2016/goo/347889", nil)
	response := httptest.NewRecorder()
	muxx := bone.New()

	muxx.Get("/", http.HandlerFunc(Bench))
	muxx.Get("/a", http.HandlerFunc(Bench))
	muxx.Get("/aas", http.HandlerFunc(Bench))
	muxx.Get("/sd", http.HandlerFunc(Bench))

	for i := ADD_PATH_COUNT; i > 0; i-- {
		muxx.Get("/produkt"+strconv.Itoa(i)+"/:num/art/:year/goo/:price", http.HandlerFunc(Bench))
	}
	muxx.Get("/produkt900/:num/art/:year/goo/:price", http.HandlerFunc(Bench))

	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

func Bench(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("b"))
}

func Bench2(rw http.ResponseWriter, req *http.Request, r *bxog.Router) {
	rw.Write([]byte("b"))
}

/*
			### Result ###

PASS ... ████████████] 100%

BenchmarkBxogMux-4      	 5000000	       330 ns/op
BenchmarkHttpRouterMux-4	 3000000	       395 ns/op
BenchmarkZeusMux-4      	  100000	     23772 ns/op
BenchmarkGorillaMux-4   	   50000	     30223 ns/op
BenchmarkGorillaPatMux-4	 1000000	      1253 ns/op
BenchmarkBoneMux2-4     	   20000	     63656 ns/op
ok  	bxogtest	13.771s

*/
