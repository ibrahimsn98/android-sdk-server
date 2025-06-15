package adb

type StopDevice struct {
	DeviceSerial string `json:"deviceSerial"`
}

type RestartDevice struct {
	DeviceSerial string `json:"deviceSerial"`
}
