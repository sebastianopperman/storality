package admin

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dashboard")
}

func GetNewPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Viewing new page")
}

func GetPageView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Viewing page %d", id)
}

func PostPageNew(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Creating new page...")
}

func DeletePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Page %d has been deleted.", id)
}