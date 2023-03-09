package models

import "github.com/YumikoKawaii/Yine/pkg/config"

type Role struct {
	User         string `json:"user" gorm:"primarykey"`
	Conversation string `json:"conversation"`
	Role         string `json:"role"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Role{})
}

func (r Role) NewRole(user string, conversation string, role string) {

	newRecord := Role{
		User:         user,
		Conversation: conversation,
		Role:         role,
	}

	db.Create(&newRecord)

}

func (r Role) ChangeRole(user string, conversation string, role string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("update roles set role = ? where id = ? and conversation = ?", role, user, conversation)
	db.Exec("set sql_safe_updates = 1")

}

func (r Role) GetRole(user string, conversation string) string {

	result := ""
	db.Raw("select role from roles where user = ? and conversation = ?", user, conversation)
	return result

}

func (r Role) RemoveUser(user string, conversation string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("delete from roles where user = ? and conversation = ?", user, conversation)
	db.Exec("set sql_safe_updates = 1")

}

func (r Role) RemoveAllUserData(user string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("delete from roles where user = ?", user)
	db.Exec("set sql_safe_updates = 1")

}

func (r Role) RemoveConversation(conversation string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("delete from roles where conversation = ?", conversation)
	db.Exec("set sql_safe_updates = 1")
}
