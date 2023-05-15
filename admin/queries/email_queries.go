package queries

import (
	"admin-server/admin/models"

	"github.com/jmoiron/sqlx"
)

type EmailQueries struct {
	*sqlx.DB
}

func (q *EmailQueries) GetEmails() ([]models.Email, error) {
	emails := []models.Email{}
	query := `SELECT email_valid_id, email_domain, company_id FROM email_valid`

	err := q.Select(&emails, query)
	if err != nil {
		return emails, err
	}

	return emails, nil
}

func (q *EmailQueries) GetEmail(e *models.EmailValidIdDto) (models.Email, error) {
	email := models.Email{}
	query := `SELECT email_valid.email_valid_id, email_valid.email_domain, company.company_id, company.company_name
				FROM email_valid
				JOIN company ON email_valid.company_id = company.company_id
				WHERE email_valid.email_valid_id = $1`

	err := q.Get(&email, query, e.EmailValidId)
	if err != nil {
		return email, err
	}

	return email, nil
}

func (q *EmailQueries) GetEmailByCompanyName(e *models.CompanyNameDto) ([]models.Email, error) {
	email := []models.Email{}
	query := `SELECT email_valid.email_valid_id, email_valid.email_domain, company.company_id, company.company_name
				FROM email_valid
				JOIN company ON email_valid.company_id = company.company_id
				WHERE email_valid.company_name = $1`

	err := q.Select(&email, query, e.CompanyName)
	if err != nil {
		return email, err
	}

	return email, nil
}

func (q *EmailQueries) GetEmailsByCompnayId(e *models.CompanyIdDto) ([]models.Email, error) {
	email := []models.Email{}
	query := `SELECT email_valid.email_valid_id, email_valid.email_domain, company.company_id, company.company_name
				FROM email_valid
				JOIN company ON email_valid.company_id = company.company_id
				WHERE email_valid.company_id = $1`

	err := q.Select(&email, query, e.CompanyId)
	if err != nil {
		return email, err
	}

	return email, nil
}

func (q *EmailQueries) GetCompanyIdByCompanyName(e *models.CompanyNameDto) (models.CompanyIdDto, error) {
	companyId := models.CompanyIdDto{}
	query := `SELECT company.company_id
				FROM email_valid
				JOIN company ON email_valid.company_id = company.company_id
				WHERE email_valid.company_id = $1`

	err := q.Get(&companyId, query, e.CompanyName)
	if err != nil {
		return companyId, err
	}

	return companyId, nil
}

func (q *EmailQueries) CreateEmail(e *models.EmailUploadEntity) error {
	query := `INSERT INTO email_valid (email_domain, company_id) VALUES ($1, $2)`

	_, err := q.Exec(query, e.EmailDomain, e.CompanyId)
	if err != nil {
		return err
	}

	return nil
}

func (q *EmailQueries) UpdateEmail(e *models.EmailUpdateEntity) error {
	query := `UPDATE email_valid SET email_domain = $1 WHERE company_id = $2 AND email_valid_id = $3`

	_, err := q.Exec(query, e.EmailDomain, e.CompanyId, e.EmailValidId)
	if err != nil {
		return err
	}

	return nil
}

func (q *EmailQueries) DeleteEmail(e *models.EmailValidIdDto) error {
	query := `DELETE FROM email_valid WHERE email_valid_id = $1`

	_, err := q.Exec(query, e.EmailValidId)
	if err != nil {
		return err
	}

	return nil
}
