package singleton

type Singletion struct {
}

var singleton *Singletion

func init() {
	singleton = &Singletion{}
}
func GetInstance() *Singletion {
	return singleton
}
