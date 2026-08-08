package checkers

type Action int

const (
	ActionMin Action = iota

	ActionFetch

	ActionInsert
	ActionUpdate
	ActionDelete
)
