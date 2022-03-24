package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostHttpMessage(sendUrl, strBody string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", sendUrl, bytes.NewBuffer([]byte(strBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	resBody := string(body)
	fmt.Println("response Body:", resBody)
	return resBody
}
