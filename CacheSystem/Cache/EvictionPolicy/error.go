package EvictionPolicy

type Error string

func (error Error) Error() string {
	return string(error)
}

const (
	NodeDoesNotExistErr Error = "Node does not exist"
)
