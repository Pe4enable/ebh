package handlers

import (
	"log"
	"net/http"
)


func (s *HandlersService) CreateSert(w http.ResponseWriter, r *http.Request) {


	result, err := s.repository.CreateSertificate()

	// IMPORTANT !!
	// value, gasPrice is in "WEI" points // 1.000.000.000.000.000.000 wei = 1 ETH
	// It's mean that gasPrice = 15000000000 equal gasPrice = 0.000000015 ETH
	//result, err := s.service.SendTransaction(from, to, valIntBig, 30000, 15000000000, password)

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, result, http.StatusOK)
}



func (s *HandlersService) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := s.repository.GetAllSertificates()

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, result, http.StatusOK)
}