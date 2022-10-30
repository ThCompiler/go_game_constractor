package director

// Director - interface type of game director
type Director interface {
	PlayScene(command SceneRequest) Result
}
