package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	g:=router()
	g.Run(":8080")
}

func router() *gin.Engine{
	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	g.GET("/rub", RUB)
	g.GET("/eur", EUR)
	g.GET("/usd", USD)
	return g
}

func RUB(c *gin.Context){
	url:= "https://www.nbrb.by/api/exrates/rates/456?periodicity=0"
	x:=client(url)
	c.JSON(http.StatusOK, x)
}
func EUR(c *gin.Context){
	url:= "https://www.nbrb.by/api/exrates/rates/451?periodicity=0"
	x:=client(url)
	c.JSON(http.StatusOK, x)
}
func USD(c *gin.Context){
	url:= "https://www.nbrb.by/api/exrates/rates/431?periodicity=0"
	x:=client(url)
	c.JSON(http.StatusOK, x)
}

func client(url string) string {

	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "error1"
	}
	//req.Header.Add("Cookie", "ASP.NET_SessionId=flhixhlhg10gq4epbj0zc0oq")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "error2"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return"error3"
	}
	return string(body)
}
