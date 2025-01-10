package routers

import (
	"beego-api-service/controllers"

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
			web.NSRouter("/", &controllers.CreateController{}, "post:CreateUser"),
			// web.NSRouter("/:identifier", &controllers.UserController{}, "get:Get;put:Update;delete:Delete"),
		),
	)

	web.AddNamespace(ns)
}
