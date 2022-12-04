package validate

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestValidator(t *testing.T) {
	type A struct {
		Name string `thy:"mobile"`
	}

	v := GetValidator()
	v.RegisterValidator("mobile", func(value reflect.Value) bool {
		v := value.String()
		fmt.Println(v)
		ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, v)
		return ok
	})

	a := A{Name: "13512345678"}
	res := ToValidateStruct(a)
	if !res {
		t.Error("a测试失败")
	}

	a1 := A{Name: "14212345678"}
	res1 := ToValidateStruct(a1)
	if res1 {
		t.Error("a1测试失败")
	}
}
