package version

import "fmt"

const (
	Version         = "v0.5.0"
	UserAgentHeader = "User-Agent"
)

func GetUserAgent() string {
	return fmt.Sprintf("smapp-sdk-go/%s", Version)
}
