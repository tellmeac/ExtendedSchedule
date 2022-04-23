package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/entity"
)

// ExtendedLesson holds the schema definition for the ExtendedLesson entity.
type ExtendedLesson struct {
	ent.Schema
}

// Fields of the ExtendedLesson.
func (ExtendedLesson) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("UserID", uuid.UUID{}),
		field.String("Description"),
		field.JSON("LessonRef", entity.LessonRef{}),
	}
}

// Edges of the ExtendedLesson.
func (ExtendedLesson) Edges() []ent.Edge {
	return nil
}
