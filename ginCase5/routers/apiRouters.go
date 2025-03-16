package routers

import (
	"github.com/gin-gonic/gin"
	"golang_Gin/ginCase5/controllers/api"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", (&api.ApiController{}).Index)
		apiRouters.GET("/userlist", (&api.ApiController{}).UserList)

	}
}
