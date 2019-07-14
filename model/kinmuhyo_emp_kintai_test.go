package model

import "testing"

func TestTableNameKintai(t *testing.T) {
	expectedTableName := "emp_kintai"
	actual := KinmhyoEmpKintai{}
	if actual.TableName() != expectedTableName {
		t.Errorf("Test faild. Table name must be %s", expectedTableName)
	}
}
