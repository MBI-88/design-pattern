package cmd


type memento struct {
	serialized string
}

func (m memento) GetState() string {
	return m.serialized
}
