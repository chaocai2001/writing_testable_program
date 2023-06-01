package storage

type RedisClient struct {
}

func NewRedisClient() *RedisClient {
	panic("failed to connect to the redis server.")
}

func (mem *RedisClient) RetiveData(token string) (string, error) {

	return "", nil
}

func (mem *RedisClient) StoreData(token string, data string) error {

	return nil
}
