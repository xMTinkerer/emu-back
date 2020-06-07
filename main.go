package main

import (
    "github.com/gin-gonic/gin"
    "time"
    "math/rand"
)

func main() {
	r := gin.Default()

	r.GET("/eat", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "You ate new cookies!",
		})
	})

	r.GET("/sleep", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        length := rand.Intn(20)
        time.Sleep(time.Duration(length + 20) * time.Second)
		c.JSON(200, gin.H{
			"message": "you slept and feel refreshed",
		})
    })
    
    r.GET("/getEmu", func( c *gin.Context ) {

        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"result": "/img/HatOff.png",
        })    

    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
