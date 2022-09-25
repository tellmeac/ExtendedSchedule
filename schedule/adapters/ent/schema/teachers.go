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

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.Other("searchVector", tsvector.TSVector{}).SchemaType(map[string]string{
			dialect.Postgres: "tsvector",
		}),
	}
}

func (Teacher) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.TeacherFunc(func(ctx context.Context, m *gen.TeacherMutation) (gen.Value, error) {
				name, ok := m.Name()
				if !ok {
					return next.Mutate(ctx, m)
				}

				m.SetSearchVector(tsvector.ToTSVector(fmt.Sprintf("%s", name)))
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (Teacher) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{}
}
