package model

import "time"

// KinmhyoEmpMastar 社員マスタテーブルマッピングオブジェクト
type KinmhyoEmpMastar struct {
	empID     int64      `gorm:"column:emp_id"`
	empName   string     `gorm:"column:emp_name"`
	lineID    string     `gorm:"column:line_id"`
	createdAt *time.Time `gorm:"column:created_at"`
	createdID int64      `gorm:"column:created_id"`
	updatedAt *time.Time `gorm:"column:updated_at"`
	updatedID int64      `gorm:"column:updated_id"`
	deletedAt *time.Time `gorm:"column:deleted_at"`
}

// TableName 社員マスタテーブルの名将返却
func (k *KinmhyoEmpMastar) TableName() string {
	return "emp_mastar"
}
