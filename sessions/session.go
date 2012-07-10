package sessions

type Session struct {
	ID     int
	Values map[string]interface{}
	Secret string
}

func (s Session) Get(key string) (v interface{}) {
	v = key

	return
}
