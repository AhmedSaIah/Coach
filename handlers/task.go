package handlers

import (
	"TaskManager/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.DB.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(task); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	if err := h.DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(tasks); if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	if err := h.DB.First(&task, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(task); if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	if err := h.DB.First(&task, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input models.Task
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Completed = input.Completed

	if err := h.DB.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(task); if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := h.DB.Delete(&models.Task{}, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
