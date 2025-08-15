package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("imageName")
		fileHeader, err := c.FormFile("image")
		/*
				          *For Multiple file uploads
						  form, err := c.MultipartForm()
						  !handle error relating to multiparts
						  if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
							return
						}
						  files := form.File["images"]
						  *loop through to access each file details
						  for _,file := range files{
						  !not sure but u should be able to access fileName from file
						  	dst := filepath.Join("./uploads/", file.Filename)
						  }
							if err = c.SaveUploadedFile(fileHeader, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
				return
			}
		*/
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
			return
		}
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
			return
		}
		defer file.Close()
		dst := filepath.Join("./uploads/", fileHeader.Filename)
		if err = c.SaveUploadedFile(fileHeader, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": fmt.Sprintf("Image uploaded with name %s", name)})
	})
	server.Run(":8083")

}
