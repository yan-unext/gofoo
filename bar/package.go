package bar

import "github.com/redis/go-redis/v9"

func Foo() *redis.Client {
	return redis.NewClient(&redis.Options{})
}
