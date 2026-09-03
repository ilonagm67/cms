package entity

type State string

const (
	StateNone State = "none"
)

type User struct {
	UserID int64
	Name   string
	State  State
	Number string
}
