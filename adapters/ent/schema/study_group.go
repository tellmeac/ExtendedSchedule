package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StudyGroup holds the schema definition for the StudyGroup entity.
type StudyGroup struct {
	ent.Schema
}

func (StudyGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("facultyName"),
	}
}
