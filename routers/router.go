package routers

import (
	"beego-api-service/controllers"
	_ "beego-api-service/docs"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/v1/api",
		web.NSNamespace("/property",
			web.NSRouter("/details/:propertyId", &controllers.PropertyDetailsController{}, "get:GetPropertyDetails"),
			web.NSRouter("/gallery/:propertyId", &controllers.PropertyImagesController{}, "get:GetPropertyImages"),
		),
		web.NSRouter("/propertyList", &controllers.BulkPropertyFetchController{}, "get:BulkPropertyFetch"),
		web.NSNamespace("/user",
			web.NSRouter("/", &controllers.UserController{}, "post:CreateUser"),
			web.NSRouter("/:identifier", &controllers.UserController{}, "get:GetUser"),
			web.NSRouter("/:identifier", &controllers.UserController{}, "put:UpdateUser"),
			web.NSRouter("/:identifier", &controllers.UserController{}, "delete:DeleteUser"),
		),
	)

	web.AddNamespace(ns)

	web.Router("/swagger/*", &controllers.SwaggerController{})
}
