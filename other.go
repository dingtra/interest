package interest

import (
	"net/http"
	"strings"
	"fmt"
)


type CallStruct struct {
	Urls []string
}


func (let *CallStruct) Route(r *http.Request) {

	url := strings.Split(r.URL.Path[len("/ajx/"):], "/")

	for _, k := range url {

		if k != ""{
			let.Urls = append(let.Urls, strings.ToLower(k))
		}
	}

	fmt.Println(let.Urls)
}


