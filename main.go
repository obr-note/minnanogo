package main

import (
	"flag"
	"fmt"
)

func main() {
	const defaultPort int = 3000
	var port = flag.Int("port", defaultPort, "use")
	flag.Parse()
	fmt.Printf("%v", *port)
}
