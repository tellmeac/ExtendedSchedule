package schema

import (
	"tellmeac/extended-schedule/domain/entity"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("UserID", uuid.UUID{}),
		field.JSON("JoinedGroups", []entity.GroupInfo{}),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return nil
}
