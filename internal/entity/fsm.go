package entity

type State string

const (
	StateAwaitingName  State = "name"
	StateAwaitingPhone State = "phone"
)

type FSM struct {
	UserID int64
	State  State
}
