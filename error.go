package errtrait

// Err основная структура ошибки сервера.
// Эта структура ошибки будет возвращена клиенту.
type Err struct {
	Message string
	Code    string
	Trait   Trait
}

// Error возвращает сообщение об ошибке
func (e Err) Error() string {
	return e.Message
}

// GetTrait возвращает trait ошибки
func (e Err) GetTrait() Trait {
	return e.Trait
}
