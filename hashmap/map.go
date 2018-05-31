package hashmap

type Map interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
	Del(string)
	Len() int
	Cap() int
}
