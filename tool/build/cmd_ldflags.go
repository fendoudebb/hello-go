package main

import "fmt"

var (
	Version   string
	BuildTime string
	GoVersion string
)

// go build -v -ldflags "-X main.Version=1.0 -X main.BuildTime=date -X 'main.GoVersion=`go version`'"
func main() {
	fmt.Printf("%s\n%s\n%s\n", Version, BuildTime, GoVersion)
}
