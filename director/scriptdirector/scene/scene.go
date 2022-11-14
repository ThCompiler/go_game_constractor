package scene

import "github.com/ThCompiler/go_game_constractor/director"

type Command uint64

const (
	NoCommand         = Command(0)
	ApplyStashedScene = Command(1)
	FinishScene       = Command(2)
	StashScene        = Command(3)
)

type Info struct {
	Text             director.Text
	Buttons          []director.Button
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
