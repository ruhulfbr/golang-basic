package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL in GO")
	strUrl := "https://jsonplaceholder.typicode.com:800/route/params?key1=value1&key2=value2"

	parsedUrl, err := url.Parse(strUrl)
	if err != nil {
		fmt.Println("Error occure while parsing URL:", err)
		return
	}

	fmt.Println("Parsed Url", parsedUrl)
	fmt.Println("Url Schema", parsedUrl.Scheme)
	fmt.Println("Url Host", parsedUrl.Host)
	fmt.Println("Url Path", parsedUrl.Path)
	fmt.Println("Url Query Params", parsedUrl.RawQuery)

	parsedUrl.Path = "new-path1/95"
	parsedUrl.RawQuery = "key9=value9&key8=value8&key7=value7&key6=value6"

	newUrl := parsedUrl.String()
	fmt.Println("New Url:", newUrl)

}
