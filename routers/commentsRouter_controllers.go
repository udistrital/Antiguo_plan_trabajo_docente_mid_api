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

	beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "ObtenerInformacionDependencias",
			Router: `/informacion_dependencias`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_mid/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "ObtenerSolicitudesDocente",
			Router: `/solicitudes_docente`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
