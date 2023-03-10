package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"prototype1.0.0/SDISUtils"
)

func main() {
	fmt.Println("Start Secure docker image upload")

	arg := os.Args[1]
	if string(arg) == "" {
		fmt.Println("INPUT THE ARGUMENTS")
		return
	}
	ibcid := SDISUtils.SecureDockerImageUpload(string(arg))
	if ibcid < 0 {
		return
	}
	fmt.Println("Start Secure docker image share")
	result := SDISUtils.BytesToInt(SDISUtils.SecureDockerImageShare(ibcid))
	fmt.Println("Start Secure docker image download")
	SDISUtils.SecureDockerImageDownload(result)

	return

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", HomePage)
	r.GET("/SecureDockerImageUpload", DockerImageUploadPage)
	r.POST("/SecureDockerImageUploadProcess", DockerImageUpload)
	r.GET("/SecureDockerImageShare", DockerImageSharePage)
	r.GET("/SecureDockerImageDownload", DockerImageDownloadPage)

	r.Run(":8888")
	//

}

func DockerImageUpload(c *gin.Context) {
	// 	c.Param(key)
	// 	SDISUtils.SecureDockerImageUpload(string(arg))
}

func HomePage(c *gin.Context) {
	c.HTML(200, "homePage.html", gin.H{
		"message": "Hello World!",
	})
}
func DockerImageUploadPage(c *gin.Context) {
	c.HTML(200, "dockerImageUploadPage.html", gin.H{
		"PageIntention": "UPLOAD",
	})
}

func DockerImageSharePage(c *gin.Context) {
	c.HTML(200, "dockerImageSharePage.html", gin.H{
		"PageIntention": "SHARE",
	})
}
func DockerImageDownloadPage(c *gin.Context) {
	c.HTML(200, "dockerImageDownloadPage.html", gin.H{
		"PageIntention": "DOWNLOAD",
	})
}
