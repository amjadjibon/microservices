package repo

import (
	"context"

	"github.com/amjadjibon/microservices/pkg/db"
	"github.com/amjadjibon/microservices/title/model"
)

type TitleRepo struct {
	db *db.Postgres
}

func NewTitleRepo(db *db.Postgres) *TitleRepo {
	return &TitleRepo{
		db: db,
	}
}

func (tr *TitleRepo) CreateTitle(ctx context.Context, title *model.Title) (*model.Title, error) {
	query, args, err := tr.db.Builder.Insert(
		"title_title",
	).Columns(
		"name",
		"description",
		"year",
		"created_at",
		"updated_at",
	).Values(
		title.TitleName,
		title.Description,
		title.Year,
		title.CreatedAt,
		title.UpdateAt,
	).Suffix("RETURNING id").ToSql()

	if err != nil {
		return nil, err
	}

	var titleID int
	err = tr.db.Pool.QueryRow(ctx, query, args...).Scan(&titleID)
	if err != nil {
		return nil, err
	}

	title.ID = titleID
	return title, nil
}

func (tr *TitleRepo) GetTitleById(ctx context.Context, titleID int) (*model.Title, error) {
	query, args, err := tr.db.Builder.Select(
		"title_id",
		"title_name",
		"description",
		"year",
		"created_at",
		"updated_at",
	).From(
		"title_title",
	).Where(
		"title_id = ?",
	).ToSql()

	if err != nil {
		return nil, err
	}

	var title model.Title
	err = tr.db.Pool.QueryRow(ctx, query, args...).Scan(
		&title.ID,
		&title.TitleName,
		&title.Description,
		&title.Year,
		&title.CreatedAt,
		&title.UpdateAt,
	)
	if err != nil {
		return nil, err
	}

	return &title, nil
}

func (tr *TitleRepo) GetAllTitle(ctx context.Context) ([]*model.Title, error) {
	query, _, err := tr.db.Builder.Select(
		"title_id",
		"title_name",
		"description",
		"year",
		"created_at",
		"updated_at",
	).From(
		"title_title",
	).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tr.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var titles []*model.Title
	for rows.Next() {
		var title model.Title
		err = rows.Scan(
			&title.ID,
			&title.TitleName,
			&title.Description,
			&title.Year,
			&title.CreatedAt,
			&title.UpdateAt,
		)
		if err != nil {
			return nil, err
		}

		titles = append(titles, &title)
	}

	return titles, nil
}

func (tr *TitleRepo) DeleteTitle(ctx context.Context, titleID int) error {
	query, args, err := tr.db.Builder.Delete(
		"title_title",
	).Where(
		"title_id = ?",
	).ToSql()

	if err != nil {
		return err
	}

	_, err = tr.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
