package app

import (
	"github.com/cjwddz/fast-model"
	"fmt"
	m "../model"
)
var AppModel model.DbModel
var CodeModel model.DbModel
var UserModel model.DbModel
var FileModel model.DbModel
var VlogModel model.DbModel
var LogModel model.DbModel

func Init(){
	app,err:= m.GetAppModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	code,err:=m.GetCodeModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	user,err:=m.GetUserModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	file,err:=m.GetFileModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	vlog,err:=m.GetVLogModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	log,err:=m.GetLogModel()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	AppModel = app
	CodeModel = code
	UserModel = user
	FileModel = file
	VlogModel = vlog
	LogModel = log
}
