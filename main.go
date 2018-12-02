package main
// Help for developers:
// go build -ldflags "-X main.version=0.0.1"
// godoc -http :9090
// root@ardeshir ~/httpdoc (master) $ graphpkg -stdout net > graph
// ardeshir.org/graph to view the SVG of all dependencies 

import (
	"net/http"
	"flag"
	// "encoding/json"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
   "os"
   "time"
   "math/rand"
	// "log"
	"fmt"
   u "github.com/ardeshir/version"
)


func process2(ch chan int) {
    n := rand.Intn(3000)
    time.Sleep(time.Duration(n) *time.Millisecond)
    ch <- n 
}

func main() {
    
    ch := make(chan int)
    go process2(ch)

    fmt.Println("Waiting for process...")
    n := <- ch
    fmt.Printf("Process took %dms\n", n)
    
    
    /*_---_______------_________________---_*/
    fmt.Println("\n\n\n");
    
    url2get  := flag.String("url", defaultURL() , "Link to download")
    flag.Parse()
    
    // resp, err := http.Get(os.Args[1])
    
    resp, err := http.Get(*url2get)
    u.ErrNil(err, "Unable to read response")
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    u.ErrNil(err, "Unable to read content")

    fmt.Println(string(content))
    
    
    
 // --------- basic housekeeping ----------- // 
 // if we're int debug mode, print out info 
 
   if debugTrue() {
    u.V( defaultVersion() )
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
func defaultVersion() string {

 if os.Getenv("DEFAULT_VERSION") != "" {
     return os.Getenv("DEFAULT_VERSION")
  }    
  
    var version string = "0.0.1"
 return version 
}

// Function to check env variable DEFAULT_URL to http get
func defaultURL() string {
    if os.Getenv("DEFAULT_URL") != "" {
        return os.Getenv("DEFAULT_URL")
    }
    return "https://httpbin.org/get"
}