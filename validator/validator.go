package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans  ut.Translator
)

func init()  {
	// 注册翻译器
	translate := zh.New()
	uni := ut.New(translate, translate)

	trans, _ = uni.GetTranslator("zh")

	//获取校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

//翻译错误信息
func TranslateList(err error) map[string][]string  {

	var result = make(map[string][]string)

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}

	return result
}


func Translate(err error) string {

	var result = make(map[string][]string)

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}

	fmt.Println(err)
	fmt.Println(err.Error())

	return "123"

	//error := err.(validator.ValidationErrors)

	//return error
}