package controllers

import (
	"fmt"
	//"plan_trabajo_docente_mid/models"	
	"github.com/astaxie/beego"
	"strconv"
	"reflect"
	"strings"
	"errors"
	"github.com/astaxie/beego/httplib"
	//"github.com/mitchellh/mapstructure"   
)

// SolicitudSoportePlanTrabajoController operations for SolicitudSoportePlanTrabajo
type SolicitudSoportePlanTrabajoController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudSoportePlanTrabajoController) URLMapping() {
	c.Mapping("ObtenerDocentesSolicitudes", c.ObtenerDocentesSolicitudes)
	c.Mapping("ObtenerInformacionDependencias", c.ObtenerInformacionDependencias)
}

// ObtenerDocentesSolicitudes
// @Title ObtenerDocentesSolicitudes
// @Description create ObtenerInfoCoordinador
// @Param	iddependencia	query 	string	true	"id de la dependencia"
// @Param	estado	query	string	true	"estado de la solicitud"
// @Param	anio	query	string	true	"anio de la solicitud"
// @Param	periodo	query	string	true	"semestre de la solicitud"
// @router /docentes_solicitudes [get]
func (c *SolicitudSoportePlanTrabajoController) ObtenerDocentesSolicitudes() {
	idDependencia := c.GetString("iddependencia")
	periodo := c.GetString("periodo")
	anio := c.GetString("anio")
	
	
	//var docente models.DocentePlanTrabajo
	var cedulasDocentes interface{}
	var docentes [] interface{}

	parametros := make(map[string]string)

	var estados string

	// orCondition para los estados: k:v,k:v
	if v := c.GetString("estados"); v != "" {
		for _, cond := range strings.Split(v, ",") {
		 	kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid orCondition key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]

			estados = estados + k + ":" +v+","
		}
	}

	fmt.Println(estados)
	strings.TrimSuffix(estados, ",")

	parametros["iddependencia"] = idDependencia
	parametros["estados"] = estados
	parametros["anio"] = anio
	parametros["periodo"] = periodo
	parametros["limit"] = "-1"

	if err := getJson(beego.AppConfig.String("UrlcrudPlan_trabajo_docente")+"/solicitud_soporte_plan_trabajo/obtener_cedulas/", &cedulasDocentes, parametros); err == nil {
		
		datosDocentes := cedulasDocentes.([]interface{})
				
		for _, cedulaColleccion := range datosDocentes{

			cedula := cedulaColleccion.(map[string]interface{})

			cedula["Persona"] = int(cedula["Persona"].(float64))
			var datosPersonas interface{}

			parametros = make(map[string]string)

			parametros["query"] = "Id:"+strconv.Itoa(cedula["Persona"].(int))
			parametros["fields"] = "Id,PrimerNombre,SegundoNombre,PrimerApellido,SegundoApellido"

			 if err := getJson(beego.AppConfig.String("UrlcrudAdministrativa")+"/informacion_persona_natural/", &datosPersonas, parametros); err == nil {
				if(reflect.TypeOf(datosPersonas) == nil){
					beego.Error("resultado de la consulta vacio la cedula "+strconv.Itoa(cedula["Persona"].(int))+" no existe")
					c.Abort("404")	
				}
				persona := datosPersonas.([]interface{})
				personaAux := persona[0].(map[string]interface{})
				docentes = append(docentes,personaAux)
			}else{
				beego.Error(err)
				c.Abort("404")				
			}		
		} 
	}else{
		beego.Error(err)
		c.Abort("404")
	}	
	c.Data["json"] = docentes
	c.ServeJSON()
}



// ObtenerInformacionDependencias
// @Title ObtenerInformacionDependencias
// @Description recibe el documento del decano o supervisor, busca de cuales dependencias es jefe
// y con los id's de esas dependencias trae informacion del nombre de oiko
// @Param	idsupervisor	query 	string	true	"id de la dependencia"
// @router /informacion_dependencias [get]
func (c *SolicitudSoportePlanTrabajoController) ObtenerInformacionDependencias() {
	idDependencia := c.GetString("idsupervisor")

	var dependencias interface{}
	//var dependenciasOikos interface{}

	var respuesta[] interface{}

	if err := getJson(beego.AppConfig.String("UrlcrudCore")+"/jefe_dependencia/?query=TerceroId:"+idDependencia, &dependencias, nil); err == nil {
		datosDependencias := dependencias.([]interface{})
		fmt.Println("______________________")
		fmt.Println(dependencias)
		fmt.Println(datosDependencias)
		fmt.Println("______________________")
		// for _, dependencia := range datosDependencias{
		// 	if err2 := getJson(beego.AppConfig.String("UrlcrudOikos")+"/dependencia?query=Id:"+strconv.Itoa(datosDependencias["DependenciaId"]), &dependenciasOikos, nil); err2 == nil {
		// 		datosDependenciasOikos := dependencias.([]interface{})		
		// 		dependenciaAux := datosDependenciasOikos[0].(map[string]interface{})
		// 		respuesta = append(respuesta, dependenciaAux)		
		// 	}else{
		// 		beego.Error(err)
		// 		c.Abort("404")				
		// 	}
		// }
	}else{
		beego.Error(err)
		c.Abort("404")
	}		
		c.Data["json"] = respuesta
		c.ServeJSON()
}

// ObtenerSolicitudesDocente
// @Title ObtenerSolicitudesDocente
// @Description recibe el documento del docente y devuelve informacion del soporte
// y con los id's de esas dependencias trae informacion del nombre de oiko
// @Param	cedula	query 	string	true	"cedula del docente"
// @Param	anio	query 	string	true	"cedula del docente"
// @Param	periodo	query 	string	true	"cedula del docente"
// @router /solicitudes_docente [get]
func (c *SolicitudSoportePlanTrabajoController) ObtenerSolicitudesDocente() {
	cedula := c.GetString("cedula")
	anio := c.GetString("anio")
	periodo := c.GetString("periodo")
	var planTrabajo []interface{}


	
	r := httplib.Get(beego.AppConfig.String("UrlcrudAcademica")+"consulta_plan_trabajo/"+cedula+"/"+anio+"/"+periodo)
		fmt.Println(r.String())	
	if err := r.ToJSON(&planTrabajo); err == nil {
			fmt.Println(r)
		}else{
			fmt.Println("error")
			fmt.Println(err)
		}
	

	// if err := getJsonWSO2(beego.AppConfig.String("UrlcrudAcademica")+"consulta_plan_trabajo/"+cedula+"/"+anio+"/"+periodo, &planTrabajo); err == nil {
	// 	fmt.Println(planTrabajo)
	// }else{
	// 	fmt.Println("no funciono")
	// 	fmt.Println(err)
	// }
}

