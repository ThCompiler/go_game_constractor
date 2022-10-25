package scene

type Command uint64

const (
	NoCommand         = Command(0)
	ApplyStashedScene = Command(1)
	FinishScene       = Command(2)
)

type Button struct {
	Title   string
	URL     string
	Payload interface{}
}

type Text struct {
	BaseText     string
	TextToSpeech string
}

type Info struct {
	Text             Text
	Buttons          []Button
	ExpectedMessages []MessageMatcher
	Err              Error
}

type BaseTextError struct {
	Message string
}

func (bte BaseTextError) GetErrorText() string {
	return bte.Message
}

func (bte BaseTextError) GetErrorScene() Scene {
	return nil
}

func (bte BaseTextError) IsErrorScene() bool {
	return false
}

type BaseSceneError struct {
	Scene Scene
}

func (bse BaseSceneError) GetErrorText() string {
	return ""
}

func (bse BaseSceneError) GetErrorScene() Scene {
	return bse.Scene
}

func (bse BaseSceneError) IsErrorScene() bool {
	return true
}
