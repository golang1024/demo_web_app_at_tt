package mdl

import (
	"demo_web_app/db"
	"demo_web_app/const"

	"fmt"
	"strconv"
)


type tbl_user struct {
	Aid   int `sql:",pk"`
	UserName string
}

type UserModel struct {
	UserId	string	`json:"id"`
	UserName	string	`json:"name"`
}

//func getShardAndAid(uid string) (int,int) {
//	if len(uid) != 10 {
//		return -1, -1
//	}
//	shard, err := strconv.Atoi(uid[0:2])
//	if err != nil {
//		return -1, -1
//	}
//	aid, err := strconv.Atoi(uid[:2])
//	if err != nil {
//		return -1, -1
//	}
//	return shard, aid
//}
//
//func getUserId(shard, aid int) string {
//	return fmt.Sprintf("%2d%8d", shard, aid)
//}
//
//func randShard() int {
//	s := rand.NewSource(time.Now().Unix())
//	return rand.New(s).Intn(100)
//}

func getDbName() string {
	return _const.DB_NAME
	//do not shard in demo
	//return fmt.Sprintf("demo_%2o", shard)
}

//we can set a formate for uid
func getUserId(aid int) string {
	return fmt.Sprintf("%d", aid)
}

//get a user by id
func GetUserById(uid string) (*UserModel, error) {
	//shard, aid := getShardAndAid(uid)
	//if shard < 0 {
	//	return nil, errors.New("unknow uid")
	//}
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}
	user := &tbl_user{Aid: aid}
	if err := pgdb.Select(user); err != nil {
		return nil, err
	}
	return &UserModel{UserId: uid, UserName:user.UserName,}, nil
}

//save a user's info
func AddUser(name string) (string, error) {
	//shard := randShard()
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return "", err
	}
	user := &tbl_user{UserName: name}
	if err = pgdb.Insert(user); err != nil {
		return "", err
	}
	return getUserId(user.Aid), nil
}

//get user info list
func GetUserList(lastId string) ([]*UserModel, error) {
	pgdb, err := db.GetDbConn(getDbName())
	if err != nil {
		return nil, err
	}
	users := []*tbl_user{}
	err = pgdb.Model(&users).Where("aid > ? ", lastId).Order("aid asc").Limit(100).Select();
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	ret := make([]*UserModel, len(users))
	for idx, i := range users {
		//ret = append(ret, &UserModel{UserId:getUserId(i.Aid), UserName:i.UserName})
		ret[idx] = &UserModel{UserId:getUserId(i.Aid), UserName:i.UserName}
	}
	return ret, nil
}