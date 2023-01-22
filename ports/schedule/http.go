package schedule

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"tellmeac/extended-schedule/pkg/errors"
	"tellmeac/extended-schedule/schedule"
	"time"
)

type Provider interface {
	GetByTeacher(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error)
	GetByGroup(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error)
}

type TeacherProvider interface {
	Search(ctx context.Context, filter string, limit int) ([]schedule.Teacher, error)
}

type GroupProvider interface {
	Search(ctx context.Context, filter string, limit int) ([]schedule.StudyGroup, error)
}

var _ ServerInterface = ServerHandler{}

func NewServerHandler(p Provider, b schedule.Builder, tp TeacherProvider, gp GroupProvider) *ServerHandler {
	return &ServerHandler{
		provider: p,
		builder:  b,
		teachers: tp,
		groups:   gp,
	}
}

type ServerHandler struct {
	provider Provider
	teachers TeacherProvider
	groups   GroupProvider
	builder  schedule.Builder
}

func (s ServerHandler) GetGroups(c *gin.Context, params GetGroupsParams) {
	defaultLimit := 40
	if params.Limit == nil {
		params.Limit = &defaultLimit
	}

	result, err := s.groups.Search(c, params.Filter, *params.Limit)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s ServerHandler) GetTeachers(c *gin.Context, params GetTeachersParams) {
	defaultLimit := 40
	if params.Limit == nil {
		params.Limit = &defaultLimit
	}

	result, err := s.teachers.Search(c, params.Filter, *params.Limit)
	if err != nil {
		errors.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
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
