package datastore

import "github.com/mboldt/assignments/student"

type Datastore struct {
	Students []*student.Student `json:"students"`
}

func New() *Datastore {
	return &Datastore{Students: []*student.Student{}}
}

func (d *Datastore) AddStudent(s *student.Student) {
	d.Students = append(d.Students, s)
}

func (d *Datastore) ListStudents() []*student.Student {
	return d.Students
}
