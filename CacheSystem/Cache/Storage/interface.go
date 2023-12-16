package Storage

type Storage interface {
	Add(string, interface{}) error
	Get(string) (interface{}, error)
	Delete(string) error
}
