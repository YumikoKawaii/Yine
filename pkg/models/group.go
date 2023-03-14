package models

import "github.com/YumikoKawaii/Yine/pkg/config"

type Group struct {
	GID    string `json:"g_id" gorm:"primarykey"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Group{})
}

func (g Group) NewGroup(gid string, name string) {

	newRecord := Group{
		GID:    gid,
		Avatar: "default",
		Name:   name,
	}

	db.Create(&newRecord)

}

func (g Group) ChangeAvatar(gid string, avatar string) {

	db.Exec("update `groups` set avatar = ? where g_id = ?", avatar, gid)

}

func (g Group) ChangeName(gid string, name string) {

	db.Exec("update `groups` set name = ? where g_id = ?", name, gid)

}

func (g Group) IsGroup(coid string) bool {

	result := ""
	db.Raw("select g_id from `groups` where g_id = ?", coid).Scan(&result)
	return result != ""

}

func (g Group) GetGroup(gid string) Group {

	result := Group{}
	db.Raw("select * from `groups` where g_id = ?", gid).Scan(&result)
	return result
}
