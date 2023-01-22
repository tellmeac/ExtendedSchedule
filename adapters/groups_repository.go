package adapters

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	sg "tellmeac/extended-sc/adapters/ent/studygroup"
	"tellmeac/extended-schedule/schedule"
)

func NewStudyGroupRepository(client *ent.Client) StudyGroupRepository {
	return StudyGroupRepository{client: client}
}

type StudyGroupRepository struct {
	client *ent.Client
}

func (r StudyGroupRepository) Search(ctx context.Context, filter string, limit int) ([]schedule.StudyGroup, error) {
	q := fmt.Sprintf("select %s, %s, %s, similarity(%s || ' ' || %s, $1) as score from %s order by score desc limit %d;",
		sg.FieldID, sg.FieldName, sg.FieldFacultyName,
		sg.FieldName, sg.FieldFacultyName,
		sg.Table, limit,
	)

	rows, err := r.client.QueryContext(ctx, q, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query groups: %w", err)
	}

	var result []schedule.StudyGroup
	var score float32
	for rows.Next() {
		g := schedule.StudyGroup{}
		if err := rows.Scan(&g.ID, &g.Name, &g.Faculty, &score); err != nil {
			return nil, fmt.Errorf("failed to read raw rows to models: %w", err)
		}

		result = append(result, g)
	}

	return result, nil
}
