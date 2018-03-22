package mdl

import (
	"demo_web_app/enum"
	"demo_web_app/db"


	"fmt"
	"strconv"
)

type tbl_relation struct {
	UidA int	`sql:",pk"`
	UidB int	`sql:",pk"`
	StateA  enum.RelationState
	StateB  enum.RelationState
}

type RelationItem struct {
	UserId string             `json:"user_id"`
	State  enum.RelationState `json:"state"`
	Type   enum.DataType      `json:"type"`
}

type Relation struct {
	UserId string
	Relation []*RelationItem
}


func ConvertOneUser(rs []*tbl_relation) (ret *Relation) {
	if rs == nil || len(rs) == 0 {
		return nil
	}
	ret = &Relation{Relation: []*RelationItem{}}
	for _, r := range rs {
		ret.UserId = fmt.Sprintf("%d", r.UidA)
		i := &RelationItem{
			UserId: fmt.Sprintf("%d", r.UidB),
			State:  CombineState(r.StateA, r.StateB),
			Type:   enum.E_TYPE_RELATION,
		}
		if i.State == enum.E_Default {
			continue
		}
		ret.Relation = append(ret.Relation, i)
	}
	return
}

func CombineState(stateA, stateB enum.RelationState) enum.RelationState {
	if stateB == enum.E_Liked && stateA == enum.E_Liked {
		return enum.E_Matched
	}
	return stateA
}

func setStateA(userIdA, userIdB int, state enum.RelationState) error {
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return err
	}
	r := &tbl_relation{UidA: userIdA, UidB:userIdB, StateA: state}
	_, err = pgdb.Model(r).OnConflict("(uida, uidb) DO UPDATE").Set("statea = EXCLUDED.statea").Insert()
	return err
}

func setStateB(userIdA, userIdB int, state enum.RelationState) error {
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return err
	}
	r := &tbl_relation{UidA: userIdA, UidB:userIdB, StateB: state}
	_, err = pgdb.Model(r).OnConflict("(uida, uidb) DO UPDATE").Set("stateb = EXCLUDED.stateb").Insert()
	return err
}


func SetState(userIdA, userIdB string, state enum.RelationState) error {
	uida, err := strconv.Atoi(userIdA)
	if err != nil {
		return err
	}
	uidb, err := strconv.Atoi(userIdB)
	if err != nil {
		return err
	}
	if err := setStateA(uida, uidb, state); err != nil {
		return err
	}
	return setStateB(uidb, uida, state)
}

func GetRelationBetween(userIdA, userIdB string) (*Relation, error) {
	uida, err := strconv.Atoi(userIdA)
	if err != nil {
		return nil, err
	}
	uidb, err := strconv.Atoi(userIdB)
	if err != nil {
		return nil, err
	}
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return nil, err
	}
	rs := []*tbl_relation{}
	err = pgdb.Model(&rs).Where("uida = ? and uidb = ?", uida, uidb).Select()
	if err != nil {
		return nil, err
	}

	return ConvertOneUser(rs), nil
}

func GetRelations(userId string) (*Relation, error){
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return nil, err
	}
	rs := []*tbl_relation{}
	err = pgdb.Model(&rs).Where("uida = ? and statea != ?", uid, enum.E_Default).Select()
	if err != nil {
		return nil, err
	}

	return ConvertOneUser(rs), nil
}


