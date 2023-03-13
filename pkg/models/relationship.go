package models

import "github.com/YumikoKawaii/Yine/pkg/utils"

type Relationship struct {
	ID     string `json:"id" gorm:"primarykey"`
	Guest  string `json:"guest"`
	Status string `json:"status"`
}

func init() {
	db.AutoMigrate(&Relationship{})
}

func (r Relationship) GetRelationship(id string, guest string) string {

	result := ""
	db.Raw("select status from relationships where id = ? and guest = ?", id, guest).Scan(&result)
	return result

}

func (r Relationship) ProcessRequest(id string, guest string) {

	i_record := &Relationship{
		ID:     id,
		Guest:  guest,
		Status: utils.SentRequest,
	}

	g_record := &Relationship{
		ID:     guest,
		Guest:  id,
		Status: utils.GotRequest,
	}

	db.Create(i_record)
	db.Create(g_record)

}

func (r Relationship) AcceptRequest(id string, guest string) {

	db.Exec("update relationships set status = ? where id = ? and guest = ?", utils.Friend, id, guest)
	db.Exec("update relationships set status = ? where id = ? and guest = ?", utils.Friend, guest, id)

}

func (r Relationship) CancelStatus(id string, guest string) {

	db.Exec("delete from relationships where id = ? and guest = ?", id, guest)
	db.Exec("delete from relationships where id = ? and guest = ?", guest, id)

}

func (r Relationship) CancelAllStatus(id string) {

	db.Exec("delete from relationships where id = ?", id)
	db.Exec("delete from relationships where guest = ?", id)

}

func (r Relationship) Block(id string, guest string) {

	db.Exec("update relationships set status = ? where id = ? and guest = ?", utils.Block, id, guest)
	db.Exec("update relationships set status = ? where id = ? and guest = ?", utils.BeBlocked, guest, id)

}
