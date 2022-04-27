package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("UserID", uuid.UUID{}),
		field.String("LessonID"),
		field.Int("Position"),
		field.Int("Weekday"),
		field.JSON("Teacher", &entity.TeacherInfo{}),
	}
}

func (ExcludedLesson) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("UserID"),
	}
}

// Edges of the ExcludedLesson.
func (ExcludedLesson) Edges() []ent.Edge {
	return nil
}
