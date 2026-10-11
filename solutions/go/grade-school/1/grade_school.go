package gradeschool

import (
    "maps"
    "slices"
)

type School struct {
    studentsPerGrade map[int][]string
	allStudentsRoster map[string]struct{}
}

func New() *School {
	return &School{
   		studentsPerGrade:  map[int][]string{},
   		allStudentsRoster: map[string]struct{}{},
   	}
}

func (s *School) Add(student string, grade int) bool {
	if _, existsInRoster := s.allStudentsRoster[student]; existsInRoster {
        return false
    }
    
	s.studentsPerGrade[grade] = append(s.studentsPerGrade[grade], student)
    s.allStudentsRoster[student] = struct{}{}

    return true
}

func (s *School) Grade(level int) []string {
	students := slices.Clone(s.studentsPerGrade[level])

	if students == nil {
		return []string{}
	}

	slices.Sort(students)
	return students
}

func (s *School) Enrollment() []string {
	ret := []string{}

    for _, grade := range slices.Sorted(maps.Keys(s.studentsPerGrade)) {
        students := s.Grade(grade)
        ret = append(ret, students...)
    }

    return ret
}
