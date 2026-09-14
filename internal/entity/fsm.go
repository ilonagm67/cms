package entity

type State string

type FSM struct {
	UserID     int64
	State      State
	DataName   string
	DataWeight int
}
