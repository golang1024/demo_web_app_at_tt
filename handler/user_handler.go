package handler

import (
	"demo_web_app/util"
	"demo_web_app/srv"
	"demo_web_app/enum"

	"net/http"

	"github.com/gorilla/mux"


)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user,err := srv.GetUserById(vars["id"])
	if err != nil {
		util.JsonReturnErr(err.Error(), http.StatusBadRequest, w)
		return
	}
	util.JsonReturnOK(user, w)
}


func GetUserList(w http.ResponseWriter, r *http.Request) {
	lastId := util.GetRequestForm(r, "lastId", "0")
	if users, err := srv.GetAllUsers(lastId); err != nil {
		util.JsonReturnErr(nil, http.StatusInternalServerError, w)
	} else {
		util.JsonReturnOK(users, w)
	}
}

func PostUser(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{}
	err := util.GetJsonBody(r, &params)
	if err != nil {
		util.JsonReturnErr(nil, http.StatusBadRequest, w)
		return
	}

	if len(params) <= 0 {
		util.JsonReturnErr(nil, http.StatusBadRequest, w)
		return
	}
	if _, ok := params["name"]; !ok {
		util.JsonReturnErr(nil, http.StatusBadRequest, w)
		return
	}

	x := &srv.UserModel{}
	x.UserName = params["name"]
	if users, err := srv.AddUser(x); err != nil {
		util.JsonReturnErr(err.Error(), http.StatusInternalServerError, w)
	} else {
		util.JsonReturnOK(users, w)
	}
}

func PutRelation(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{}
	err := util.GetJsonBody(r, &params)
	vars := mux.Vars(r)
	if err != nil || vars == nil || len(params) <= 0 {
		util.JsonReturnErr(1, http.StatusBadRequest, w)
		return
	}

	_, ok1 := params["state"]
	_, ok2 := vars["ida"]
	_, ok3 := vars["idb"]

	if !ok1 || !ok2 || !ok3 {
		util.JsonReturnErr(2, http.StatusBadRequest, w)
		return
	}
	var state enum.RelationState
	switch params["state"] {
	case "liked":
		state = enum.E_Liked
	case "disliked":
		state = enum.E_Disliked
	default:
		util.JsonReturnErr(3, http.StatusBadRequest, w)
		return
	}

	if r, err := srv.AddRelation(vars["ida"], vars["idb"], state); err != nil {
		util.JsonReturnErr(err.Error(), http.StatusInternalServerError, w)
	} else if r != nil && r.Relation != nil && len(r.Relation) > 0 {
		util.JsonReturnOK(r.Relation, w)
	} else {
		util.JsonReturnOK(nil, w)
	}
}

func GetRelation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	if vars == nil {
		util.JsonReturnErr(nil, http.StatusBadRequest, w)
		return
	}

	_, ok := vars["id"]

	if !ok {
		util.JsonReturnErr(nil, http.StatusBadRequest, w)
		return
	}

	if r, err := srv.GetRelation(vars["id"]); err != nil  {
		util.JsonReturnErr(err.Error(), http.StatusInternalServerError, w)
	} else if r != nil && r.Relation != nil && len(r.Relation) > 0 {
		util.JsonReturnOK(r.Relation, w)
	} else {
		util.JsonReturnOK(nil, w)
	}
}

