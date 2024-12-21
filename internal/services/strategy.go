package services

import (
	"github.com/syntaxsdev/mercury/models"
	"go.mongodb.org/mongo-driver/bson"
)

type StrategyService struct {
	db *MongoService
}

// Create new StrategyService
func NewStrategyService(db *MongoService) *StrategyService {
	return &StrategyService{db: db}
}

// Get All Strategies
func (s *StrategyService) GetAllStrategies() ([]models.Strategy, error) {
	var results []models.Strategy
	err := s.db.All("strategies", bson.M{}, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// // Create A Strategy
// func CreateStrategy(m *MongoService) (*models.Strategy, error) {
// 	var newStrategy models.Strategy
// 	if err := json.NewDecoder(r.Body).Decode(&newStrategy); err != nil {
// 		http.Error(w, "Invalid payload", http.StatusBadRequest)
// 		return
// 	}
// 	res, err := f.MongoService.Insert("strategies", newStrategy)
// 	if err != nil {
// 		http.Error(w, "Could not insert strategy", http.StatusInternalServerError)
// 		return
// 	}

// 	WriteHttp(w, http.StatusCreated, "Strategy successfully created", res)
// }
