package valid

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Error struct {
	Key     string
	Message string
}
type Errors []*Error

func (e *Error) Error() string {
	return e.Message
}

func (e Errors) Error() string {
	return strings.Join(e.Errors(), ",")
}

func (e Errors) Errors() []string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndVaild(c *gin.Context, v interface{}) (bool, Errors) {
	var errs Errors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &Error{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
