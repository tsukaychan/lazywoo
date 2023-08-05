/**
 * @author tsukiyo
 * @date 2023-08-06 1:36
 */

package cmd

import (
	"github.com/spf13/cobra"
	"lazywoo/internal/sql2struct"
	"log"
)

var (
	username, password, host, charset, dbType, dbName, tableName string
)

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "please input db username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "please input db password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "please input db host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "please input db charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "please input db type")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "please input db name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "please input table name")
}

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql conversion and processing",
	Long:  "sql conversion and processing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql conversion",
	Long:  "sql conversion",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		if err := dbModel.Connect(); err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}
