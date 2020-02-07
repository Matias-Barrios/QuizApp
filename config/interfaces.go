package config

// IenvironmentFetcher : Interface type for environment fetcher
type IenvironmentFetcher interface {
	GetValue(key string) (string, error)
}
