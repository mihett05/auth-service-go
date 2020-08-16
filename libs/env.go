package libs

import "os"

var envDefault = map[string]string {
	"KEY": "key",
	"PORT": "8000",
	"DATABASE_URL": "postgres://postgres:1@localhost:5432/auth_go",
}

func EnvDefault(key string) string {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key)
	}
	return envDefault[key]
}
