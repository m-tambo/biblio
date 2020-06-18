package handler

import (
	"net/http/httptest"
	"testing"
)

func TestGetAuthorsHandler(t *testing.T) {
	handler := NewGetAuthorsHandler(MockDAO{})

	srv := httptest.NewServer(handler)
	srv.Start()
	defer srv.Close()

}
