package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	commonmodels "tellmeac/extended-schedule/common/models"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("Email").NotEmpty(),
		field.JSON("BaseGroup", &commonmodels.GroupInfo{}),
		field.JSON("ExcludeRules", []commonmodels.ExcludeRule{}),
		field.JSON("ExtendedGroupLessons", []commonmodels.ExtendedGroupLessons{}),
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
