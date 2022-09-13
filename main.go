package main

import (
	"converter/controller"
	"converter/util"
	"embed"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

//go:embed config/* templates/*
var content embed.FS

var methods = []string{
	"张三封 --> ZhangSanFeng",
	"张三封 --> ZhangSF",
	"张三封 --> ZSF",
}

func main() {

	var port string
	defaultPort := "9090"

	var platform = runtime.GOOS

	file, _ := ioutil.ReadFile("./config/default_port.txt")

	if string(file) != "" {
		defaultPort = string(file)
	}

	flag.StringVar(&port, "port", defaultPort, "server port")

	flag.Parse()

	r := gin.Default() // 返回默认路由引擎

	r.SetFuncMap(template.FuncMap{
		"getMethod": func(meth string) string {
			split := strings.Split(meth, " ")
			return split[len(split)-1]
		},
	})
	box := packr.NewBox("./templates")

	r.StaticFS("/lib", box)

	r.LoadHTMLGlob("./templates/index.html")

	//r.Static("/lib", "./templates/lib")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"port":    port,
			"methods": methods,
			"suffixs": util.GetSuffixs(),
		})
	})

	//-----------------------------------
	r.POST("/upload", controller.UploadController)
	r.POST("/convert", controller.ConvertController)
	r.POST("/download", controller.DownloadController)
	//------------------------------------
	// 启动服务
	err := util.OpenBrowser(platform, port)
	if err != nil {
		fmt.Printf("[ERROR] open default browser failed on platform: %s", platform)
	}

	r.Run(":" + port)
}
