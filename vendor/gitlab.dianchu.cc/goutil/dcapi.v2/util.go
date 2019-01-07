/*
Author : Haoyuan Liu
Time   : 2018/6/26
*/
package dcapi

import "reflect"

//Success 参数resp为struct 或者nil，设置并返回Response
func Success(resp interface{}) Response {
	response := NewResponse()
	if resp == nil {
		return response
	}
	response.SetCode(0)
	response.SetInfo("Success")
	val := reflect.ValueOf(resp)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i ++ {
		ft := typ.Field(i)
		fv := val.Field(i)
		response.Set(ft.Name, fv.Interface())
	}
	return response
}

//Failed 设置并返回Response
func Failed(code int, info string) Response {
	resp := NewResponse()
	resp.SetCode(code)
	resp.SetInfo(info)
	return resp
}

//FailedWithErr 设置并返回Response
func FailedWithErr(err error) Response {
	resp := NewResponse()
	resp.SetError(err)
	return resp
}
