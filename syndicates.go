package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Syndicate - представляет собой ОРГАНИЗАЦИЮ среднего порядка с экономикой построенной на криминале. имеет 2 типа параметров Operation и Market
type Syndicate struct {
	Name       string
	Operation  map[string]int
	Market     map[string]int
	Adjustment map[string]int
}

func allMarkets() []string {
	markets := []string{
		"Computer Crime",
		"Controlled Substanses",
		"Counterfeiting and Forgery",
		"Fencing",
		"Fixing",
		"Gambling",
		"Hijacking",
		"Kidnapping",
		"Loansharking",
		"Pornography",
		"Prostitution",
		"Protection",
		"Robbery",
		"Shadow Servises",
		"Smuggling",
		"White-Collar Crime",
	}
	return markets
}

func allOperations() []string {
	markets := []string{
		"Enforcement",
		"Fiscal",
		"Intelligence",
		"Management",
		"Reputation",
	}
	return markets
}

//NewSyndicate - создает рандомный синдикат со случайными параметрами.
//TODO: придумать схему контролируемого создания организаций.
func NewSyndicate(name string) *Syndicate {
	s := Syndicate{}
	s.Name = name
	s.Market = make(map[string]int)
	markets := allMarkets()
	for i := range markets {
		s.Market[markets[i]] = randInt(1, 10)
	}
	s.Operation = make(map[string]int)
	operations := allOperations()
	for i := range operations {
		s.Operation[operations[i]] = randInt(1, 10)
	}
	s.Adjustment = make(map[string]int)
	return &s
}

//Rating - Возвращает числовое значение рейтинга по имени
func (s *Syndicate) Rating(ratingName string) (int, error) {
	rating := -9999
	if val, ok := s.Market[ratingName]; ok {
		rating = val
	}
	if val, ok := s.Operation[ratingName]; ok {
		rating = val
	}
	if rating < -9000 {
		return rating, errors.New("Error: unknown parametr '" + ratingName + "'")
	}
	return rating, nil
}

func (s *Syndicate) AdjustedRating(ratingName string) (int, error) {
	rating, ratError := s.Rating(ratingName)
	if ratError != nil {
		return -9999, ratError
	}
	if adj, ok := s.Adjustment[ratingName]; ok {
		rating = rating + adj
	}
	return rating, nil
}

func (s *Syndicate) reportRating(ratingName string) (string, error) {
	report := ""
	if val, ok := s.Market[ratingName]; ok {
		report = ratingName + ": " + strconv.Itoa(val)
	}
	if val, ok := s.Operation[ratingName]; ok {
		report = ratingName + ": " + strconv.Itoa(val)
	}
	if report == "" {
		return report, errors.New("Error: unknown parametr '" + ratingName + "'")
	}
	return report, nil
}

//FullReport - Возвращает форматированную стену текста с описанием всех характеристик.
func (s *Syndicate) FullReport() string {
	netRating := 0
	report := "Syndicate Report:" + "\n"
	report = report + "Markets" + "\n"

	markets := allMarkets()
	for i := range markets {
		marketRep, err := s.reportRating(markets[i])
		if err != nil {
			fmt.Println(err)
		}
		report = report + marketRep + "\n"
		marketRating, _ := s.Rating(markets[i])
		netRating = netRating + marketRating
	}
	report = report + "--------------------" + "\n"
	report = report + "Operations" + "\n"

	operations := allOperations()
	for i := range operations {
		operationsRep, err := s.reportRating(operations[i])
		if err != nil {
			fmt.Println(err)
		}
		report = report + operationsRep + "\n"
		operationsRating, _ := s.Rating(operations[i])
		netRating = netRating + operationsRating
	}
	report = report + "--------------------" + "\n"
	report = report + "Total Net Rating: " + strconv.Itoa(netRating) + "\n"
	return report
}

func (s *Syndicate) NaturalCycle() {
	for key, val := range s.Market {

		adjRating, _ := s.AdjustedRating(key)
		hits, outcome, resultArray, glitch := sr3SimpleTest(adjRating+randInt(0, 3), val)
		fmt.Println(key, hits, outcome, resultArray, glitch)
		if hits == 0 {
			change := (randInt(1, 6) + 1) / 2
			fmt.Println("   ", key, "reduced by", change)
			s.Market[key] = s.Market[key] - change
			if s.Market[key] < 0 {
				s.Market[key] = 0
			}
			continue
		}
		if hits > 0 && hits < 4 {
			//fmt.Println(key, "Base Rating remains the same")
			continue
		}
		if hits > 3 && hits < 7 {
			change := 1
			fmt.Println("   ", key, "increased by", change)
			s.Market[key] = s.Market[key] + change
			continue
		}
		if hits > 6 {
			change := 2
			fmt.Println("   ", key, "increased by", change)
			s.Market[key] = s.Market[key] + change
			continue
		}
	}
}

func (s *Syndicate) assessEffects() {
	marketAct := randInt(0, 3)
	operationsAct := randInt(0, 1)
	for market, _ := range s.Market {
		if marketAct < 1 {
			break
		}
		s.Adjustment[market] = randInt(-1, 0)

	}
	for operation, _ := range s.Operation {
		if operationsAct < 1 {
			break
		}
		s.Adjustment[operation] = randInt(-1, 0)

	}
}

func (s *Syndicate) determineTestOrder() (order []string) {
	resourcePool := s.resourcePool()

	//any Markets that were directly affected by shadowruns
	for market, _ := range s.Market {
		if _, ok := s.Adjustment[market]; ok {
			order = append(order, market)
			maxAdjResource, _ := s.Rating(market)
			if maxAdjResource < resourcePool {
				maxAdjResource = resourcePool
			}
			resourcePool = resourcePool - maxAdjResource
			s.Adjustment[market] = s.Adjustment[market] + randInt(0, maxAdjResource)
		}
	}

	//each of the syndicate's Operations
	for operation, _ := range s.Operation {
		order = append(order, operation)
		maxAdjResource, _ := s.Rating(operation)
		if maxAdjResource < resourcePool {
			maxAdjResource = resourcePool
		}
		resourcePool = resourcePool - maxAdjResource
		s.Adjustment[operation] = s.Adjustment[operation] + randInt(0, maxAdjResource)
	}
	//any other markets of his or her choice

	return order
}

func (s *Syndicate) resourcePool() int {
	rPool := 0
	for _, val := range s.Operation {
		rPool = rPool + val
	}
	return rPool
}
