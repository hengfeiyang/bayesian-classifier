package storage

type Storager interface {
	Save(data interface{}) error
	Load(data interface{}) error
}

func NewStorage(config map[string]string) (handler Storager, err error) {
	switch config["adapter"] {
	case "file":
		handler, err = NewFileStorage(config["path"])
	case "redis":
		handler, err = NewRedistorage(config)
	default:
		handler, err = NewFileStorage(config["path"])
	}
	return
}
