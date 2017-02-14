package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for index := 0; index < 10; index++ {
		go func() {
			log.Print(index)
			PostToAPI("http://localhost:5000/api/logs")
		}()
	}

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
}

func PostToAPI(url string) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, _ := netClient.Get(url)

	log.Println(response.Status)
}
