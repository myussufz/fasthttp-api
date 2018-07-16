package api

import (
	"log"
	"testing"
)

func TestGetFetch(t *testing.T) {

	resp, err := Fetch("https://www.google.com").ToString()
	log.Println(resp, err)
	if err != nil {
		t.Fatal("unexpected call")
	}

	resp, err = Fetch("https://cdnjs.cloudflare.com/ajax/libs/vue/2.5.16/vue.common.js").ToString()
	log.Println(resp, err)
	if err != nil {
		t.Fatal("unexpected call")
	}
}
