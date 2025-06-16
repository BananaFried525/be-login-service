package cache

import "fmt"

const (
	AnonymousToken = "token:anonymous"
)

func GetAnonymousTokenKey(key string) string {
	return fmt.Sprintf("%s:%s", AnonymousToken, key)
}
