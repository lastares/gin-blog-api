package initialize

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go-gin-blog-api/global"
	"reflect"
)

func Validate() {
	global.Validate = validator.New()
	uni := ut.New(zh.New())
	translator, _ := uni.GetTranslator("zh")

	// 注册一个翻译器
	err := zh_translations.RegisterDefaultTranslations(global.Validate, translator)
	if err != nil {
		fmt.Println(err)
	}
	//注册一个函数，获取struct tag里自定义的label作为字段名
	global.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	zh_translations.RegisterDefaultTranslations(global.Validate, translator)
	global.Translator = translator
}

