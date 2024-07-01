package example

import (
	"errors"
	"net/http"

	"github.com/binamico/errtrait"
)

// TraitErr - интерфейс, который определяет методы для получения ошибки в виде строки и
// получения связанного с ней признака (Trait).
type TraitErr interface {
	error
	Trait() errtrait.Trait
}

//nolint:cyclop // потому что switch тут будет быстрее, чем map
func errorInterpreter(err error) (int, error) {
	var traitErr TraitErr
	if ok := errors.As(err, &traitErr); ok {
		switch traitErr.Trait() {
		case errtrait.Validation:
			return http.StatusBadRequest, err
		case errtrait.Internal:
			return http.StatusInternalServerError, err
		case errtrait.NotFound:
			return http.StatusNotFound, err
		case errtrait.Conflict:
			return http.StatusConflict, err
		case errtrait.UnAuthorized:
			return http.StatusUnauthorized, err
		case errtrait.Forbidden:
			return http.StatusForbidden, err
		case errtrait.BadRequest:
			return http.StatusBadRequest, err
		default:
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusInternalServerError, errDefault(err)
	}
}

// errDefault - отдает ошибку по-умолчанию
func errDefault(err error) error {
	return &errtrait.Err{
		Message: err.Error(),
		Code:    "INTERNAL_ERROR",
	}
}
