package errtrait

import (
	"errors"
)

// Err основная структура ошибки сервера.
// Эта структура ошибки будет возвращена клиенту.
type Err struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Trait   Trait  `json:"-"`
}

// Error возвращает сообщение об ошибке
func (e Err) Error() string {
	return e.Message
}

// GetTrait возвращает trait ошибки
func (e Err) GetTrait() Trait {
	return e.Trait
}

// HasTrait проверяет, имеется ли у ошибки переданный trait.
func HasTrait(err error, trait Trait) bool {
	var traitErr Err
	if ok := errors.As(err, &traitErr); !ok {
		return false
	}

	return traitErr.Trait == trait
}
