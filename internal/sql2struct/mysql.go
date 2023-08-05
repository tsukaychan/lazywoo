/**
 * @author tsukiyo
 * @date 2023-08-05 10:17
 */

package sql2struct

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	var tableColumns []*TableColumn
	var err error
	if err = m.DBEngine.Where("TABLE_SCHEMA = ? AND TABLE_NAME = ?", dbName, tableName).Find(&tableColumns).Error; err != nil {
		return nil, err
	}
	return tableColumns, err
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}
