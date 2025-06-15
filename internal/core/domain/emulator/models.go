package emulator

type Start struct {
	AVDName string   `json:"avdName"`
	Args    []string `json:"args"`
}
