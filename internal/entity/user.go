package entity

type State string

const (
	StateNone        State = "none"
	StateWaitingName State = "name"
)

type User struct {
	ID     int64
	Name   string
	State  State
	Number string
}
