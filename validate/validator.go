package validate

import (
	"fmt"
	"reflect"
	"strings"
)

var globalValidator *ValidatorCollector

func init() {
	globalValidator = newValidatorCollector()
}

func newValidatorCollector() *ValidatorCollector {
	return &ValidatorCollector{validations: make(map[string]Validator)}
}

func GetValidator() *ValidatorCollector {
	return globalValidator
}

type Validator struct {
	ValidateFunc func(value reflect.Value) bool
}

type ValidatorCollector struct {
	validations map[string]Validator
}

func (v *ValidatorCollector) RegisterValidator(name string, validatorFunc func(value reflect.Value) bool) {
	validator := Validator{ValidateFunc: validatorFunc}
	v.validations[name] = validator
}

func (v *ValidatorCollector) Run(name string, value reflect.Value) bool {
	if f, ok := v.validations[name]; ok {
		return f.ValidateFunc(value)
	}
	panic("未实现validator:" + name)
}

func ToValidateStruct(data any) bool {
	dataValue := reflect.ValueOf(data)
	fmt.Println(dataValue.Kind())
	if dataValue.Kind() == reflect.Ptr && dataValue.Elem().Kind() == reflect.Struct {
		dataValue = dataValue.Elem()
	} else if dataValue.Kind() != reflect.Struct {
		panic("must be *struct or struct")
	}

	for i := 0; i < dataValue.NumField(); i++ {
		// 验证值
		if !toValidateField(dataValue, i) {
			return false
		}
	}

	return true
}

func toValidateField(value reflect.Value, index int) bool {
	filed := value.Type().Field(index)
	// 获取对应字段的目标标签
	targetTag := filed.Tag.Get(TagName)
	fmt.Println(targetTag)

	if targetTag == "" { // 没有标签直接判断下一个字段
		return true
	}

	targetTagValues := strings.Split(targetTag, ";")

	// 验证值
	for i := 0; i < len(targetTagValues); i++ {
		if !globalValidator.Run(targetTagValues[i], value.Field(index)) {
			return false
		}
	}
	return true
}
