package memory

type StateRepository struct {
	states map[int64]string
}

func NewStateRepository() *StateRepository {
	return &StateRepository{
		states: make(map[int64]string, 0),
	}
}
