package toy

type LocalMapStore struct {
	storage map[string]string
}

func NewLocalMapStore() Storage {
	return &LocalMapStore{
		storage: make(map[string]string),
	}
}

func (lms *LocalMapStore) RetiveData(token string) (string, error) {
	return lms.storage[token], nil
}

func (lms *LocalMapStore) StoreData(token string, data string) error {
	lms.storage[token] = data
	return nil
}
