package main

import (
	"fmt"
	"net/http"
)

func main() {
	key := "ff0965a16d38556b5e961784090b60c2"
	url := "https://rebrickable.com/api/v3/lego/minifigs/"
	for i := 100; i > 0; i-- {
		//amt := strconv.Itoa(i)
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", "key "+key)
		if err != nil {
			fmt.Println("Error making request")
			return
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error getting server")
			return
		}
		defer response.Body.Close()
		amt := response.Status
		fmt.Println(amt + " bottles of beer on the wall, " + amt + " bottles of beer, take one down, pass it around, " + amt + " bottles of beer on the wall!")
	}

}
