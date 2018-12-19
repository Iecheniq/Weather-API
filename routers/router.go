// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"weather/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/weather",
		beego.NSInclude(&controllers.WeatherController{}),
<<<<<<< HEAD
		beego.NSNamespace("/scheduler",
			beego.NSInclude(&controllers.SchedulerController{})),
=======
>>>>>>> v2
	)
	beego.AddNamespace(ns)
}
