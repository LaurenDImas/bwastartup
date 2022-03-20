package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be email format"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func FormatValidationError(err error) []ErrorMsg {
	// var errors []string
	coral := []string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	fmt.Println(coral)

	out := make([]ErrorMsg, len(err.(validator.ValidationErrors)))

	for i, fe := range err.(validator.ValidationErrors) {
		out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
	}

	return out
}
