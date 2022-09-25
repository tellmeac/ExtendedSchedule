package schedule

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"tellmeac/extended-schedule/common/errors"
	"tellmeac/extended-schedule/schedule"
	"time"
)

type FacultyProvider interface {
	Faculties(c context.Context) ([]schedule.Faculty, error)
}

type Provider interface {
	GetByTeacher(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error)
	GetByGroup(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error)
}

func NewServerHandler(f FacultyProvider, p Provider, b schedule.Builder) *ServerHandler {
	return &ServerHandler{
		faculties: f,
		provider:  p,
		builder:   b,
	}
}

type ServerHandler struct {
	faculties FacultyProvider
	provider  Provider
	builder   schedule.Builder
}

func (s ServerHandler) GetGroups(c *gin.Context, params GetGroupsParams) {
	//TODO implement me
	panic("implement me")
}

func (s ServerHandler) GetTeachers(c *gin.Context, params GetTeachersParams) {
	//TODO implement me
	panic("implement me")
}

func (s ServerHandler) GetScheduleByGroupId(c *gin.Context, id string, params GetScheduleByGroupIdParams) {
	result, err := s.provider.GetByGroup(c, id, params.From.Time, params.To.Time)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s ServerHandler) GetScheduleByTeacherId(c *gin.Context, id string, params GetScheduleByTeacherIdParams) {
	result, err := s.provider.GetByTeacher(c, id, params.From.Time, params.To.Time)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s ServerHandler) GetUsersSchedule(c *gin.Context, params GetUsersScheduleParams) {
	result, err := s.builder.Personal(c, string(params.Email), params.From.Time, params.To.Time)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s ServerHandler) GetLessonsByGroupId(c *gin.Context, id string, params GetLessonsByGroupIdParams) {
	groupSchedule, err := s.provider.GetByGroup(c, id, params.From.Time, params.To.Time)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, uniqueLessons(groupSchedule))
}

func uniqueLessons(groupSchedule schedule.Schedule) []LessonWithoutContext {
	allLessons := make([]schedule.Lesson, 0)
	for _, day := range groupSchedule.Days {
		allLessons = append(allLessons, day.Lessons...)
	}

	uniqueLessons := lo.UniqBy(allLessons, func(l schedule.Lesson) string {
		return l.ID
	})

	return lo.Map(uniqueLessons, func(l schedule.Lesson, _ int) LessonWithoutContext {
		return LessonWithoutContext{
			Groups:  l.Groups,
			ID:      l.ID,
			Kind:    l.Kind,
			Name:    l.Name,
			Teacher: l.Teacher,
		}
	})
}
