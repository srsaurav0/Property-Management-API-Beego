package test

import (
	"path/filepath"
	"runtime"

	_ "beego-api-service/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
// func TestBeego(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	logs.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 	        })
// 	        Convey("The Result Should Not Be Empty", func() {
// 	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 	        })
// 	})
// }
