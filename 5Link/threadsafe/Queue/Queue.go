package Queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex //锁
}

//新建一个队列
func NewQueue() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)
	return queue
}

//每个方法都是添加上锁和解锁，保证线程安全，解决对同一个资源的同时访问问题

//解决了线程安全
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock() //延迟unlock
	return q.len
}

//解决了线程安全
func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) { //delete值
	q.lock.Lock()
	defer q.lock.Unlock()
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, el)
	q.len++

	return
}

func (q *Queue) Peek() interface{} { //get
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queue[0]
}
