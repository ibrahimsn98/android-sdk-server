package domain

import "fmt"

var validSDKVersions = []string{
	"13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "19.0", "20.0",
	"latest",
}

func CheckSdkVersion(sdkVersion string) (string, error) {
	for _, v := range validSDKVersions {
		if v == sdkVersion {
			return sdkVersion, nil
		}
	}
	return "", fmt.Errorf("SDK version is not valid")
}
