package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	url := "http://localhost:8000/api/auth/register"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `"name":"khanhvn2301",` + "" + `"email":"khanhvungoc2301999@gmail.com",` + "" + `"password":"12345678",` + "" + `"passwordConfirm":"12345678",` + "" + `"photo":""` + "" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
