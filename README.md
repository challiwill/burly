# burly
A convenience library for really long URL generation in Go

For encoding a struct of URL components as URL:

``` go
package main

import (
  "fmt"
  "github.com/challiwill/burly/url"
)

type URLStruct struct {
  Protocol string `comp:"protocol"`
  Host string `comp:"domain"`
  Path string `comp:"path"`
  Option1 string `comp:"some-param"`
  Option2 string `comp:"some-other-param"`
}

func main() {
  myURL, _ := url.Parse(URLStruct{
    Protocol: "https",
    Host:     "www.myhost.com",
    Path:     "/a/path",
    Option1:  "my-value",
    Option2:  "123",
  })
  // returns net/url.URL for 'https://www.myhost.com/a/path?some-param=my-value&some-other-param=123'
  fmt.Println(myURL) 
}
```

### DISCLAIMER: 
This was more of an experiment and is not very performant. I learned before diving too deep into this that there is a library provided by google that does the bulk of what we wanted already. Probably everyone should use this instead: https://github.com/google/go-querystring.

