package system

import (
	"errors"
	"os"
)

func FindAndroidSDKPath() (string, error) {
	if sdkPath := os.Getenv("ANDROID_HOME"); sdkPath != "" {
		return sdkPath, nil
	}
	if sdkPath := os.Getenv("ANDROID_SDK_ROOT"); sdkPath != "" {
		return sdkPath, nil
	}
	return "", errors.New("no android sdk path")
}
