package cache

type Cache interface {
	Read(key string) (interface{}, error)
	Write(key string, obj interface{}, timeoutSeconds int32) error
}
