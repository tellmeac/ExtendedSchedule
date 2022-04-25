package aggregates

import (
	"errors"
	"tellmeac/extended-schedule/domain/entity"
)

type Section struct {
	Position int
	Lessons  []entity.Lesson
}

func (section *Section) Join(other Section) error {
	if section.Position != other.Position {
		return errors.New("must be equal position number")
	}

	section.Lessons = JoinLessons(section.Lessons, other.Lessons)
	return nil
}

func JoinSections(x []Section, y []Section) []Section {
	var joinedResult = x
	for i := 0; i < len(x); i++ {
		joinedResult[i].Lessons = JoinLessons(joinedResult[i].Lessons, y[i].Lessons)
	}

	return joinedResult
}

func JoinLessons(a []entity.Lesson, b []entity.Lesson) []entity.Lesson {
	type lessonJoinKey struct {
		ID         string
		LessonType string
	}

	var joinMap = make(map[lessonJoinKey]*entity.Lesson)
	var key lessonJoinKey
	for i := 0; i < len(a); i++ {
		key = lessonJoinKey{
			ID:         a[i].ID,
			LessonType: a[i].LessonType,
		}
		joinMap[key] = &a[i]
	}
	for i := 0; i < len(b); i++ {
		key = lessonJoinKey{
			ID:         b[i].ID,
			LessonType: b[i].LessonType,
		}
		joinMap[key] = &b[i]
	}

	var joinedResult = make([]entity.Lesson, 0, len(joinMap))
	for _, lesson := range joinMap {
		joinedResult = append(joinedResult, *lesson)
	}

	return joinedResult
}
