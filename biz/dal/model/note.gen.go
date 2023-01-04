// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNote = "note"

// Note mapped from table <note>
type Note struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UserID    int64          `gorm:"column:user_id" json:"user_id"`
	Title     string         `gorm:"column:title" json:"title"`
	Content   string         `gorm:"column:content" json:"content"`
}

// TableName Note's table name
func (*Note) TableName() string {
	return TableNameNote
}