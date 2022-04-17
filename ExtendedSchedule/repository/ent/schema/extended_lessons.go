package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/values"
)

// ExtendedLesson holds the schema definition for the ExtendedLesson entity.
type ExtendedLesson struct {
	ent.Schema
}

// Fields of the ExtendedLesson.
func (ExtendedLesson) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("ConfigID", uuid.UUID{}),
		field.JSON("Context", values.LessonContext{}),
		field.Bool("IsPrivate").Default(true),
		field.JSON("Intervals", values.LessonInterval{}),
	}
}

// Edges of the ExtendedLesson.
func (ExtendedLesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserConfig", UserConfig.Type).Ref("ExtendedLessons").Unique().Field("ConfigID").Required(),
	}
}
