/**
 * @author tsukiyo
 * @date 2023-08-05 11:56
 */

package sql2struct

import (
	"testing"
)

func TestGorm(t *testing.T) {
	db := NewDBModel(&DBInfo{
		UserName: "root",
		Password: "for.nothing",
		Host:     "124.70.190.134",
		Charset:  "utf8",
	})
	db.Connect()
	columns, err := db.GetColumns("dev", "blog_tab")
	if err != nil {
		t.Error(err)
	}
	for _, c := range columns {
		t.Logf("%v", *c)
	}
}
