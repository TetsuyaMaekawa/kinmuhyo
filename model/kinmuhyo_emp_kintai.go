package model

import "time"

// KinmhyoEmpKintai 社員勤怠テーブルマッピングオブジェクト
type KinmhyoEmpKintai struct {
	empID          int64      `gorm:"column:emp_id"`
	kintai         string     `gorm:"column:kintai"`
	meirei         string     `gorm:"column:mairei"`
	bikou          string     `gorm:"column:bikou"`
	gyomunaiyo     string     `gorm:"column:gyomunaiyo"`
	syukkinntaikin string     `gorm:"column:syukkintaikin"`
	createdAt      *time.Time `gorm:"column:created_at"`
	createdID      int64      `gorm:"column:created_id"`
	updatedAt      *time.Time `gorm:"column:updated_at"`
	updatedID      int64      `gorm:"column:updated_id"`
	deletedAt      *time.Time `gorm:"column:deleted_at"`
}

// TableName 社員勤怠テーブル名称返却
func (k *KinmhyoEmpKintai) TableName() string {
	return "emp_kintai"
}
