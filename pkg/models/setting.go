package models

type Setting struct {
	ID       string `json:"id" gorm:"primarykey"`
	Public   bool   `json:"public"`
	Stranger bool   `json:"stranger"`
}

func init() {
	db.AutoMigrate(&Setting{})
}

func (s Setting) NewSetting(id string) {

	db.Create(Setting{
		ID:       id,
		Public:   true,
		Stranger: true,
	})

}

func (s Setting) ChangePublic(id string, p bool) {

	db.Exec("update settings set public = ? where id = ?", p, id)

}

func (s Setting) ChangeStranger(id string, stranger bool) {

	db.Exec("update settings set stranger = ? where id = ?", stranger, id)

}

func (s Setting) GetPublic(id string) bool {

	result := false
	db.Raw("select public from settings where id = ?", id).Scan(&result)
	return result

}

func (s Setting) GetStranger(id string) bool {

	result := false
	db.Raw("select stranger from settings where id = ?", id).Scan(&result)
	return result

}
