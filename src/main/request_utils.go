package main

import (
	"net/http"
	"encoding/json"
)


func getDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		blocks := get_blocks()
		rooms := get_rooms()
		response := ResponseGetAll{true, blocks,rooms}
		productJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()
}
//
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		request := RequestForUpdate{}
		err := decoder.Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}else {
			wipe_table()
		}
		for i := 0; i < len(request.Blocks); i++{
			add_block(request.Blocks[i])
		}
		for i := 0; i < len(request.Rooms); i++{
			add_room(request.Rooms[i])
		}
		response := ResponsePost{true}
		productJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productJson)
	} else {
		//productJson, _ := json.Marshal(users)
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		//w.Write(productJson)
	}
	defer r.Body.Close()
}