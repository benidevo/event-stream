package interfaces

type ResponseMessage interface {
	Serialize() ([]byte, error)
}
