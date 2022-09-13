package main

import (
	"converter/assets"
	"converter/conf"
	"converter/controller"
	"converter/util"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"runtime"
)

func main() {

	var IP string
	defaultIP := "127.0.0.1"

	var port string
	defaultPort := "9090"

	var platform = runtime.GOOS

	flag.StringVar(&port, "port", defaultPort, "server port")
	flag.StringVar(&IP, "ip", defaultIP, "server ip")

	flag.Parse()

	r := gin.Default() // 返回默认路由引擎

	r.StaticFS("assets", http.FS(assets.Static))

	funcMap := template.FuncMap{
		"getMethod": func(meth string) string {
			return conf.MethodMap[meth]
		},
	}

	r.SetHTMLTemplate(template.Must(template.New("").Funcs(funcMap).ParseFS(assets.Templates, "templates/*")))

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":      IP,
			"port":    port,
			"methods": conf.Methods,
			"suffixs": conf.Suffixs,
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
