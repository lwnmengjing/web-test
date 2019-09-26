package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
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
		Cache:      autocert.DirCache("./"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
	//r.Run(":80")
}
