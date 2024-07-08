package router

// 路由文件
import (
	con "backend/controller"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 解决跨域问题
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的响应头
		AllowCredentials: true,                                                // 允许传递凭据（例如 Cookie）
		MaxAge:           12 * time.Hour,                                      // 预检请求的有效期
	}))
	// 设置静态文件目录
	r.Static("/dist", "./dist")
	// 测试GET请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//注册
	r.POST("/register", con.Register)
	//登录
	r.POST("/login", con.Login)
	//登出
	r.POST("/logout", con.Logout)
	//查询用户的类型
	r.POST("/getInfo", middleware.JWTAuthMiddleware(), con.GetInfo)
	//农产品上链
	r.POST("/uplink", middleware.JWTAuthMiddleware(), con.Uplink)
	// 获取农产品的上链信息
	r.POST("/getFruitInfo", con.GetFruitInfo)
	// 获取用户的农产品ID列表
	r.POST("/getFruitList", middleware.JWTAuthMiddleware(), con.GetFruitList)
	// 获取所有的农产品信息
	r.POST("/getAllFruitInfo", middleware.JWTAuthMiddleware(), con.GetAllFruitInfo)
	// 获取农产品上链历史(溯源)
	r.POST("/getFruitHistory", middleware.JWTAuthMiddleware(), con.GetFruitHistory)
	// 上传IPFS文件
	r.POST("/ipfsUpload", con.IpfsUpload)
	// 下载IPFS文件到服务器
	r.POST("/ipfsDownload", con.IpfsDownload)
	// 下载服务器文件
	r.GET("/download", con.Download)
	return r
}
