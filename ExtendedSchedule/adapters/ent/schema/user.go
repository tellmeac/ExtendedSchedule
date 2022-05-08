package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserInfo holds the schema definition for the UserInfo entity.
type UserInfo struct {
	ent.Schema
}

// Fields of the UserInfo.
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("Email").NotEmpty(),
	}
}

func (UserInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Email").Unique(),
	}
}

// Edges of the UserInfo.
func (UserInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("JoinedGroups", JoinedGroups.Type),
		edge.To("ExcludedLessons", ExcludedLesson.Type),
	}
}
