package observer

import "fmt"

//每个obsever对外暴露一个接口，方便更新自己的状态
type Observer interface {
	Update(Subject)
}

type ConcreateObserver struct {
	state int
	name  string
}

func (c *ConcreateObserver) Update(s Subject) {
	c.state = s.GetState()
	fmt.Printf("%v observe %v\n", c.name, c.state)
}

func NewConcreateObserver(name string) *ConcreateObserver {
	return &ConcreateObserver{
		name: name,
	}
}

type Subject interface {
	Attach(Observer)
	Notify()
	SetState(state int)
	GetState() int
}

type ConcreateSubject struct {
	state   int
	objList []Observer
}

func NewConcreateSubject() *ConcreateSubject {
	return &ConcreateSubject{
		objList: make([]Observer, 0),
	}
}

func (s *ConcreateSubject) Attach(obj Observer) {
	s.objList = append(s.objList, obj)
}

func (s *ConcreateSubject) SetState(state int) {
	s.state = state
	s.Notify()
}

func (s *ConcreateSubject) GetState() int {
	return s.state
}

func (s *ConcreateSubject) Notify() {
	for _, v := range s.objList {
		v.Update(s)
	}
}
