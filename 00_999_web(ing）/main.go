package main

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"pageredder/db"
	"pageredder/handler"
	"pageredder/model"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	db.DBInit()
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{"CountUpRow": handler.CountRowCount, "json2String": handler.String2Slice, "dateFormat": handler.Dateformat})
	r.LoadHTMLGlob("templates/*.html")
	v1 := r.Group("/v1")
	{
		v1.GET("/", handler.RouterHandler)
		v1.GET("/c", handler.RouterHandler)
	}
	v2 := r.Group("/v2")
	{

		v2.GET("/", handler.RouterHandlerForJson)
		v2.GET("/c", handler.RouterHandlerForJson)
	}

	conf := ReadConfig()

	s := &http.Server{
		Addr:         conf.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	s.ListenAndServe()
}
func ReadConfig() model.Conf {
	var conf model.Conf
	conf = model.Conf{
		Port:     ":8080",
		RtimeOut: 10,
		WtimeOut: 10,
	}

	f, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("read config error.")
	}

	if err := yaml.Unmarshal(f, &conf); err != nil {
		log.Println("Yaml2Struct Unmarshal error !")
	}
	return conf
}
