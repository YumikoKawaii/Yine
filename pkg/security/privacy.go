package security

import (
	"github.com/YumikoKawaii/Yine/pkg/models/relationship"
	"github.com/YumikoKawaii/Yine/pkg/models/setting"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var (
	Relationship relationship.Relationship
	Setting      setting.Setting
)

func IsAccessable(id string, guest string) bool {

	relationship := Relationship.GetRelationshipBetween(id, guest)
	public := Setting.GetStranger(guest)

	return relationship == utils.Friend || (relationship != utils.Block && relationship != utils.BeBlocked && public)

}
