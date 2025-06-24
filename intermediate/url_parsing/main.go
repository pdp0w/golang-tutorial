package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment] <---- this is the structure of url

	rawUrl := "https://example.com:8080/path?query=parament#frags"

	parseUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return
	}

	fmt.Println("Schema:", parseUrl.Scheme)
	fmt.Println("Host:", parseUrl.Host)
	fmt.Println("Port:", parseUrl.Port())
	fmt.Println("Path:", parseUrl.Path)
	fmt.Println("Query:", parseUrl.RawQuery)
	fmt.Println("Fragments:", parseUrl.Fragment)

	rawUrl2 := "https://example.com/path?name=pdp0w&age=21"
	parseUrl1, err := url.Parse(rawUrl2)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return
	}

	queryParams := parseUrl1.Query()
	fmt.Println(queryParams)
	fmt.Println(queryParams.Get("name"))
	fmt.Println(queryParams.Get("age"))

	// building url
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("key", "value")
	query.Set("age", "21")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Built Url:", baseUrl.String())

	values := url.Values{}
	// add key-value pairs

	values.Add("name", "Jane")
	values.Add("age", "32")

	// Encode
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)
}
