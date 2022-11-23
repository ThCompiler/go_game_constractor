package scene

type Payload interface{}

type Button struct {
	Text    string  `yaml:"text" xml:"text" json:"text"`
	URL     string  `yaml:"url,omitempty" xml:"url,omitempty" json:"url,omitempty"`
	Payload Payload `yaml:"payload,omitempty" json:"payload,omitempty" xml:"payload,omitempty"`
	ToScene string  `yaml:"toScene,omitempty" json:"to_scene,omitempty" xml:"toScene,omitempty"`
}
