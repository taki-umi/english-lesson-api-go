package model

import "time"

// BaseEntity - ベースエンティティを定義します。
type BaseEntity struct {
	createdAt time.Time
	updatedAt time.Time
	createUserID string
	updateUserID string
}

// NewBaseEntity - ベースエンティティを生成します。
func NewBaseEntity(createUserID string, updateUserID string) BaseEntity {
	t := time.Now()

	return BaseEntity{
		createdAt: t,
		updatedAt: t,
		createUserID: createUserID,
		updateUserID: updateUserID,
	}
}

// NewBaseEntityForRebuild - ベースエンティティを再構築します。
func NewBaseEntityForRebuild(createUserID string, updateUserID string, createdAt time.Time, updatedAt time.Time) BaseEntity {
	return BaseEntity{
		createdAt: createdAt,
		updatedAt: updatedAt,
		createUserID: createUserID,
		updateUserID: updateUserID,
	}
}

// CreatedAt - 作成日時を取得します。
func (base BaseEntity) CreatedAt() time.Time {
	return base.createdAt
}

// UpdatedAt - 更新日時を取得します。
func (base BaseEntity) UpdatedAt() time.Time {
	return base.updatedAt
}

// CreateUserID - 作成ユーザーIDを取得します。
func (base BaseEntity) CreateUserID() string {
	return base.createUserID
}

// UpdateUserID - 更新ユーザーIDを取得します。
func (base BaseEntity) UpdateUserID() string {
	return base.updateUserID
}