package entity

type State string

const (
	StateNone           State = "none"
	StateAwaitingName   State = "name"
	StateAwaitingNumber State = "number"
)

type User struct {
	UserID int64
	Name   string
	State  State
	Number string
}
