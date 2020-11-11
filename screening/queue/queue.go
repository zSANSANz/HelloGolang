package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue {
	var q Queue
	q.push()
	q.pop()
	q.Contains()
	q.Len()
	q.Keys()
	return q
}
