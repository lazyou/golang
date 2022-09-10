package main

import (
    "fmt"
)

// student 行为
type StudentTalk interface {
    talk()
}

// teacher 行为
type TeacherTalk interface {
    say()
}

// people 行为
type PeopleTalk interface {
    StudentTalk
    TeacherTalk
}

type Person struct {
    TeacherTalk
    StudentTalk
}

// Struct Student 和 Teacher 分别实现了 StudentTalk 和 TeacherTalk。
type Student struct{}

func (s *Student) talk() {
    fmt.Println("student talk")
}

type Teacher struct{}

func (t *Teacher) say() {
    fmt.Println("teacher say")
}

// meet 的参数 PeopleTalk，而且实现了 PeopleTalk 必然实现了 StudentTalk 和 TeacherTalk
func meet(p PeopleTalk)  {
    fmt.Println("====>people meet<====")
    meetTeacher(p)
    meetStudent(p)
}

func meetTeacher(ps TeacherTalk)  {
    fmt.Println("====>teacher meet<====")
    ps.say()
}

func meetStudent(ps StudentTalk) {
    fmt.Println("====>student meet<====")
    ps.talk()
}

func main() {
    t := Teacher{}
    s := Student{}

    // Person 实现了 PeopleTalk 方法，通过 Teacher 和 Student 实例
    p := Person{
        TeacherTalk: &t, 
        StudentTalk: &s,
    }
    
    meet(p)
}
