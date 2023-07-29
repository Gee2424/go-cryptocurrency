package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response struct for displaying the response of the POST request
type Response struct {
	Index    int    `json:"index"`
	Message  string `json:"message"`
	BPM      int    `json:"BPM"`
	PrevHash string `json:"prev_hash"`
	Hash     string `json:"hash"`
}

// handleGetBlockchain handles the GET request to return the blockchain
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// handleWriteBlock handles the POST request to add a new block
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var message Response

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], message.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, message)
		return
	}
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		respondWithJSON(w, r, http.StatusCreated, newBlock)
	}
	mutex.Unlock()
}

// respondWithJSON sends JSON response
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// Run starts the HTTP server
func Run() error {
	http.HandleFunc("/get_blockchain", handleGetBlockchain)
	http.HandleFunc("/write_block", handleWriteBlock)
	return http.ListenAndServe(":8080", nil)
}
