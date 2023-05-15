package models

import (
	"time"

	"github.com/google/uuid"
)

// @Description 공지사항 모델
type Notice struct {
	NoticeId       int       `db:"notice_id" json:"notice_id"`
	NoticeTitle    string    `db:"notice_title" json:"notice_title" example:"notice_title"`
	NoticeContents string    `db:"notice_contents" json:"notice_contents" example:"notice_contents"`
	ProfileId      uuid.UUID `db:"profile_id" json:"profile_id" example:"profile_id" validate:"required,uuid"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

// 공지사항 업데이트시 엔티티 Notice -> NoticeUploadEntity
func (n *Notice) SetUploadNoticeToEntity() *NoticeUploadEntity {
	return &NoticeUploadEntity{
		NoticeTitle:    n.NoticeTitle,
		NoticeContents: n.NoticeContents,
		ProfileId:      n.ProfileId,
	}
}

// 공지사항 삭제시 Notice -> NoticeDeleteEntity
func (n *Notice) SetDeleteNoticeToEntity() *NoticeDeleteEntity {
	return &NoticeDeleteEntity{
		NoticeId:  n.NoticeId,
		ProfileId: n.ProfileId,
	}
}

func (n *Notice) SetNoticeIdDto() *NoticeIdDto {
	return &NoticeIdDto{
		NoticeId: n.NoticeId,
	}
}

// @Description 공지사항 조회시 아이디
type NoticeIdDto struct {
	NoticeId int `json:"notice_id"`
}

// @Description NoticeIdEntity
type NoticeIdEntity struct {
	NoticeId int `db:"notice_id"`
}

// 공지사항 아이디 값, NoticeIdDto Setter
func (n *NoticeIdDto) SetNoticeId(id int) {
	n.NoticeId = id
}

// 공지사항 조회 및 등록시 NoticeIdDto -> NoticeIdEntity
// 조회시엔 dto entity 구분 없어도 될듯
func (n *NoticeIdDto) SetNoticeIdToEntity() *NoticeIdEntity {
	return &NoticeIdEntity{
		NoticeId: n.NoticeId,
	}
}

// @Description 공지사항 삭제
// 로직 수정 필요함 나중엔 id값만 받고 jwt에서 profile_id 가져와서 검사해야함
type NoticeDeleteDto struct {
	NoticeId  int       `json:"notice_id" example:"notice_id"`
	ProfileId uuid.UUID `db:"profile_id"`
}

func (n *NoticeDeleteDto) SetNoticeIdDto() *NoticeIdDto {
	return &NoticeIdDto{
		NoticeId: n.NoticeId,
	}
}

// Notice삭제시 사용하는 엔티티
type NoticeDeleteEntity struct {
	NoticeId  int       `json:"notice_id" example:"notice_id"`
	ProfileId uuid.UUID `db:"profile_id"`
}

// @Description 공지사항 등록 폼
type NoticeUploadDto struct {
	NoticeTitle    string    `db:"notice_title" json:"notice_title" example:"notice_title"`
	NoticeContents string    `db:"notice_contents" json:"notice_contents" example:"notice_contents"`
	ProfileId      uuid.UUID `db:"profile_id" json:"profile_id" example:"profile_id" validate:"required,uuid"`
}

// Notice upload entity
type NoticeUploadEntity struct {
	NoticeTitle    string    `db:"notice_title" json:"notice_title"`
	NoticeContents string    `db:"notice_contents" json:"notice_contents"`
	ProfileId      uuid.UUID `db:"profile_id" json:"profile_id"`
}

func (n *NoticeUploadDto) SetNoticeDtoToEntity() *NoticeUploadEntity {
	return &NoticeUploadEntity{
		NoticeTitle:    n.NoticeTitle,
		NoticeContents: n.NoticeContents,
		ProfileId:      n.ProfileId,
	}
}

func (n *NoticeUploadEntity) SetNoticeEntityToDto() *NoticeUploadDto {
	return &NoticeUploadDto{
		NoticeTitle:    n.NoticeTitle,
		NoticeContents: n.NoticeContents,
		ProfileId:      n.ProfileId,
	}
}
