package handlers

import (
	"ConfigCrafter/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

func ConvertHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {

		}
	}(fileContent)

	bytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	inputFilename := file.Filename
	var output string
	var outputFilename string

	if utils.IsYAMLFile(inputFilename) {
		output = utils.YAMLToProperties(bytes)
		outputFilename = changeExtension(inputFilename, "properties")
	} else if strings.HasSuffix(inputFilename, ".properties") {
		output = utils.PropertiesToYAML(bytes)
		outputFilename = changeExtension(inputFilename, "yaml")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", outputFilename))
	c.Data(http.StatusOK, "text/plain", []byte(output))
}

func changeExtension(filename, newExt string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + newExt
}
