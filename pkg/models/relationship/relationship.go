package relationship

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

type Relationship struct {
	User   string `json:"user"`
	Guest  string `json:"guest"`
	Status string `json:"status"`
}

var (
	db *gorm.DB
)

func init() {
	config.Connect()
	db = config.GetDB()
}

func (r Relationship) GetRelationshipBetween(user string, guest string) string {

	result := ""
	db.Raw("select status from relationships where user = ? and guest = ?", user, guest).Scan(&result)
	return result

}

func (r Relationship) SentRequest(user string, guest string) {

	db.Create(Relationship{
		User:   user,
		Guest:  guest,
		Status: utils.SentRequest,
	})

	db.Create(Relationship{
		User:   guest,
		Guest:  user,
		Status: utils.GotRequest,
	})

}

func (r Relationship) AcceptRequest(user string, guest string) {

	db.Exec("update relationships set status = ? where user = ? and guest = ?", utils.Friend, user, guest)
	db.Exec("update relationships set status = ? where user = ? and guest = ?", utils.Friend, guest, user)

}

func (r Relationship) CancelStatus(user string, guest string) {

	db.Exec("delete from relationships where user = ? and guest = ?", user, guest)
	db.Exec("delete from relationships where user = ? and guest = ?", guest, user)

}

func (r Relationship) CancelAllStatus(user string) {

	db.Exec("delete from relationships where user = ?", user)
	db.Exec("delete from relationships where guest = ?", user)

}

func (r Relationship) Block(user string, guest string) {

	db.Exec("update relationships set status = ? where user = ? and guest = ?", utils.Block, user, guest)
	db.Exec("update relationships set status = ? where user = ? and guest = ?", utils.BeBlocked, guest, user)

}

func (r Relationship) GetRelationship(user string) []Relationship {

	result := make([]Relationship, 0)
	db.Raw("select * from relationships where user = ?", user).Scan(&result)
	return result

}
