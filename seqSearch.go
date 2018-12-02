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
	// "log"
	"fmt"
   u "github.com/ardeshir/version"
)

// search function to find a value
func SeqSearch(data []int, value int) bool {
   size := len(data)
   for i := 0; i < size; i++ {
     if value == data[i] {
        return true
     }
   }
   return false
}


// summarizing function for ints
func SumArray(data []int) int {

  size  := len(data)
  total := 0
  for index := 0; index < size; index++ {
     total = total + data[index]
  }

   return total
}

func PrintSlice(data []int) {
  fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

/*   MAIN TO RUN FUNCTIONS */
func main() {

    arr := []int{1,2,3,4,5,6,7,8,9}

    PrintSlice(arr)

    fmt.Printf("Total is %d\n ", SumArray(arr))
    val := 9

    if  SeqSearch(arr, val)  {
       fmt.Printf("The value %d was found\n", val)
    } else {
       fmt.Printf("The value %d was not found\n", val)
    }

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
