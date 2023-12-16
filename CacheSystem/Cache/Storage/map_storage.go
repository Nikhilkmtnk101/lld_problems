package Storage

type MapStorage struct {
	capacity int
	storage  map[string]interface{}
}

func NewMapStorage(capacity int) Storage {
	return &MapStorage{
		capacity: capacity,
		storage:  make(map[string]interface{}),
	}
}

func (m *MapStorage) Add(key string, value interface{}) error {
	if len(m.storage) == m.capacity {
		return FullStorageErr
	}
	m.storage[key] = value
	return nil
}

func (m *MapStorage) Get(key string) (interface{}, error) {
	if value, exists := m.storage[key]; exists {
		return value, nil
	} else {
		return nil, KeyDoesNotExistErr
	}
}

func (m *MapStorage) Delete(key string) error {
	if _, exists := m.storage[key]; exists {
		delete(m.storage, key)
		return nil
	}
	return KeyDoesNotExistErr
}
