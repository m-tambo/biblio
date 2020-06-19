package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"biblio/schema"
)

func TestGetAuthorsHandler(t *testing.T) {
	handler := NewGetAuthorsHandler(MockDAO{})

	srv := httptest.NewServer(handler)
	defer srv.Close()

	res, err := srv.Client().Get(srv.URL)
	if err != nil {
		fmt.Println("errorrrr:", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	var un []schema.Author

	json.Unmarshal(body, &un)
	fmt.Println(un)

}
