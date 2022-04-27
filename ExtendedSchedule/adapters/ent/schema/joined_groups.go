package schema

import (
	"entgo.io/ent/schema/index"
	"tellmeac/extended-schedule/domain/entity"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// JoinedGroups holds the schema definition for the JoinedGroups entity.
type JoinedGroups struct {
	ent.Schema
}

// Fields of the UserConfig.
func (JoinedGroups) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("UserID", uuid.UUID{}),
		field.JSON("JoinedGroups", []entity.GroupInfo{}),
	}
}

func (JoinedGroups) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("UserID").Unique(),
	}
}

// Edges of the UserConfig.
func (JoinedGroups) Edges() []ent.Edge {
	return nil
}
