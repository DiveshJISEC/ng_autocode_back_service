package repository

type transientCache struct {
	Id        string
	DataCache map[string]interface{}
}
