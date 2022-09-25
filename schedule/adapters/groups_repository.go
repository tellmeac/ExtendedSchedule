package adapters

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	sg "tellmeac/extended-schedule/adapters/ent/studygroup"
	"tellmeac/extended-schedule/schedule"
)

func NewStudyGroupRepository(client *ent.Client) StudyGroupRepository {
	return StudyGroupRepository{client: client}
}

type StudyGroupRepository struct {
	client *ent.Client
}

func (r StudyGroupRepository) SearchGroups(ctx context.Context, search string) ([]schedule.StudyGroup, error) {
	q := fmt.Sprintf("select %s, %s, %s from %s where %s @@ to_tsquery('?');",
		sg.Table, sg.FieldID, sg.FieldName, sg.FieldFacultyName, sg.FieldSearchVector)

	rows, err := r.client.QueryContext(ctx, q, search)
	if err != nil {
		return nil, fmt.Errorf("failed to query groups: %w", err)
	}

	var result []schedule.StudyGroup
	for rows.Next() {
		g := schedule.StudyGroup{}
		if err := rows.Scan(g.ID, g.Name, g.Faculty); err != nil {
			return nil, fmt.Errorf("failed to read raw rows to models: %w", err)
		}

		result = append(result, g)
	}

	return result, nil
}
