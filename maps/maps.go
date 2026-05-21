package main


import "fmt"



func main() {
	websites := map[string]string{
		"Google": "https://.google.com",
		"Amazone web service": "https://aws.com",
	}

	fmt.Println(websites["Google"])

	delete(websites, "Google")
	fmt.Println(websites)
}