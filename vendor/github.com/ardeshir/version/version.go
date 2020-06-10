package version

// Help for developers:
// go build -ldflags "-X main.version=0.0.1"
// godoc -http :9090
// root@ardeshir ~/httpdoc (master) $ graphpkg -stdout net > graph
// ardeshir.org/graph to view the SVG of all dependencies

import (
	"fmt"
	"log"
	"time"
)

// URL exported
var URL = "https://api.univs.io/my-ip"

// V sets verion
func V(version string) {

	fmt.Printf("Version: %s  Univrs.io\n", version)
	fmt.Printf("Current time: %s\n", time.Now())

}

// ErrNil easy errors
func ErrNil(err error, mesg string) {
	if err != nil {
		log.Fatalln(mesg)
	}
}
