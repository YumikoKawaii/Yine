package models

import (
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

type Conversation struct {
	User     string `json:"user"`
	CoID     string `json:"co_id"`
	Role     string `json:"role"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

func init() {
	db.AutoMigrate(&Conversation{})
}

func (c Conversation) NewConnect(coid string, user string, role string) {

	db.Create(&Conversation{
		CoID:     coid,
		User:     user,
		Role:     role,
		Nickname: "",
	})

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

func (c Conversation) IsMember(coid string, user string) bool {

	result := ""
	db.Raw("select role from conversations where co_id = ? and user = ?", coid, user).Scan(&result)
	return result != ""

}

func (c Conversation) IsAdmin(coid string, user string) bool {

	result := ""
	db.Raw("select role from conversations where co_id = ? and user = ?", coid, user).Scan(&result)
	return result == utils.Admin
}

func (c Conversation) RemoveConversation(coid string) {

	db.Exec("delete from conversations where co_id = ?", coid)

}

func (c Conversation) IsConversationExist(coid string) bool {

	result := ""
	db.Raw("select distinct co_id from conversations where co_id = ?", coid).Scan(&result)
	return result != ""
}

func (c Conversation) IsConversationBetween(user string, guest string) bool {

	result := ""
	db.Raw("select user from conversations where user = ? and co_id = ?", user, guest)

	return result != ""
}

func (c Conversation) GetReceivers(coid string, id string) []string {

	result := make([]string, 0)

	db.Raw("select user from conversations where co_id = ? and user <> ?", coid, id).Scan(&result)

	return result

}
