package url

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Parse(urlStruct interface{}) (*url.URL, error) {
	var protocol string
	var domain string
	var path string
	var params string
	var field reflect.StructField
	var comp string

	typ := reflect.ValueOf(urlStruct).Type()
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("I need my argument of type interface{} to be a struct - you gave me a %s '%#v'", typ.Kind(), urlStruct)
	}

	val := reflect.Indirect(reflect.ValueOf(urlStruct))

	for i := 0; i < typ.NumField(); i++ {
		field = typ.Field(i)
		if field.Type.Kind() != reflect.String {
			return nil, fmt.Errorf("I need all the fields in my argument to be strings - you gave me a struct with a field '%s' which is a %s", field.Name, field.Type.Kind())
		}

		comp = field.Tag.Get("url")
		switch comp {
		case "":
			// ignore

		case "protocol":
			p := val.FieldByName(field.Name)
			protocol = string(p.String())

		case "domain":
			d := val.FieldByName(field.Name)
			domain = string(d.String())

		case "path":
			p := val.FieldByName(field.Name)
			path = p.String()

		default:
			p := val.FieldByName(field.Name)
			if params != "" {
				params += "&"
			}
			pStr := p.String()
			encode := field.Tag.Get("encode")
			if encode != "false" {
				pStr = url.QueryEscape(pStr)
			}
			params += comp + "=" + pStr
		}
	}

	return &url.URL{
		Scheme:   protocol,
		Host:     domain,
		Path:     makeSafePath(path),
		RawQuery: params,
	}, nil
}

func makeSafePath(path string) string {
	var p string
	if !strings.HasPrefix(path, "/") {
		p = "/"
	}
	return p + path
}
