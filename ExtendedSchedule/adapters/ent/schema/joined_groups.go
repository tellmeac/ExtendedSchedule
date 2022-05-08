package schema

import (
	"entgo.io/ent/schema/edge"
	"tellmeac/extended-schedule/domain/entity"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// JoinedGroups holds the schema definition for the JoinedGroups entity.
type JoinedGroups struct {
	ent.Schema
}

// Fields of the JoinedGroups.
func (JoinedGroups) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.UUID("UserID", uuid.UUID{}),
		field.JSON("JoinedGroups", []entity.GroupInfo{}),
	}
}

// Edges of the JoinedGroups.
func (JoinedGroups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserInfo", UserInfo.Type).Ref("JoinedGroups").Unique().Field("UserID").Required(),
	}
}
