package main

// Help for developers:
// go build -ldflags "-X main.version=0.0.1"
// godoc -http :9090
// root@ardeshir ~/httpdoc (master) $ graphpkg -stdout net > graph
// ardeshir.org/graph to view the SVG of all dependencies

import (
	"flag"
	"net/http"
	"strings"

	// "encoding/json"
	// "encoding/base64"
	// "bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	// "log"
	"fmt"

	u "github.com/ardeshir/version"
)

// URL sets url
var URL = "https://httpbin.org/"

func process2(ch chan int) {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {

	ch := make(chan int)
	go process2(ch)

	fmt.Println("Waiting for process...")
	n := <-ch
	fmt.Printf("Process took %dms\n", n)

	/*_---_______------_________________---_*/
	fmt.Println("\n\n\n")

	url2get := flag.String("url", defaultURL2(URL), "Link to download")
	flag.Parse()

	// create a client

	client := http.DefaultClient
	req, err3 := http.NewRequest("GET", URL, nil)
	u.ErrNil(err3, "Unable to get")
	resp3, err3 := client.Do(req)
	u.ErrNil(err3, "Unable to get request with Client")
	defer resp3.Body.Close()
	content3, err3 := ioutil.ReadAll(resp3.Body)
	u.ErrNil(err3, "Undable to read body")
	fmt.Println(string(content3))

	// resp, err := http.Get(os.Args[1])
	resp, err := http.Get(*url2get)
	u.ErrNil(err, "Unable to read response")
	defer resp.Body.Close()
	// content, err := ioutil.ReadAll(resp.Body)
	u.ErrNil(err, "Unable to read content")

	// fmt.Println(string(content))

	resp2, err1 := http.Post("https://httpbin.org/post", "text/plain",
		strings.NewReader("this is the request string"))
	u.ErrNil(err1, "Unable to read response")
	defer resp2.Body.Close()
	content2, err1 := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(content2))

	// --------- basic housekeeping ----------- //
	// if we're int debug mode, print out info

	if debugTrue() {
		u.V(defaultVersion2())
	}
}

// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {

	if os.Getenv("DEFAULT_DEBUG") != "" {
		return true
	}
	return false
}

// Function to check env variable DEFAULT_VERSION string
func defaultVersion2() string {

	if os.Getenv("DEFAULT_VERSION") != "" {
		return os.Getenv("DEFAULT_VERSION")
	}

	version := "0.0.1"
	return version
}

// Function to check env variable DEFAULT_URL to http get
func defaultURL2(url string) string {
	if os.Getenv("DEFAULT_URL") != "" {
		return os.Getenv("DEFAULT_URL")
	} else if url == "" {
		return "https://httpbin.org/get"
	}
	return url
}
