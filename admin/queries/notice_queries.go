package queries

import (
	"admin-server/admin/models"

	"github.com/jmoiron/sqlx"
)

type NoticeQueries struct {
	*sqlx.DB
}

func (q *NoticeQueries) GetNotices() ([]models.Notice, error) {
	notices := []models.Notice{}
	query := `SELECT notice_id, notice_title, notice_contents, profile_id, created_at, updated_at FROM notice WHERE use_yn = true`

	err := q.Select(&notices, query)
	if err != nil {
		return notices, err
	}

	return notices, nil
}

func (q *NoticeQueries) GetNotice(n *models.NoticeIdDto) (models.Notice, error) {
	notice := models.Notice{}

	query := `SELECT notice_id, notice_title, notice_contents, profile_id, created_at, updated_at FROM notice WHERE notice_id = $1 and use_yn = true`

	err := q.Get(&notice, query, n.NoticeId)
	if err != nil {
		return notice, err
	}

	return notice, nil
}

func (q *NoticeQueries) CreateNotice(n *models.NoticeUploadEntity) error {
	query := `INSERT INTO notice (notice_title, notice_contents, profile_id) VALUES ($1, $2, $3);`

	_, err := q.Exec(query, n.NoticeTitle, n.NoticeContents, n.ProfileId)
	if err != nil {
		return err
	}

	return nil
}

func (q *NoticeQueries) UpdatedNotice(nid *models.NoticeIdEntity, n *models.NoticeUploadEntity) error {
	query := `UPDATE notice SET notice_title = $1, notice_contents = $2 WHERE profile_id = $3 and notice_id = $4`

	_, err := q.Exec(query, n.NoticeTitle, n.NoticeContents, n.ProfileId, nid.NoticeId)
	if err != nil {
		return err
	}

	return nil
}

func (q *NoticeQueries) DeleteNotice(n *models.NoticeDeleteEntity) error {
	query := `DELETE FROM notice WHERE notice_id = $1 AND profile_id = $2`

	_, err := q.Exec(query, n.NoticeId, n.ProfileId)
	if err != nil {
		return nil
	}

	return nil
}
