package hashmap

type Map interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
	Len() int
	Cap() int
}
