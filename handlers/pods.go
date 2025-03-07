package handlers

import (
	"encoding/json"
	"net/http"

	"k8s-pod-deployer/k8s"

	"github.com/gorilla/mux"
	v1 "k8s.io/api/core/v1"
)

func CreatePod(w http.ResponseWriter, r *http.Request) {
	var pod v1.Pod
	if err := json.NewDecoder(r.Body).Decode(&pod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPod, err := k8s.CreatePod(&pod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdPod)
}

func GetPod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	pod, err := k8s.GetPod(name, "default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pod)
}

func DeletePod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if err := k8s.DeletePod(name, "default"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
