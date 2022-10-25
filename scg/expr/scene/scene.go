package scene

type Scene struct {
	Text        Text      `yaml:"text" json:"text" xml:"text"`
	NextScene   string    `yaml:"nextScene" json:"next_scene" xml:"nextScene"`
	IsInfoScene bool      `yaml:"isInfoScene,omitempty" json:"is_info_scene,omitempty" xml:"isInfoScene,omitempty"`
	Matchers    []Matcher `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
	Errors      []string  `yaml:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
}

func (s *Scene) IsValid() (bool, error) {
	return s.Text.IsValid()
}

type GoodByeScene struct {
	Scene
	Name string `yaml:"name" json:"name" xml:"name"`
}
