package main

import (
	"api-http/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// var wg sync.WaitGroup
// var mut sync.Mutex

// func sendRequest(url string) {
// 	defer wg.Done()
// 	now := time.Now()
// 	res, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	mut.Lock()
// 	defer mut.Unlock()
// 	fmt.Printf("[%d] %s\n", res.StatusCode, url)
// 	fmt.Println("elapsed", time.Since(now))
// }

// func main() {

// 	go sendRequest("https://www.facebook.com/")
// 	go sendRequest("https://www.google.com/")
// 	go sendRequest("https://www.bing.com/")
// 	wg.Add(3)
// 	wg.Wait()
// 	// go func() {
// 	// 	mux := http.NewServeMux()
// 	// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 		fmt.Printf("server: %s /\n", r.Method)

// 	// 	})
// 	// 	server := http.Server{
// 	// 		Addr:    fmt.Sprintf(":%d", serverPort),
// 	// 		Handler: mux,
// 	// 	}
// 	// 	if err := server.ListenAndServe(); err != nil {
// 	// 		if !errors.Is(err, http.ErrServerClosed) {
// 	// 			fmt.Printf("error running http server: %s\n", err)
// 	// 		}
// 	// 	}
// 	// }()

// 	// time.Sleep(100 * time.Millisecond)

// 	// requestURL := fmt.Sprintf("https://www.facebook.com")
// 	// //requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
// 	// res, err := http.Get(requestURL)
// 	// if err != nil {
// 	// 	fmt.Printf("error making http request: %s\n", err)
// 	// 	//os.Exit(1)
// 	// }

// 	// fmt.Printf("client: got response!\n")
// 	// fmt.Printf("client: status code: %d\n", res.StatusCode)
// }

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/input-url", controllers.C_GetUrlResponse).Methods("POST")
	fmt.Printf("Starting server at port 8082\n")
	log.Fatal(http.ListenAndServe(":8084", r))

}
