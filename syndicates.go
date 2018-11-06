package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Syndicate - представляет собой ОРГАНИЗАЦИЮ среднего порядка с экономикой построенной на криминале. имеет 2 типа параметров Operation и Market
type Syndicate struct {
	Name      string
	Operation map[string]int
	Market    map[string]int
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
	return &s
}

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
