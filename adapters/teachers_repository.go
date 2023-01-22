package adapters

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/adapters/ent/teacher"
	"tellmeac/extended-schedule/schedule"
)

func NewTeachersRepository(client *ent.Client) TeacherRepository {
	return TeacherRepository{client: client}
}

type TeacherRepository struct {
	client *ent.Client
}

func (r TeacherRepository) Search(ctx context.Context, filter string, limit int) ([]schedule.Teacher, error) {
	q := fmt.Sprintf("select %s, %s, similarity(%s, $1) as score from %s order by score desc limit %d",
		teacher.FieldID, teacher.FieldName,
		teacher.FieldName,
		teacher.Table, limit,
	)

	rows, err := r.client.QueryContext(ctx, q, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query teachers: %w", err)
	}

	var result []schedule.Teacher
	var score float32
	for rows.Next() {
		g := schedule.Teacher{}
		if err := rows.Scan(&g.ID, &g.Name, &score); err != nil {
			return nil, fmt.Errorf("failed to read raw rows to models: %w", err)
		}

		result = append(result, g)
	}

	return result, nil
}
