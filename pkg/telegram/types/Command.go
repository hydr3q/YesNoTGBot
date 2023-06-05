package types

type Command struct {
	Regex   string
	Handler func(*Context, []string)
}
