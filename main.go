package main

import (
	"fmt"

	"github.com/challiwill/burly/url"
)

type URLStruct struct {
	Protocol string `comp:"protocol"`
	Host     string `comp:"domain"`
	Path     string `comp:"path"`
	Option1  string `comp:"some-param"`
	Option2  string `comp:"some-other-param"`
}

func main() {
	myURL, _ := url.Parse(URLStruct{
		Protocol: "https",
		Host:     "www.myhost.com",
		Path:     "/a/path",
		Option1:  "my-value",
		Option2:  "123",
	})
	fmt.Println(myURL) // returns URL that corresponds to 'https://www.myhost.com/a/path?some-param=my-value&some-other-param=123'
}
