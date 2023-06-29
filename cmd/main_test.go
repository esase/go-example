package main

import (
	"testing"

	"git.crg.one/scm/go/supplier-hub.git/test/suppliers/example"
)

func TestAllRoutes(t *testing.T) {
	router := setupRouter()

	example.RunAllTests(router, t)
}
