package main

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDecodeAndDecode(t *testing.T) {
	s := "AjaXXe"
	r := Decode(s)
	t.Log(r)
	s2 := Encode(r)
	t.Log(s2)
	if s2 != s {
		t.Fail()
	}
}