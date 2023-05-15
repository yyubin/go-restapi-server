package models

import "github.com/google/uuid"

type Email struct {
	EmailValidId int       `db:"email_valid_id" json:"email_valid_id"`
	EmailDomain  string    `db:"email_domain" json:"email_domain"`
	CompanyId    uuid.UUID `db:"company_id" json:"company_id"`
	CompanyName  string    `db:"company_name" json:"company_name"`
}

type EmailValidIdDto struct {
	EmailValidId int `db:"email_valid_id" json:"email_valid_id"`
}

func (e *EmailValidIdDto) SetEmailValidDto(id int) {
	e.EmailValidId = id
}

type EmailUploadDto struct {
	EmailDomain string    `json:"email_domain"`
	CompanyId   uuid.UUID `json:"company_id"`
}

type EmailUploadEntity struct {
	EmailDomain string    `db:"email_domain"`
	CompanyId   uuid.UUID `db:"company_id"`
}

type EmailUpdateDto struct {
	EmailValidId int       `json:"email_valid_id"`
	EmailDomain  string    `json:"email_domain"`
	CompanyId    uuid.UUID `json:"company_id"`
}

type EmailUpdateEntity struct {
	EmailValidId int       `db:"email_valid_id"`
	EmailDomain  string    `db:"email_domain"`
	CompanyId    uuid.UUID `db:"company_id"`
}

func (e *EmailUploadDto) SetEmailUploadDtoToEntity() *EmailUploadEntity {
	return &EmailUploadEntity{
		EmailDomain: e.EmailDomain,
		CompanyId:   e.CompanyId,
	}
}

func (e *EmailUpdateDto) SetEmailUpdateDtoToEntity() *EmailUpdateEntity {
	return &EmailUpdateEntity{
		EmailValidId: e.EmailValidId,
		EmailDomain:  e.EmailDomain,
		CompanyId:    e.CompanyId,
	}
}

func (e *Email) SetEmailToEntity() *EmailUploadEntity {
	return &EmailUploadEntity{
		EmailDomain: e.EmailDomain,
		CompanyId:   e.CompanyId,
	}
}

type Company struct {
	CompanyId      uuid.UUID `db:"company_id" json:"company_id"`
	CompanyName    string    `db:"company_name" json:"company_name"`
	IndustryId     int       `db:"industry_id" json:"industry_id"`
	CompanyImageId int       `db:"company_image_id" json:"company_image_id"`
}

type CompanyIdDto struct {
	CompanyId uuid.UUID `db:"company_id" json:"company_id"`
}

func (c *CompanyIdDto) SetCompanyIdDto(id uuid.UUID) {
	c.CompanyId = id
}

type CompanyNameDto struct {
	CompanyName string `db:"company_name" json:"company_name"`
}

func (c *CompanyNameDto) SetCompanyName(name string) {
	c.CompanyName = name
}
