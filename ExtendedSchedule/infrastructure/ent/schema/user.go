package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/userconfig"
	"tellmeac/extended-schedule/domain/values"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("Email").NotEmpty(),
		field.JSON("BaseGroup", &values.GroupInfo{}),
		field.JSON("ExcludeRules", []userconfig.ExcludeRule{}),
		field.JSON("ExtendedGroupLessons", []userconfig.ExtendedGroupLessons{}),
	}
}

func (UserConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Email").Unique(),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
