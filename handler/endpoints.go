package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-pgdb/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (env *Env) GetAllForums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllForums")
	var forums = env.PGDB.GetAllForums()
	sendResponse(w, forums)
}

func (env *Env) CreateForum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateForum")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var forum models.Forum
	err := json.Unmarshal(reqBody, &forum)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var updatedForum = env.PGDB.CreateForum(forum)
	sendResponse(w, updatedForum)
}

func (env *Env) UpdateForum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateForum")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedForum models.Forum
	err := json.Unmarshal(reqBody, &updatedForum)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	updatedForum, err = env.PGDB.UpdateForum(updatedForum)
	if err != nil {
		http.Error(w, "Error Updating Forum", http.StatusBadRequest)
		return
	} else {
		sendResponse(w, updatedForum)
	}
}

func (env *Env) GetForumById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetForumById")
	vars := mux.Vars(r)
	forumId := vars["id"]
	id, err := strconv.Atoi(forumId)
	if err == nil {
		fmt.Println("Endpoint Hit: GetForumById " + forumId)
		var claimAuthorityLimitSummaries = env.PGDB.GetForumById(id)
		sendResponse(w, claimAuthorityLimitSummaries)
	}
}

func (env *Env) GetThreadsByForumId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetThreadsByForumId")
	vars := mux.Vars(r)
	forumId := vars["id"]
	id, err := strconv.Atoi(forumId)
	if err == nil {
		fmt.Println("Endpoint Hit: GetThreadsByForumId " + forumId)
		threads := env.PGDB.GetThreadsByForumId(id)
		sendResponse(w, threads)
	}
}

func (env *Env) GetPostsByThreadId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetPostsByThreadId")
	vars := mux.Vars(r)
	threadId := vars["id"]
	id, err := strconv.Atoi(threadId)
	if err == nil {
		fmt.Println("Endpoint Hit: GetPostsByThreadId " + threadId)
		posts := env.PGDB.GetPostsByThreadId(id)
		sendResponse(w, posts)
	}
}

func (env *Env) GetPostsByForumId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetPostsByForumId")
	vars := mux.Vars(r)
	forumId := vars["id"]
	id, err := strconv.Atoi(forumId)
	if err == nil {
		fmt.Println("Endpoint Hit: GetPostsByForumId " + forumId)
		posts := env.PGDB.GetPostsByForumId(id)
		sendResponse(w, posts)
	}
}

func sendResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}