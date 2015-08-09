package models

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CheckErr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}

func Get(_url string) []byte {
	resp, err := http.Get(_url)
	if CheckErr(err) {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if CheckErr(err) {
		return nil
	}
	return b
}
