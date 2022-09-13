package controller

import (
	"converter/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Data struct {
	Method string `json:"method" binding:"required"`
	Suffix string `json:"suffix" binding:"required"`
	Names  string `json:"names" binding:"required"`
}

type File struct {
	Type string `json:"type" binding:"required"`
	Text string `json:"text" binding:"required"`
}

func UploadController(c *gin.Context) {
	var names string
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	open, _ := file.Open()
	p := make([]byte, file.Size)
	open.Read(p)

	split := strings.Split(file.Filename, ".")
	fileType := split[len(split)-1]
	names = util.GetNamesFromFile(fileType, p)
	c.JSON(http.StatusOK, gin.H{"names": names})
}

func ConvertController(c *gin.Context) {
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	method := data.Method
	suffix := data.Suffix
	names := data.Names

	result, err := util.DoConvert(method, suffix, names)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, gin.H{"names": result})
}

func DownloadController(c *gin.Context) {
	var file File

	if err := c.ShouldBindJSON(&file); err != nil {
		c.Data(http.StatusOK, "", nil)
	}

	var bts []byte
	if file.Type == "txt" {
		bts = util.GenerateTxtFile(file.Text)
	} else if file.Type == "xlsx" {
		bts = util.GenerateExcelFile(file.Text)
	}

	c.Data(http.StatusOK, "", bts)
}
