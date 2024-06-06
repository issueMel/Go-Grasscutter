package player

type SceneLoadState uint32

const (
	NONE    SceneLoadState = 0
	LOADING SceneLoadState = 1
	INIT    SceneLoadState = 2
	LOADED  SceneLoadState = 3
)
