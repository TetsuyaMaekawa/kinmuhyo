package rds

import "rootship.co.jp/kinmuhyo/model"

// InsertEmpMastar マスタ情報の登録
func (c Client) InsertEmpMastar(command []string, lineID string) error {
	empMastar := model.KinmuhyoEmpMastar{}
	empMastar.EmpID = command[1]
	empMastar.EmpName = command[2]
	empMastar.LineID = "hoge"

	c.DB.Create(&empMastar)

	return nil
}
