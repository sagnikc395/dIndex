package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	simulator *IndexSimulator
	metrics   *MetricsCollector
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer() {
	http.HandleFunc("/simulate", s.handleSimulation)
	http.HandleFunc("/metrics", s.handleMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) handleSimulation(w http.ResponseWriter, r *http.Request) {
	// Handle simulation requests
}

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	// Return metrics in JSON format
	json.NewEncoder(w).Encode(SimulationResult{})
}
