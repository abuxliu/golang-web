package model

import (
	"time"
	m "github.com/cjwddz/fast-model"
)

/*
DROP TABLE IF EXISTS t_file;
CREATE TABLE t_file (
"id" serial NOT NULL,
"key" varchar(16) UNIQUE COLLATE "default",
"type" varchar(16) COLLATE "default",
"name" varchar(128) COLLATE "default",
"owner" varchar(128) COLLATE "default",
"created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
WITH (OIDS=FALSE);
 */
type T_File struct {
	ID        int64     `json:"id"`
	Key       string    `json:"key"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
}

func GetFileModel() (m.DbModel, error){
	sc:=m.SqlController {
		TableName:      "t_file",
		InsertColumns:  []string{"key","type","name","owner","created_at"},
		QueryColumns:   []string{"id","key","type","name","owner","created_at"},
		InSertFields:   insertFileFields,
		QueryField2Obj: queryFileField2Obj,
	}
	return m.GetModel(sc)
}

func insertFileFields(obj interface{}) []interface{} {
	file:=obj.(T_File)
	return []interface{}{
		file.Key, file.Type, file.Name, file.Owner, file.CreatedAt,
	}
}
func queryFileField2Obj(fields []interface{}) interface{} {
	file:=T_File{
		ID:        m.GetInt64(fields[0],0),
		Key:       m.GetString(fields[1]),
		Type:      m.GetString(fields[2]),
		Name:      m.GetString(fields[3]),
		Owner:     m.GetString(fields[4]),
		CreatedAt: m.GetTime(fields[5],time.Now()),
	}
	return file
}