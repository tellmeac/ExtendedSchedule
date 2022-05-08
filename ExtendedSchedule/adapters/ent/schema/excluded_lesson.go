package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/entity"
)

// ExcludedLesson holds the schema definition for the ExcludedLesson entity.
type ExcludedLesson struct {
	ent.Schema
}

// Fields of the ExcludedLesson.
func (ExcludedLesson) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.UUID("UserID", uuid.UUID{}),
		field.String("LessonID"),
		field.Int("Position"),
		field.Int("Weekday"),
		field.JSON("Teacher", &entity.TeacherInfo{}),
	}
}

// Edges of the ExcludedLesson.
func (ExcludedLesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserInfo", UserInfo.Type).Ref("ExcludedLessons").Unique().Field("UserID").Required(),
	}
}
