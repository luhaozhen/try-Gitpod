/*
 * @Date: 2022-03-02 17:24:21
 * @FilePath: \tmp\main.go
 * @Description:
 */
package main

import "fmt"

func main() {
	mymap := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	for k, v := range mymap {
		fmt.Println(k, " = ", v)
	}
}

// func main() {

// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello World")
// 	})

// 	router.GET("/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"Example": "Hello Gin",
// 		})
// 	})

// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname")

// 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	})

// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	router.Run(":8000")
// }
