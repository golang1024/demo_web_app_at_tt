package srv

import (
	"demo_web_app/enum"
	"demo_web_app/mdl"

	"errors"
)

type UserModel struct {
	mdl.UserModel
	UserType enum.DataType `json:"type"`
}

//get a user by it's id
func GetUserById(uid string) (*UserModel, error) {
	userinfo, err := mdl.GetUserById(uid)
	if err != nil {
		return nil, err
	} else  if userinfo == nil {
		return nil, nil
	}
	return &UserModel{*userinfo, enum.E_TYPE_USER}, nil
}

//get a list with users, who id is over the input id, order by uid asc
func GetAllUsers(lastId string) ([]*UserModel, error) {
	userinfos, err := mdl.GetUserList(lastId)
	if err != nil {
		return nil, err
	} else if userinfos == nil || len(userinfos) == 0 {
		return nil, nil
	}
	ret := []*UserModel{}
	for _, i := range userinfos {
		ret = append(ret, &UserModel{*i, enum.E_TYPE_USER})
	}
	return ret, nil
}

//add a user,return user's id or err
func AddUser(user *UserModel) (string, error) {
	return mdl.AddUser(user.UserName)
}


//create relation from uida to uidb with state
func AddRelation(uida, uidb string, state enum.RelationState) (*mdl.Relation, error) {
	//check if user exist
	if user, err := GetUserById(uida); err != nil || user == nil {
		return nil, errors.New("no such user")
	}
	if user, err := GetUserById(uidb); err != nil || user == nil {
		return nil, errors.New("no such user")
	}
	//write realation into db
	if err := mdl.SetState(uida, uidb, state); err != nil {
		return nil, err
	}
	//get the lastest relation
	return mdl.GetRelationBetween(uida, uidb)
}

//get a user's relation by userid
func GetRelation(uid string) (*mdl.Relation, error) {
	return mdl.GetRelations(uid)
}