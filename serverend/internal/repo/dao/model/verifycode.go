package model

type VerifyCode struct {
	Id        int64  `gorm:"column:id;primary_key;auto_increment"`
	Biz       string `gorm:"column:biz"`
	Target    string `gorm:"column:target"`
	Code      string `gorm:"column:code"`
	ExpiresAt int64  `gorm:"column:expires_at"`
	IsUsed    int    `gorm:"column:is_used"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (VerifyCode) TableName() string {
	return "verify_code"
}
