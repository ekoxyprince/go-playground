package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func PostUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := Response{
			Success:    false,
			Message:    "Method not Allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		resp := Response{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	name := r.FormValue("imageName")
	file, header, err := r.FormFile("image")
	if err != nil {
		resp := Response{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer file.Close()
	dst, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		resp := Response{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		resp := Response{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := Response{
		Success:    true,
		Message:    "File upload successful " + name,
		StatusCode: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func PostUploads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := Response{
			Success:    false,
			Message:    "Method not Allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		resp := Response{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	name := r.FormValue("imageName")
	files := r.MultipartForm.File["images"]
	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			resp := Response{
				Success:    false,
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		defer file.Close()
		dst, err := os.Create("./uploads/" + header.Filename)
		if err != nil {
			resp := Response{
				Success:    false,
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			resp := Response{
				Success:    false,
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	resp := Response{
		Success:    true,
		Message:    "File upload successful " + name,
		StatusCode: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
