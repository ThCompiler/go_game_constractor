package scene

type Error struct {
	Base  string `yaml:"base,omitempty" json:"base,omitempty" xml:"base,omitempty"`
	Name  string `yaml:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Text  string `yaml:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	Scene string `yaml:"scene,omitempty" json:"scene,omitempty" xml:"scene,omitempty"`
}

func (e Error) IsBase() bool {
	return e.Base != ""
}

func (e Error) IsText() bool {
	return e.Text != "" && e.Name != ""
}

func (e Error) IsScene() bool {
	return e.Scene != ""
}

func (e Error) IsValid() bool {
	n := 0
	if e.IsBase() {
		n++
	}

	if e.IsScene() {
		n++
	}

	if e.IsText() {
		n++
	}

	return n == 1
}
