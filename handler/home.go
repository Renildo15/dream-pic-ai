package handler

import (
	"dreampicai/view/home"
	"fmt"
	"net/http"
	"time"
)

func HandleLongProcess(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Second * 5)
	return home.UserLikes(10000).Render(r.Context(), w)
}

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAutheticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	fmt.Printf("%+v\n", user.Account)
	return home.Index().Render(r.Context(), w)
}