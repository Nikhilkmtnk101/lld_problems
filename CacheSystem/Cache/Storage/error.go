package Storage

type Error string

func (error Error) Error() string {
	return string(error)
}

const (
	KeyDoesNotExistErr Error = "Key does not exist"
	FullStorageErr     Error = "Cache Storage is full"
)
