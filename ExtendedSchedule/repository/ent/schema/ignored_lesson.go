package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/values"
)

// IgnoredLesson holds the schema definition for the IgnoredLesson entity.
type IgnoredLesson struct {
	ent.Schema
}

// Fields of the IgnoredLesson.
func (IgnoredLesson) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("ConfigID", uuid.UUID{}),
		field.JSON("Context", values.LessonContext{}),
		field.JSON("Intervals", values.LessonInterval{}),
	}
}

// Edges of the IgnoredLesson.
func (IgnoredLesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserConfig", UserConfig.Type).Ref("IgnoredLessons").Unique().Field("ConfigID").Required(),
	}
}
