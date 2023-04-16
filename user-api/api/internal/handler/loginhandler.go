package handler

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"

	zt "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/user-api/api/internal/logic"
	"go-zero-demo/user-api/api/internal/svc"
	"go-zero-demo/user-api/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		zhT := zh.New() // 中文翻译器
		uni := ut.New(zhT)
		trans, _ := uni.GetTranslator("zh")
		validate := validator.New()
		if err := zt.RegisterDefaultTranslations(validate, trans); err != nil {
			panic(err)
		}
		validate.RegisterValidation("CustomValidateFun", CustomValidateFun) // 自定义函数
		err := validate.StructCtx(r.Context(), req)
		if err != nil {

			errors, ok := err.(validator.ValidationErrors)
			if ok {
				fmt.Println(removeStructName(errors.Translate(trans)))
				for _, e := range errors.Translate(trans) {
					err := fmt.Errorf(e)
					httpx.ErrorCtx(r.Context(), w, err)
				}
			} else {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			return
		}
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

// 自定义函数
func CustomValidateFun(f1 validator.FieldLevel) bool {
	// f1 包含了字段相关信息
	// f1.Field() 获取当前字段信息
	// f1.Param() 获取tag对应的参数
	// f1.FieldName() 获取字段名称

	return f1.Field().String() == "lpj"
}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}
