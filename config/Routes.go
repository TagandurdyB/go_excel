package config

import (
	controllers "auto_excel/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//Admin
	//Blog post
	r.GET("/", controllers.Dashboard{}.Index)
	r.GET("/new-add", controllers.Dashboard{}.NewItem)
	r.GET("/excel", controllers.Dashboard{}.Excel)
	r.POST("/add", controllers.Dashboard{}.Add)

	r.GET("/delete/:id", controllers.Dashboard{}.Delete)

	// r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	// r.POST("/admin/update/:id", admin.Dashboard{}.Update)
	// r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	//Userops
	// r.GET("/admin/login", admin.Userops{}.LoginIndex)
	// r.POST("/admin/do_login", admin.Userops{}.Login)
	// r.GET("/admin/sign-up", admin.Userops{}.SignUpIndex)
	// r.POST("/admin/do_sign-up", admin.Userops{}.SignUp)
	// r.POST("/admin/log_out", admin.Userops{}.LogOut)
	//Categories
	// r.GET("/admin/categories", admin.Categories{}.Index)
	// r.POST("/admin/categories/add", admin.Categories{}.Add)
	// r.GET("/admin/categories/delete/:id", admin.Categories{}.Delete)
	//Serve Files
	r.ServeFiles("/views/assets/*filepath", http.Dir("views/assets/"))
	r.ServeFiles("/.base/uploads/*filepath", http.Dir(".base/uploads/"))

	// r.ServeFiles("/netije.xlsx", http.Dir(".base//netije.xlsx"))

	// http.Handle("/admin/assets/*filepath", http.FileServer(http.Dir("v/assets")))
	// http.Handle("/uploads/*filepath", http.FileServer(http.Dir("uploads")))

	// static := httprouter.New()
	// static.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	// static.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	// r.NotFound = static

	return r
}
