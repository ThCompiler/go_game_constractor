package scene

//go:generate mockgen -destination=mocks/mock_scene.go -package=mock_scene -mock_names=Scene=MockScene . Scene
//go:generate mockgen -destination=mocks/mock_error.go -package=mock_scene -mock_names=Error=MockError . Error
//go:generate mockgen -destination=mocks/mock_message_matcher.go -package=mock_scene -mock_names=MessageMatcher=MockMessageMatcher . MessageMatcher

type Error interface {
	GetErrorText() string
	GetErrorScene() Scene
	IsErrorScene() bool
}

type Scene interface {
	GetSceneInfo(ctx *Context) (sceneInfo Info, withReact bool)
	React(ctx *Context) Command
	Next() Scene
}

type MessageMatcher interface {
	Match(message string) (isMatch bool, searchedString string)
	GetMatchedName() string
}
