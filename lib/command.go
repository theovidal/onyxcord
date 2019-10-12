package lib

type Command struct {
	Name string
	Description string
	Usage string
	Category string
	Execute interface{}
}