package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http, _ := http.Get("https://julekalender-backend.knowit.no/rails/active_storage/blobs/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBNdz09IiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--0af4f790dec929a13e3615fdac11667323ea1597/tall.txt?disposition=inline")
	//bodybytes, _ := io.ReadAll(resp)
	fmt.Println(http.ContentLength)

}

func getnumber(s string, p int) (int, int) {
	trans := map[string]int{
		"en":  1,
		"to":  2,
		"tre": 3,
	}

	substr := s[p:len(s)]

	maxval := -1
	length := 0

	for key, element := range trans {
		if strings.HasPrefix(substr, key) {
			if element > maxval {
				maxval = element
				length = len(key)
			}
		}
	}

	return maxval, length

}

/*
func translate(s string) []int {
	trans := map[string]int{
		"en": 1,
	}


	return []int{5}
}*/
