package interfaces

type CacheRepository interface {
	Set(key string, entry []byte) error
	Get(key string) ([]byte, error)
}
