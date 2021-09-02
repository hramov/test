package ports

type CachePort interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}
