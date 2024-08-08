package handlers

import "net/http"


func HandlerReadiness(w http.ResponseWriter, r *http.Request){
    type response struct {
        Status string `json:"status"`
    }
    respondWithJSON(w,http.StatusOK,response{
        Status: "ok",
    })
}

func HandleErr(w http.ResponseWriter, r *http.Request) {
    respondWithError(w,http.StatusInternalServerError,"Internal Server Error")

}
