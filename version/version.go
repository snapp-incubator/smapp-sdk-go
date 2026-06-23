package version

import "fmt"

// Version is set to "dev" by default and overridden at build time via:
//   -ldflags "-X github.com/snapp-incubator/smapp-sdk-go/version.Version=vX.Y.Z"
var Version = "dev"

const UserAgentHeader = "User-Agent"

func GetUserAgent() string {
	return fmt.Sprintf("smapp-sdk-go/%s", Version)
}
