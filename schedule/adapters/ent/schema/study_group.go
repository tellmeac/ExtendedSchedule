package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"fmt"
	gen "tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/adapters/ent/hook"
	"tellmeac/extended-schedule/common/tsvector"
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
		field.Other("searchVector", tsvector.TSVector{}).SchemaType(map[string]string{
			dialect.Postgres: "tsvector",
		}),
	}
}

func (StudyGroup) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.StudyGroupFunc(func(ctx context.Context, m *gen.StudyGroupMutation) (gen.Value, error) {
				name, ok := m.Name()
				if !ok {
					return next.Mutate(ctx, m)
				}

				facultyName, ok := m.FacultyName()
				if !ok {
					return next.Mutate(ctx, m)
				}

				m.SetSearchVector(tsvector.ToTSVector("russian", fmt.Sprintf("%s %s", name, facultyName)))
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (StudyGroup) Indexes() []ent.Index {
	return []ent.Index{}
}

func (StudyGroup) Edges() []ent.Edge {
	return []ent.Edge{}
}
