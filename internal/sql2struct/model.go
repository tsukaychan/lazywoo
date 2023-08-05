/**
 * @author tsukiyo
 * @date 2023-08-05 11:47
 */

package sql2struct

import "gorm.io/gorm"

type DBModel struct {
	DBEngine *gorm.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

func (t TableColumn) TableName() string {
	return "COLUMNS"
}
