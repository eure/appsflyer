package util

import (
	"os"
)

func GetAPIToken() string {
	return os.Getenv("APPSFLYER_API_TOKEN")
}
