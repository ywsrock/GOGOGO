package main

import (
	"context"
	"fmt"
	swagger "swagger/client"
)

//
// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc gin.HandlerFunc
// }
//
// type Routes []Route
//
// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/",
// 		f,
// 	},
// }

// func f(context *gin.Context) {
// 	context.JSON(200, "Ok")
// }

func main() {
	cfg := swagger.NewConfiguration()
	c := swagger.NewAPIClient(cfg)
	c.ChangeBasePath("http://localhost:8081")
	ctx := context.Background()
	a, b, e := c.DevelopersApi.SearchInventory(ctx, &swagger.DevelopersApiSearchInventoryOpts{})
	fmt.Printf("%s\n%s\n%s", a, b, e)
}

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
//
// 	for _, route := range routes {
// 		r.Handle(route.Method, route.Pattern, route.HandlerFunc)
// 	}
// 	return r
// }
