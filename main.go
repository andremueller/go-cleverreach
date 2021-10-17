package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	config = clientcredentials.Config{
		ClientID:     "", // TODO get that from $HOME/.config/go-cleverreach/cleverreach.toml
		ClientSecret: "",
		TokenURL:     "https://rest.cleverreach.com/oauth/token.php",
	}
)

const url = "https://rest.cleverreach.com/v3"

func main() {
	client := config.Client(context.Background())

	// the client will update its token if it's expired
	resp, err := client.Get(url + "/mailings.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	} else {
		fmt.Println(resp)
	}
}
