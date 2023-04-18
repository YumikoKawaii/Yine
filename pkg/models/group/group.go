package group

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Group struct {
	GID    string `json:"group_id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (g Group) NewGroup(group_id string, name string) {

	db.Create(Group{
		GID:    group_id,
		Avatar: "default",
		Name:   name,
	})

}

func (g Group) ChangeAvatar(group_id string, avatar string) {

	db.Exec("update `groups` set avatar = ? where group_id = ?", avatar, group_id)

}

func (g Group) ChangeName(group_id string, name string) {

	db.Exec("update `groups` set name = ? where group_id = ?", name, group_id)

}

func (g Group) IsGroup(coid string) bool {

	result := ""
	db.Raw("select group_id from `groups` where group_id = ?", coid).Scan(&result)
	return result != ""

}

func (g Group) GetGroup(group_id string) Group {

	result := Group{}
	db.Raw("select * from `groups` where group_id = ?", group_id).Scan(&result)
	return result
}
