package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"log"
	"strings"
	"time"
)

const (
	DB_USER = "root"
	DB_PASS = "123456"
	DB_NAME = "shorten"
	TB_NAME = "shorten_v1"
)

type Config struct {
	DSN string
}

type Data struct {
	c  *Config
	DB *sql.DB
}

var D *Data

type retInfo struct {
	Code      int    `json:"ret"`
	Msg       string `json:"msg"`
	ShortCode string `json:"short_code,omitempty"`
	ShortUrl  string `json:"short_url,omitempty"`
}

func main() {
	conf := &Config{DSN: fmt.Sprintf("%v:%v@/%v?charset=utf8", DB_USER, DB_PASS, DB_NAME)}
	db, err := conf.NewMysql()
	defer func() {
		log.Println("db closing")
		err = db.Close()
		checkErr(err)
	}()
	D = &Data{conf, db,}

	g := gin.Default()
	g.GET("/", Index)
	g.GET("/:code", Redirect)
	g.POST("/shorten", Shorten)
	addr := ":8080"
	err = g.Run(addr)
	checkErr(err)
}

func (c *Config) NewMysql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", c.DSN)
	return
}

func Index(ctx *gin.Context) {
	ctx.String(200, "Hello，世界!")
}

func Redirect(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		ctx.AbortWithStatus(404)
	}
	id := Decode(code)
	log.Println(fmt.Sprintf("Decode: %s --> %d", code, id))

	var url string
	err := D.DB.QueryRow("SELECT url FROM "+TB_NAME+" WHERE id=?", id).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			url = ""
		} else {
			checkErr(err)
		}
	}
	if url != "" {
		url = html.UnescapeString(url)
		log.Println(fmt.Sprintf("Redirect to %s", url))
		ctx.Redirect(301, url)
		return
	}
	ctx.AbortWithStatus(404)
}

func Shorten(ctx *gin.Context) {
	ret := retInfo{}
	url := ctx.DefaultPostForm("url", "")
	if url == "" {
		ret.Code = -1
		ret.Msg = "need param: url"
		ctx.JSON(200, ret)
		return
	}
	log.Println("url: " + url)

	if strings.Index(url, "http://") == -1 && strings.Index(url, "https://") == -1 {
		ret.Code = -2
		ret.Msg = "invalid url style without \"http(s)://\""
		ctx.JSON(200, ret)
		return
	}

	stmt, err := D.DB.Prepare("INSERT IGNORE INTO " + TB_NAME + " SET sign=?,url=?,create_at=?")
	checkErr(err)
	defer func() {
		log.Println("prepared statement closing")
		err = stmt.Close()
		checkErr(err)
	}()
	hash := md5.Sum([]byte(url))
	md5str := fmt.Sprintf("%x", hash)
	result, err := stmt.Exec(md5str, html.EscapeString(url), time.Now().Unix())
	checkErr(err)
	lastId, err := result.LastInsertId()
	checkErr(err)

	id := int(lastId)
	if id == 0 {
		err = D.DB.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE sign='%s'", TB_NAME, md5str)).Scan(&id)
		checkErr(err)
		log.Printf("url exists and find ID:%v\n", id)
	}
	short := Encode(id)

	// update id
	stmt, err = D.DB.Prepare("UPDATE " + TB_NAME + " SET code=?,update_at=? WHERE id=? AND code=''")
	checkErr(err)
	defer func() {
		log.Println("prepared statement closing")
		err = stmt.Close()
		checkErr(err)
	}()
	result, err = stmt.Exec(short, time.Now().Unix(), id)
	checkErr(err)
	if _, err := result.LastInsertId(); err != nil {
		log.Println("update code:" + short)
	}

	ret.Code = 0
	ret.Msg = "success"
	ret.ShortCode = short
	ret.ShortUrl = ctx.Request.Host + "/" + short
	ctx.JSON(200, ret)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
