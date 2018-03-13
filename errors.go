package main

import (
	lane "gopkg.in/oleiade/lane.v1"
	"errors"
	"fmt"
)

type Error struct {
	Err   error
	Cause string
}

var (
	queue = lane.NewQueue()
)

type Errors struct{}

func Queue() *lane.Queue {
	if queue == nil {
		queue = lane.NewQueue()
	}

	return queue
}

func (e *Errors) Push(err error, cause string){
	q := Queue()

	er := Error{
		Err: err,
		Cause: cause,
	}

	q.Enqueue(er)

	return

}

func (e Errors) Has() bool {
	q := Queue()

	if q.Size() >0 {
		return true
	}

	return false
}

func (e *Errors) First() Error {
	q := Queue()

	v := q.First()
	if v !=nil {
		er := v.(Error)
		return er
	}

	return Error{}

}

func (e *Errors ) Count() int {
	q := Queue()

	return q.Size()
}

func (e *Errors ) DequeueAll() []Error  {
	q := Queue()
	ers := make([]Error, 0)

	t := q.Size()

	for i:=0; i<t; i++ {
		er := q.Dequeue()
		err := er.(Error)
		ers = append(ers, err)
	}

	return ers
}

func div(e *Errors, a, b int)  (int, bool){
	if e.Has() {
		return 0, true
	}

	if a == 0 || b == 0 {
		err := errors.New("division by zero")
		e.Push(err, "Zero value exception")
		return 0, true
	}

	result := a/b
	return result, false
}

func mul(e *Errors, a, b int)  (int, bool){
	if e.Has() {
		return 0, true
	}

	if a == 0 || b == 0 {
		err := errors.New("multiplication by zero")
		e.Push(err, "Zero mult exception")
		return 0, true
	}

	result := a*b
	return result, false
}

func main() {

	e := Errors{}

	d, _ := div(&e, 0, 3)
	m, _ := mul(&e, 0, d)

	if e.Has() {
		fmt.Println(e.First().Cause)
		return
	}

	fmt.Println(d + m)
}
