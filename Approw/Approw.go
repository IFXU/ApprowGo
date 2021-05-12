package Approw

import (
	"fmt"
	"net/http"
	/*
		"encoding/json"

		"io/ioutil"
		"log"

		"net/url"
	*///authing "github.com/authing/authing-go-sdk"
	//prettyjson "github.com/hokaccha/go-prettyjson"
	//"github.com/kelvinji2009/graphql"
)

type User struct {
	Email    string
	Password string
}

type id_token struct {
	value string
}

func VerifyIdtoken(w http.ResponseWriter, r *http.Request) {
	/*
		var p id_token
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		params := "id_token=" + url.QueryEscape(p.value)
		path := fmt.Sprintf("https://authingagoralogin.authing.cn/api/v2/oidc/validate_token?%s", params)

		resp, err := http.Get(path)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(string(body))*/
	fmt.Println("hello!")
}
