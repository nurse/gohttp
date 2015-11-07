package main

/*
#include <ruby/ruby.h>

VALUE NewGoStruct(VALUE klass, void *p);
void *GetGoStruct(VALUE obj);

VALUE gohttp_hello(VALUE);
VALUE gohttpGet(VALUE,VALUE);
VALUE gohttpHead(VALUE,VALUE);
int hash_of_ary_of_str2values_i(VALUE,VALUE,VALUE);
VALUE gohttpPostForm(VALUE,VALUE,VALUE);
VALUE gohttp_get_url(VALUE,VALUE);
VALUE gourlQueryEscape(VALUE,VALUE);
VALUE gourlQueryUnescape(VALUE,VALUE);

VALUE goio_to_s(VALUE);
VALUE goioRead(VALUE);
VALUE goioClose(VALUE);

VALUE reqMethod(VALUE);
VALUE reqURL(VALUE);
VALUE reqProto(VALUE);
VALUE reqHeader(VALUE);
VALUE reqBody(VALUE);
VALUE reqContentLength(VALUE);
VALUE reqTransferEncoding(VALUE);
VALUE reqHost(VALUE);
VALUE reqForm(VALUE);
VALUE reqForm(VALUE);
VALUE reqForm(VALUE);
VALUE reqForm(VALUE);
VALUE reqRemoteAddr(VALUE);
VALUE reqRequestURI(VALUE);

VALUE respContentLength(VALUE);
VALUE respStatus(VALUE);
VALUE respStatusCode(VALUE);
VALUE respProto(VALUE);
VALUE respTransferEncoding(VALUE);
VALUE respHeader(VALUE);
VALUE respBody(VALUE);
VALUE respRequest(VALUE);
*/
import "C"

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "net/url"
import "reflect"
import "unsafe"

func main() {
}

var rb_cGoIO C.VALUE
var rb_cGohttp C.VALUE
var rb_cGourl C.VALUE
var rb_cRequest C.VALUE
var rb_cResponse C.VALUE

//export goobj_retain
func goobj_retain(obj unsafe.Pointer) {
	objects[obj] = true
}

//export goobj_free
func goobj_free(obj unsafe.Pointer) {
	delete(objects, obj)
}

func goioNew(klass C.VALUE, r interface{}) C.VALUE {
	v := reflect.ValueOf(r)
	return C.NewGoStruct(klass, unsafe.Pointer(&v))
}

//export goio_to_s
func goio_to_s(self C.VALUE) C.VALUE {
	v := (*reflect.Value)(C.GetGoStruct(self))
	switch v.Kind() {
	case reflect.Struct:
		return RbString(fmt.Sprintf("#<GoIO:%s>", v.Type().String()))
	case reflect.Chan, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return RbString(fmt.Sprintf("#<GoIO:%s:0x%X>", v.Type().String(), v.Pointer()))
	default:
		return RbString(fmt.Sprintf("#<GoIO:%s>", v.Type().String()))
	}
}

//export goioRead
func goioRead(self C.VALUE) C.VALUE {
	v := (*reflect.Value)(C.GetGoStruct(self))
	r := v.Interface().(io.Reader)
	body, err := ioutil.ReadAll(r)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	return RbBytes(body)
}

//export goioClose
func goioClose(self C.VALUE) C.VALUE {
	v := (*reflect.Value)(C.GetGoStruct(self))
	if c, ok := v.Interface().(io.Closer); ok {
		c.Close()
		return C.Qtrue
	}
	return C.Qfalse
}

func reqNew(klass C.VALUE, r *http.Request) C.VALUE {
	return C.NewGoStruct(klass, unsafe.Pointer(r))
}

//export reqMethod
func reqMethod(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return RbString(req.Method)
}

//export reqProto
func reqProto(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return RbString(req.Proto)
}

//export reqHeader
func reqHeader(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	h := C.rb_hash_new()
	for key, value := range req.Header {
		C.rb_hash_aset(h, RbString(key), StrSlice2RbArray(value))
	}
	return h
}

//export reqBody
func reqBody(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	if req.Body == nil {
		return C.Qnil
	}
	return goioNew(rb_cGoIO, req.Body)
}

//export reqContentLength
func reqContentLength(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return INT64toNUM(req.ContentLength)
}

//export reqTransferEncoding
func reqTransferEncoding(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return StrSlice2RbArray(req.TransferEncoding)
}

//export reqHost
func reqHost(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return RbString(req.Host)
}

//export reqRemoteAddr
func reqRemoteAddr(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return RbString(req.RemoteAddr)
}

//export reqRequestURI
func reqRequestURI(self C.VALUE) C.VALUE {
	req := (*http.Request)(C.GetGoStruct(self))
	return RbString(req.RequestURI)
}

func respNew(klass C.VALUE, r *http.Response) C.VALUE {
	return C.NewGoStruct(klass, unsafe.Pointer(r))
}

//export respContentLength
func respContentLength(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return INT64toNUM(resp.ContentLength)
}

//export respHeader
func respHeader(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	h := C.rb_hash_new()
	for key, value := range resp.Header {
		C.rb_hash_aset(h, RbString(key), StrSlice2RbArray(value))
	}
	return h
}

//export respStatus
func respStatus(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return RbString(resp.Status)
}

//export respStatusCode
func respStatusCode(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return INT2NUM(resp.StatusCode)
}

//export respProto
func respProto(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return RbString(resp.Proto)
}

//export respTransferEncoding
func respTransferEncoding(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return StrSlice2RbArray(resp.TransferEncoding)
}

//export respBody
func respBody(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	if resp.Body == nil {
		return C.Qnil
	}
	return goioNew(rb_cGoIO, resp.Body)
}

//export respRequest
func respRequest(self C.VALUE) C.VALUE {
	resp := (*http.Response)(C.GetGoStruct(self))
	return reqNew(rb_cRequest, resp.Request)
}

//export gohttp_hello
func gohttp_hello(dummy C.VALUE) C.VALUE {
	fmt.Printf("hello, world\n")
	return C.Qnil
}

//export gohttpGet
func gohttpGet(dummy C.VALUE, urlstr C.VALUE) C.VALUE {
	str := RbGoString(urlstr)
	resp, err := http.Get(str)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	return respNew(rb_cResponse, resp)
}

//export gohttpHead
func gohttpHead(dummy C.VALUE, urlstr C.VALUE) C.VALUE {
	str := RbGoString(urlstr)
	resp, err := http.Head(str)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	return respNew(rb_cResponse, resp)
}

//export hash_of_ary_of_str2values_i
func hash_of_ary_of_str2values_i(key C.VALUE, ary C.VALUE, arg C.VALUE) C.int {
	values := (*url.Values)(unsafe.Pointer(uintptr(arg)))
	values.Add(RbGoString(key), RbGoString(ary))
	return C.ST_CONTINUE
}

//export gohttpPostForm
func gohttpPostForm(dummy C.VALUE, urlstr C.VALUE, data C.VALUE) C.VALUE {
	str := RbGoString(urlstr)
	val := url.Values{}
	C.rb_hash_foreach(data, (*[0]byte)(C.hash_of_ary_of_str2values_i), C.VALUE(uintptr(unsafe.Pointer(&val))))
	resp, err := http.PostForm(str, val)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	return respNew(rb_cResponse, resp)
}

//export gohttp_get_url
func gohttp_get_url(dummy C.VALUE, urlstr C.VALUE) C.VALUE {
	u, err := url.Parse(RbGoString(urlstr))
	if err != nil {
		return C.Qnil
	}
	return RbString(u.String())
}

//export gourlQueryEscape
func gourlQueryEscape(dummy C.VALUE, s C.VALUE) C.VALUE {
	return RbString(url.QueryEscape(RbGoString(s)))
}

//export gourlQueryUnescape
func gourlQueryUnescape(dummy C.VALUE, s C.VALUE) C.VALUE {
	src := RbGoString(s)
	str, err := url.QueryUnescape(src)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s' is not valid pct-encoded", src)
	}
	return RbString(str)
}

//export Init_gohttp
func Init_gohttp() {
	sNew := "new"
	str_new := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&sNew)))[0]))

	rb_cGoIO = rb_define_class("GoIO", C.rb_cIO)
	C.rb_undef_alloc_func(rb_cGoIO)
	C.rb_undef_method(C.rb_class_of(rb_cGoIO), str_new)
	rb_define_method(rb_cGoIO, "to_s", C.goio_to_s, 0)
	rb_define_method(rb_cGoIO, "inspect", C.goio_to_s, 0)
	rb_define_method(rb_cGoIO, "read", C.goioRead, 0)
	rb_define_method(rb_cGoIO, "close", C.goioClose, 0)

	rb_cRequest = rb_define_class("Request", C.rb_cObject)
	C.rb_undef_alloc_func(rb_cRequest)
	C.rb_undef_method(C.rb_class_of(rb_cRequest), str_new)
	rb_define_method(rb_cRequest, "method", C.reqMethod, 0)
	//rb_define_method(rb_cRequest, "url", C.reqURL, 0)
	rb_define_method(rb_cRequest, "proto", C.reqProto, 0)
	rb_define_method(rb_cRequest, "header", C.reqHeader, 0)
	rb_define_method(rb_cRequest, "body", C.reqBody, 0)
	rb_define_method(rb_cRequest, "content_length", C.reqContentLength, 0)
	rb_define_method(rb_cRequest, "transfer_encoding", C.reqTransferEncoding, 0)
	rb_define_method(rb_cRequest, "host", C.reqHost, 0)
	//rb_define_method(rb_cRequest, "form", C.reqForm, 0)
	//rb_define_method(rb_cRequest, "post_form", C.reqForm, 0)
	//rb_define_method(rb_cRequest, "multipart_form", C.reqForm, 0)
	//rb_define_method(rb_cRequest, "trailer", C.reqForm, 0)
	rb_define_method(rb_cRequest, "remote_addr", C.reqRemoteAddr, 0)
	rb_define_method(rb_cRequest, "request_uri", C.reqRequestURI, 0)

	rb_cResponse = rb_define_class("Response", C.rb_cObject)
	C.rb_undef_alloc_func(rb_cResponse)
	C.rb_undef_method(C.rb_class_of(rb_cResponse), str_new)
	rb_define_method(rb_cResponse, "size", C.respContentLength, 0)
	rb_define_method(rb_cResponse, "status", C.respStatus, 0)
	rb_define_method(rb_cResponse, "status_code", C.respStatusCode, 0)
	rb_define_method(rb_cResponse, "proto", C.respProto, 0)
	rb_define_method(rb_cResponse, "transfer_encoding", C.respTransferEncoding, 0)
	rb_define_method(rb_cResponse, "header", C.respHeader, 0)
	rb_define_method(rb_cResponse, "body", C.respBody, 0)
	rb_define_method(rb_cResponse, "request", C.respRequest, 0)

	rb_cGohttp = rb_define_class("Gohttp", C.rb_cObject)
	rb_define_singleton_method(rb_cGohttp, "hello", C.gohttp_hello, 0)
	rb_define_singleton_method(rb_cGohttp, "get", C.gohttpGet, 1)
	rb_define_singleton_method(rb_cGohttp, "head", C.gohttpHead, 1)
	//rb_define_singleton_method(rb_cGohttp, "post", C.gohttpPost, 1)
	rb_define_singleton_method(rb_cGohttp, "post_form", C.gohttpPostForm, 2)
	//rb_define_method(rb_cGohttp, "initialize", C.gohttp_initialize, 1)
	rb_define_method(rb_cGohttp, "get_url", C.gohttp_get_url, 1)

	rb_cGourl = rb_define_class("Gourl", C.rb_cObject)
	rb_define_singleton_method(rb_cGourl, "query_escape", C.gourlQueryEscape, 1)
	rb_define_singleton_method(rb_cGourl, "query_unescape", C.gourlQueryUnescape, 1)
}
