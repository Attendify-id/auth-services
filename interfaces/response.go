package interfaces

type ResponseJSON struct {
	Status  bool
	Message string
	Data    interface{}
	Error   interface{}
}
