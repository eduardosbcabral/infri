package main

import (
	"context"
	"encoding/json"
	"github.com/eduardosbcabral/infri/models"
	"github.com/eduardosbcabral/infri/node"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Server struct {
	repository node.Repository
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	node := r.URL.Path[len("/nodes/"):]
	switch r.Method {
	case http.MethodPost:
		s.StoreNode(w, r)
	case http.MethodGet:
		s.GetNode(w, node)
	}
}

func (s *Server) StoreNode(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var node models.Node

	err = json.Unmarshal(b, &node)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ctx := context.Background()

	s.repository.Store(ctx, &node)

	w.WriteHeader(200)
}

func (s *Server) GetNode(w http.ResponseWriter, nodeID string) {
	id, err := strconv.ParseInt(nodeID, 10, 64)
	if err != nil {
		w.Write([]byte("id is not integer"))
		w.WriteHeader(500)
	}

	ctx := context.Background()

	node, err := s.repository.GetById(ctx, id)
	if err != nil {
		w.Write([]byte("node not found"))
		w.WriteHeader(500)
	}

	nodeM, _ := json.Marshal(node)

	w.Write(nodeM)
	w.WriteHeader(200)
}