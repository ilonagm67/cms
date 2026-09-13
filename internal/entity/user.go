package entity

const (
	StateAwaitingUserName     State = "user_name"
	StateAwaitingUserPhone    State = "user_phone"
	StateAwaitingUserQuestion State = "user_question"
)

type User struct {
	ID     int64
	Name   string
	Number string
}
