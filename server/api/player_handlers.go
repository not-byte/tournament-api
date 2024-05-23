package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"tournament_api/server/model"

	"github.com/gorilla/mux"
)

func (s *Server) handleGetTeamPlayers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id := new(big.Int)
	_, err := fmt.Sscan(idStr, id)
	if err != nil {
		http.Error(w, "Invalid team ID "+err.Error(), http.StatusBadRequest)
		return

	}

	players, err := s.store.GetPlayers()
	if err != nil {
		http.Error(w, "Error while fetching players "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]model.PlayerDTO{"players": players})
}

func (s *Server) handleGetAllPlayers(w http.ResponseWriter, r *http.Request) {

	players, err := s.store.GetPlayers()
	if err != nil {
		http.Error(w, "Error while fetching players "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]model.PlayerDTO{"players": players})
}
