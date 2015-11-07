package main

/*
#include <ruby/ruby.h>

const char *
rstring_ptr(VALUE str) {
    return RSTRING_PTR(str);
}
int
rstring_len(VALUE str) {
    return RSTRING_LENINT(str);
}
void
rb_raise2(VALUE exc, const char *str) {
    rb_raise(exc, "%s", str);
}

void goobj_retain(void *);
void goobj_free(void *);

static const rb_data_type_t go_type = {
    "GoStruct",
    {NULL, goobj_free, NULL},
    0, 0, RUBY_TYPED_FREE_IMMEDIATELY|RUBY_TYPED_WB_PROTECTED
};

VALUE
NewGoStruct(VALUE klass, void *p)
{
    goobj_retain(p);
    return TypedData_Wrap_Struct((klass), &go_type, p);
}

void *
GetGoStruct(VALUE obj)
{
    void *val;
    return TypedData_Get_Struct((obj), void *, &go_type, (val));
}
*/
import "C"
import "unsafe"
import "fmt"

var objects = make(map[interface{}]bool)

func rb_define_method(klass C.VALUE, name string, fun unsafe.Pointer, args int) {
	cname := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&name)))[0]))
	C.rb_define_method(klass, cname, (*[0]byte)(fun), C.int(args))
}

func rb_define_singleton_method(klass C.VALUE, name string, fun unsafe.Pointer, args int) {
	cname := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&name)))[0]))
	C.rb_define_singleton_method(klass, cname, (*[0]byte)(fun), C.int(args))
}

func RbGoString(str C.VALUE) string {
	C.rb_string_value(&str)
	return C.GoStringN(C.rstring_ptr(str), C.rstring_len(str))
}

func RbBytes(bytes []byte) C.VALUE {
	if len(bytes) == 0 {
		return C.rb_str_new(nil, C.long(0))
	}
	cstr := (*C.char)(unsafe.Pointer(&bytes[0]))
	return C.rb_str_new(cstr, C.long(len(bytes)))
}

func RbString(str string) C.VALUE {
	if len(str) == 0 {
		return C.rb_utf8_str_new(nil, C.long(0))
	}
	cstr := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&str)))[0]))
	return C.rb_utf8_str_new(cstr, C.long(len(str)))
}

func rb_define_class(name string, parent C.VALUE) C.VALUE {
	cname := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&name)))[0]))
	v := C.rb_define_class(cname, parent)
	return v
}

func rb_raise(exc C.VALUE, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	cstr := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&str)))[0]))
	C.rb_raise2(exc, cstr)
}

func INT2NUM(n int) C.VALUE {
	return C.rb_int2inum(C.long(n))
}

func INT64toNUM(n int64) C.VALUE {
	return C.rb_ll2inum(C.longlong(n))
}

func StrSlice2RbArray(slice []string) C.VALUE {
	ary := C.rb_ary_new_capa(C.long(len(slice)))
	for _, val := range slice {
		C.rb_ary_push(ary, RbString(val))
	}
	return ary
}
