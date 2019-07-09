package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDecode(t *testing.T) {
	s := "AjaXXe"
	r := Decode(s)
	t.Log(r)
	s2 := Encode(r)
	t.Log(s2)
	if s2 != s {
		t.Fail()
	}
}

func TestShortenV1(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.33.10:3306)/?charset=utf8")
	if err != nil {
		t.Error(err)
	}
	row, err := db.Query("SELECT max(`id`) FROM shorten.shorten_v1")
	if err != nil {
		t.Error(err)
	}

	for row.Next() {
		var v interface{}
		err = row.Scan(&v)
		if err != nil {
			t.Error(err)
		}
		t.Logf("结果:%+v\n",v)
	}


	// insert
	db.Prepare("INSERT IGNORE INTO shorten.shorten_v1 (`sign`,`url`)")

	err = db.Close()
	if err != nil {
		t.Error(err)
	}
}
