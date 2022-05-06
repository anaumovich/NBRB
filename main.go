package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type exchange struct {
	CurID           int     `json:"Cur_ID"`
	Date            string  `json:"Date"`
	CurAbbreviation string  `json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `json:"Cur_Name"`
	CurOfficialRate float64 `json:"Cur_OfficialRate"`
}


func main()  {
	g:=router()
	err := g.Run(":8090")
	if err != nil {
		return 
	}
}

func router() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.GET("/rub", sendExchange)
	g.GET("/eur", sendExchange)
	g.GET("/usd", sendExchange)
	return g
}

func newExchange() exchange {
	return exchange{CurID: 0, Date:"" ,CurAbbreviation: "",CurScale:0 ,CurName: "",CurOfficialRate: 0}
}

func checker(c *gin.Context) string {
	switch c.FullPath() {
	case "/rub":
		return "456"
	case "/usd":
		return "431"
	case "/eur":
		return "451"
	}
	return "checker error"
}
func client(c *gin.Context, target interface{}) error {
	ch:= checker(c)
	url:= "https://www.nbrb.by/api/exrates/rates/"+ch+"?periodicity=0"
	log.Println(url)

	customTransport := &(*http.DefaultTransport.(*http.Transport)) // make shallow copy
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport,Timeout: 10 * time.Second}

	r, err := client.Get(url)

	if err != nil {
		log.Println("error code",err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)


	return json.NewDecoder(r.Body).Decode(target) ///write to struct fields
}

func sendExchange(c *gin.Context){
	currency := newExchange()   //create empty struct
	err := client(c, &currency) //call write to strict fields

	if err != nil {
		return
	}

	abv := strconv.FormatFloat(currency.CurOfficialRate, 'f', 4, 64)

	 log.Println ("request time")

	c.JSON(http.StatusOK, normalize(abv))
}


func normalize(s string) string {
	var newStr strings.Builder
	for _,r:= range s{
		switch r {
			case '.': newStr.WriteRune(44)
		default: newStr.WriteRune(r)
		}
	}
	return newStr.String()
}
