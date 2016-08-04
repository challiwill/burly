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
  Protocol string `url:"protocol"`
  Host string `url:"domain"`
  Path string `url:"path"`
  Option1 string `url:"some-param" encode:"false"`
  Option2 string `url:"some-other-param"`
}

func main() {
  myURL, _ := url.Parse(URLStruct{
    Protocol: "https",
    Host:     "www.myhost.com",
    Path:     "/a/path",
    Option1:  "my/value",
    Option2:  "other-value",
  })
  // returns net/url.URL for 'https://www.myhost.com/a/path?some-param=my/value&some-other-param=other%2Fvalue'
  fmt.Println(myURL)
}
```

### DISCLAIMER:
This was more of an experiment and is not very performant. I learned before diving too deep into this that there is a library provided by google that does the bulk of what we wanted already. Probably everyone should use this instead: https://github.com/google/go-querystring.

#### How is this different from go-querystring

Primarily we needed a way to mark some parameters to not be encoded in the traditional query safe way. To do this properly this library returns a net/url.URL instead of just handling the Values.
Some ways in which this library is worse than go-querystring:
1. Currently it does not provide the ability of type safety on param values (I hope to implement this eventually)
1. Currently it is not as performant (I think)

### NOTES

Currently this library does not support URLs with user info, or fragments. A standard URL object represents `scheme://[userinfo@]host/path[?query][#fragment]` or `scheme:opaque[?query][#fragment]`. Currently this library only supports `scheme://host/path[?query]`
