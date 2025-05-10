package version

import "fmt"

const (
	Version         = "v0.9.20"
	UserAgentHeader = "User-Agent"
)

func GetUserAgent() string {
	return fmt.Sprintf("smapp-sdk-go/%s", Version)
}
