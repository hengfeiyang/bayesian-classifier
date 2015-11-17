package storage

type RedisStorage struct {
}

// TODO: 还没想好用Redis的方法
// 可能存储的数据很大, 直接放到Redis中不好, 要变成小KEY
func NewRedistorage(config map[string]string) (*RedisStorage, error) {
	return &RedisStorage{}, nil
}

func (t *RedisStorage) Save(data interface{}) error {
	return nil
}

func (t *RedisStorage) Load(data interface{}) error {
	return nil
}
