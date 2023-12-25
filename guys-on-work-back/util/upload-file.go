package util

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(context *gin.Context) (*string, error) {
	
	// Extract the file from the form
	file, fileHeader, err := context.Request.FormFile("file")
	if err == nil {

		// Save the file to a specific location (you might want to change this)
		fileName := filepath.Join("./uploads", fileHeader.Filename)
		err = context.SaveUploadedFile(fileHeader, fileName)
		if err != nil {
			return nil, err
		}

	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	return &fileHeader.Filename, nil

}

func DownloadFile(context *gin.Context) {

	pathFile := context.Param("path")

	if pathFile == "" {

		context.JSON(400, gin.H{"error": "ID é obrigatório!"})
		return

	}

	filePath := filepath.Join("./uploads", pathFile)
	context.File(filePath)
	return
}
