package main

import (
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.Request.Host)

		c.String(200, "pong")
	})

	//log.Fatal(autotls.Run(r, "*.jifenhua.cn"))
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		//HostPolicy: autocert.HostWhitelist("jifenhua.cn", "*.jifenhua.cn"),
		//Cache:      autocert.DirCache("/data/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
