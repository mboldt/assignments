package datastore

import "github.com/mboldt/assignments/student"

type Datastore struct {
	students []*student.Student
}

func New() *Datastore {
	return &Datastore{students: []*student.Student{}}
}

func (d *Datastore) AddStudent(s *student.Student) {
	d.students = append(d.students, s)
}

func (d *Datastore) ListStudents() []*student.Student {
	return d.students
}