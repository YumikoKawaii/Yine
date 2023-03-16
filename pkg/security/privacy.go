package security

import (
	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Relationship models.Relationship
var Setting models.Setting

func IsAccessable(id string, guest string) bool {

	relationship := Relationship.GetRelationshipBetween(id, guest)
	public := Setting.GetStranger(guest)

	return relationship == utils.Friend || (relationship != utils.Block && relationship != utils.BeBlocked && public)

}
