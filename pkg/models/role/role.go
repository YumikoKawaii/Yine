package role

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Role struct {
	User     string `json:"user"`
	ConvID   string `json:"conv_id"`
	Role     string `json:"role"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (r Role) NewConnect(conv_id string, user string, role string) {

	db.Create(Role{
		User:     user,
		ConvID:   conv_id,
		Role:     role,
		Nickname: "",
		Status:   "",
	})

}

func (r Role) ChangeRole(conv_id string, user string, role string) {
	db.Exec("update roles set role = ? where conv_id = ? and user = ?", role, conv_id, user)
}

func (r Role) RemoveUser(conv_id string, user string) {
	db.Exec("delete from roles where conv_id = ? and user = ?", conv_id, user)
}

func (r Role) RemoveUserData(user string) {

	db.Exec("delete from conversation where user = ?", user)

}

func (r Role) ChangeNickname(conv_id string, user string, nickname string) {

	db.Exec("update roles set nickname = ? where conv_id = ? and user = ?", nickname, conv_id, user)

}

func (r Role) IsMember(conv_id string, user string) bool {

	result := ""
	db.Raw("select role from roles where conv_id = ? and user = ?", conv_id, user).Scan(&result)
	return result != ""

}

func (r Role) IsAdmin(conv_id string, user string) bool {

	result := ""
	db.Raw("select role from roles where conv_id = ? and user = ?", conv_id, user).Scan(&result)
	return result == utils.Admin
}

func (r Role) RemoveConversation(conv_id string) {

	db.Exec("delete from roles where conv_id = ?", conv_id)

}

func (r Role) IsConversationExist(conv_id string) bool {

	result := ""
	db.Raw("select distinct conv_id from roles where conv_id = ?", conv_id).Scan(&result)
	return result != ""
}

func (r Role) IsConversationBetween(user string, guest string) bool {

	result := ""
	db.Raw("select user from roles where user = ? and conv_id = ?", user, guest).Scan(&result)
	return result != ""
}

func (r Role) GetReceivers(conv_id string, id string) []string {

	result := make([]string, 0)
	db.Raw("select user from roles where conv_id = ? and user <> ?", conv_id, id).Scan(&result)
	return result

}

func (r Role) GetAllConversation(id string) []string {

	result := make([]string, 0)
	db.Raw("select conv_id from roles where user = ?", id).Scan(&result)
	return result
}

func (r Role) GetGroupMember(gid string) []Role {

	result := make([]Role, 0)
	db.Raw("select * from roles where conv_id = ?", gid).Scan(&result)
	return result

}

func (r Role) GetPartner(user string, conv_id string) Role {

	result := Role{}
	db.Raw("select * from roles where conv_id = ? and user = ?", user, conv_id).Scan(&result)
	return result

}
