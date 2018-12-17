package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "ObtenerDocentesSolicitudes",
			Router: `/docentes_solicitudes`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
