package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"week02/dao"
	"week02/service"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		user, err := service.UserService.One(id)
		if err == dao.RecordNotFound {
			fmt.Printf("stack trace: \n%+v\n", err)
			_, _ = w.Write([]byte("用户不存在"))
			return
		}
		if err != nil {
			log.Println(fmt.Sprintf("%+v", err))
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, _ := json.Marshal(user)
		_, _ = w.Write(data)
	})
	mux.HandleFunc("/rows", func(w http.ResponseWriter, r *http.Request) {
		res, err := service.UserService.Rows()
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, _ := json.Marshal(res)
		_, _ = w.Write(data)
	})
	log.Fatal(http.ListenAndServe(":8888", mux))
}
