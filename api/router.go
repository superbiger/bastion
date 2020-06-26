package api

import (
	"bastion/controller"
	"bastion/internal/middleware"
	"bastion/models"
	"bastion/pkg"
	"bastion/pkg/pprof"
	"encoding/gob"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"html/template"
)

func init() {
	// sessions encode
	gob.Register(models.MUser{})
	gob.Register(models.StatAdmin{})
}

func Init() *gin.Engine {
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))

	// 全局中间件
	r.Use(
		middleware.Recovery(),
		middleware.RequestLog(),
		middleware.Translation(),
		sessions.Sessions("bastion", store))

	// 模版
	r.SetFuncMap(template.FuncMap{
		"fmtDate":   pkg.FmtDate,
		"parseHtml": pkg.ParseHtml,
	})
	r.LoadHTMLGlob("web/views/*")

	// 静态资源
	r.Static("/web/public", "web/public")
	r.StaticFile("favicon.ico", "web/public/favicon.ico")

	// swagger doc
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// prometheus 监控
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	// pprof
	pprof.Register(r)

	// system
	system := r.Group("/")
	SysRegister(system)

	// 统计
	stat := r.Group("/")
	StatRegister(stat)

	return r
}

func SysRegister(r *gin.RouterGroup) {
	s := controller.System{}
	r.GET("/", s.Info)
	r.GET("/sys/error", s.Error)
}

func StatRegister(r *gin.RouterGroup) {
	m := controller.Monitor{}

	// https://github.com/gin-contrib/cors/issues/29
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Accept-Encoding",
			"Authorization", "Cache-Control", "X-Requested-With", "User-Agent", "Referrer", "Host", "Token"},

		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	r.POST("/api/stat/admin/login", m.AdminLogin)
	r.POST("/api/stat/admin/create", m.AdminCreate)
	r.POST("/api/stat/admin/update", middleware.StatAdminAuth(), m.AdminUpdate)
	r.GET("/api/stat/admin/info", middleware.StatAdminAuth(), m.AdminInfo)
	r.GET("/api/stat/admin/list", middleware.StatAdminAuth(), m.AdminList)

	r.POST("/api/stat/project", middleware.StatAdminAuth(), m.CreateProject)
	r.GET("/api/stat/projects", middleware.StatAdminAuth(), m.FindAllProjects)
	r.GET("/api/stat/project/:id", middleware.StatAdminAuth(), m.FindProjectById)

	r.POST("/api/stat/error", m.CreateError)
	r.GET("/img/stat/error", m.ImgCreateError)

	r.GET("/api/stat/device", m.FindDeviceByUid)
	r.GET("/api/stat/devices", m.FindAllDevice)
	r.GET("/api/stat/errors", m.FindErrorsWithParams)

	r.POST("/api/stat/behavior", m.CreateBehavior)
	r.GET("/api/stat/behaviors", middleware.StatAdminAuth(), m.FindAllBehaviors)
	r.GET("/api/stat/behavior/:id", middleware.StatAdminAuth(), m.FindBehaviorById)

	r.GET("/api/test/fail", m.TestFail)
	r.GET("/api/test/error", m.TestError)
	r.GET("/api/test/timeout", m.TestTimeOut)
}
