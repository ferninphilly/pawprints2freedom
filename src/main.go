package main

import (
	"fmt"

	fb "github.com/ferninphilly/pawprints2freedom/facebook"
)

func main() {
	fbStruct := fb.NewFBEndpoints()
	accessToken := fbStruct.GetPageAccessToken()
	fmt.Println("Here is the response from fb ", accessToken)
}
