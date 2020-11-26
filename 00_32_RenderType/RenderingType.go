package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Rendering JSON gin.H{}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "OK"})
	})

	//Rendering JSON struct
	r.GET("/jsonStruct", func(c *gin.Context) {
		var person struct {
			Name string `json:"name"`
			Age  int
		}
		person.Name = "My gin Test"
		person.Age = 18

		c.JSON(http.StatusOK, person)
	})

	// Rendering XML
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"test": "xml"})
	})

	server01 := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//Rendering static file
	//set static path static,staticFs,staticFile
	r.Static("/assets", "./assets")
	r.StaticFS("/fs", http.Dir("./"))
	r.StaticFile("/file", "./assets/test.txt")

	// Serving data from reader
	r.GET("/reader", func(c *gin.Context) {
		//http.GET
		res, err := http.Get("https://github.com/gin-gonic/gin#xml-json-yaml-and-protobuf-rendering")
		if err != nil || res.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := res.Body
		// newReader := res.Body

		defer reader.Close()
		// defer newReader.Close()

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment;"`,
		}

		// f, err := os.Create("test.html")

		// if err != nil {
		// 	log.Fatal("create file error!")
		// }
		// bs, err := ioutil.ReadAll(newReader)
		// fmt.Fprint(f, string(bs))

		c.DataFromReader(http.StatusOK, res.ContentLength, res.Header.Get("Content-type"), reader, extraHeaders)
	})

	//Rendering html
	r.GET("/html", func(c *gin.Context) {
		r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "test.html", gin.H{"title": "testHTML"})
	})

	if err := server01.ListenAndServe(); err != nil {
		log.Fatal("start serve error!")
	} else {
		log.Println("started!")
	}
}
