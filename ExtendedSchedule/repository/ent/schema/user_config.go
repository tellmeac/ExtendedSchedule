package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/values"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.JSON("JoinedGroups", values.JoinedGroups{}),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ExcludedLessons", ExcludedLesson.Type),
		edge.To("ExtendedLessons", ExtendedLesson.Type),
	}
}
