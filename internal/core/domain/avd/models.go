package avd

type CreateAVD struct {
	Name        string   `json:"name"`
	PackagePath string   `json:"packagePath"`
	Options     []string `json:"options"`
}

type DeleteAVD struct {
	Name string `json:"name"`
}
