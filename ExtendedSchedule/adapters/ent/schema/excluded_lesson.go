package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.String("LessonRef"),
	}
}

// Edges of the ExcludedLesson.
func (ExcludedLesson) Edges() []ent.Edge {
	return nil
}
