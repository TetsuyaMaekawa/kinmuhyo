package model

import "testing"

func TestTableNameMastar(t *testing.T) {
	expectedTableName := "emp_mastar"
	actual := KinmhyoEmpMastar{}
	if actual.TableName() != expectedTableName {
		t.Errorf("Test faild. Table name must be %s", expectedTableName)
	}
}
