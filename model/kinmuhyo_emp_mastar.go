package model

import "time"

// KinmuhyoEmpMastar 社員マスタテーブルマッピングオブジェクト
type KinmuhyoEmpMastar struct {
	EmpID     string     `gorm:"column:emp_id"`
	EmpName   string     `gorm:"column:emp_name"`
	LineID    string     `gorm:"column:line_id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	CreatedID int64      `gorm:"column:created_id"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedID int64      `gorm:"column:updated_id"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// TableName 社員マスタテーブルの名将返却
func (k *KinmuhyoEmpMastar) TableName() string {
	return "emp_mastar"
}
