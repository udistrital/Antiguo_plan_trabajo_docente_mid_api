package controllers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	//"math/big"
	"net/http"
	"net/url"
	//"reflect"
	//"strings"
	//"time"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	//"github.com/udistrital/plan_trabajo_docente_mid/models"
	//"plan_trabajo_docente_mid/models"
)

func sendJson(urlTarget string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		json.NewEncoder(b).Encode(datajson)
	}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, urlTarget, b)
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//se agrega funcionalidad de recibir paramaetros
func getJson(urlTarget string, target interface{}, parametros map[string]string ) error {

	if(parametros != nil){
		param := url.Values{}

		for k, v := range parametros {
			param.Add(k,v)
		}
	
		urlTarget = urlTarget + "?" +param.Encode()
	}	

		fmt.Println(urlTarget)

	r, err := http.Get(urlTarget)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&target)
}


func getJsonMap(urlTarget string, target *interface{}, parametros map[string]string) error {

	if(parametros != nil){
		param := url.Values{}
		for k, v := range parametros {
			param.Add(k,v)
		}
		urlTarget = urlTarget + "?" +param.Encode()
	}	
	response, err := http.Get(urlTarget)
	if err != nil {
		return err
	}
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
	}

	err = json.Unmarshal(responseData, target)
	if err != nil {
        return err
	}
	return err
}


func getXml(urlTarget string, target interface{}) error {
	r, err := http.Get(urlTarget)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return xml.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// 	_, err := http.Get("https://golang.org/")
	
	
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		// fmt.Println("_________________________")
		// body, err := ioutil.ReadAll(r.Body)
		// fmt.Println("_________________________")
		return err
	}else{
		fmt.Println("__________err_______________")
		fmt.Println(err)
		fmt.Println("___________err______________")
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

