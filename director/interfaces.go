package director

type Director interface {
	PlayScene(command SceneRequest) Result
}
