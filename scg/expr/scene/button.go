package scene

type payload interface{}

type Button struct {
	Name    string  `yaml:"name" xml:"name" json:"name"`
	URL     string  `yaml:"url,omitempty" xml:"url,omitempty" json:"url,omitempty"`
	Payload payload `yaml:"payload,omitempty" json:"payload,omitempty" xml:"payload,omitempty"`
}
