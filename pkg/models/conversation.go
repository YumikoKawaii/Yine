package models

import "github.com/YumikoKawaii/Yine/pkg/config"

type Conversation struct {
	CoID     string `json:"co_id" gorm:"primarykey"`
	User     string `json:"user"`
	Role     string `json:"role"`
	Nickname string `json:"nickname"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Conversation{})
}

func (c Conversation) NewConnect(coid string, user string, role string) {

	newRecord := Conversation{
		CoID:     coid,
		User:     user,
		Role:     role,
		Nickname: "",
	}

	db.Create(&newRecord)

}

func (c Conversation) ChangeRole(coid string, user string, role string) {

	db.Exec("update conversations set role = ? where co_id = ? and user = ?", role, coid, user)

}

func (c Conversation) RemoveUser(coid string, user string) {

	db.Exec("delete from conversations where co_id = ? and user = ?", coid, user)

}

func (c Conversation) RemoveUserData(user string) {

	db.Exec("delete from conversation where user = ?", user)

}

func (c Conversation) ChangeNickname(coid string, user string, nickname string) {

	db.Exec("update conversations set nickname = ? where co_id = ? and user = ?", nickname, coid, user)

}

func (c Conversation) GetRole(coid string, user string) string {

	result := ""
	db.Raw("select role from conversations where co_id = ? and user = ?", coid, user).Scan(&result)
	return result

}

func (c Conversation) RemoveConversation(coid string) {

	db.Exec("delete from conversations where co_id = ?", coid)

}
