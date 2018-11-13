package main

import (
	"errors"
	"fmt"
	"strconv"
)

//type AllSyndicates map[string]*Syndicate

//Syndicate - представляет собой ОРГАНИЗАЦИЮ среднего порядка с экономикой построенной на криминале. имеет 2 типа параметров Operation и Market
type Syndicate struct {
	Name       string
	Operation  map[string]int
	Market     map[string]int
	Adjustment map[string]int
	err        error
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
		s.Market[markets[i]] = randInt(0, 10)
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
func (s *Syndicate) Rating(ratingName string) int {
	if val, ok := s.Market[ratingName]; ok {
		return val
	}
	if val, ok := s.Operation[ratingName]; ok {
		return val
	}
	s.err = errors.New("Error: unknown parametr '" + ratingName + "'")
	//llog.Error(rating < -9000, "Error: unknown parameter '"+ratingName+"'")
	return -1
}

//SetRating - Изменяет числовое значение рейтинга по имени
func (s *Syndicate) SetRating(ratingName string, newRating int) {
	if s.err != nil {
		return
	}
	found := false
	if _, ok := s.Market[ratingName]; ok {
		found = true
		s.Market[ratingName] = newRating
	}
	if _, ok := s.Operation[ratingName]; ok {
		found = true
		s.Market[ratingName] = newRating
	}
	if !found {
		s.err = errors.New("Error: unknown parametr '" + ratingName + "'")
	}
}

func (s *Syndicate) increaseRating(ratingName string) {
	found := false
	if _, ok := s.Market[ratingName]; ok {
		found = true
		if s.Market[ratingName] < 10 {
			s.Market[ratingName] = s.Market[ratingName] + 1
		}
	}
	if _, ok := s.Operation[ratingName]; ok {
		found = true
		if s.Market[ratingName] < 10 {
			s.Market[ratingName] = s.Market[ratingName] + 1
		}
	}
	if !found {
		s.err = errors.New("Error: unknown parametr '" + ratingName + "'")
		return
	}
	fmt.Println(s.Name, ratingName, "increased...")
}

func (s *Syndicate) decreaseRating(ratingName string) {
	found := false
	if _, ok := s.Market[ratingName]; ok {
		found = true
		if s.Market[ratingName] > 0 {
			s.Market[ratingName] = s.Market[ratingName] - 1
		}
	}
	if _, ok := s.Operation[ratingName]; ok {
		found = true
		if s.Market[ratingName] > 0 {
			s.Market[ratingName] = s.Market[ratingName] - 1
		}
	}
	if !found {
		s.err = errors.New("Error: unknown parametr '" + ratingName + "'")
		return
	}
	fmt.Println(s.Name, ratingName, "decreased...")
}

//FullReport - Возвращает форматированную стену текста с описанием всех характеристик.
func (s *Syndicate) FullReport() string {
	netRating := 0
	report := "Syndicate Report: " + s.Name + "\n"
	report = report + "Markets" + "\n"

	markets := allMarkets()
	for i := range markets {
		marketRep := s.Rating(markets[i])

		report = report + markets[i] + ": " + strconv.Itoa(marketRep) + "\n"
		marketRating := s.Rating(markets[i])
		netRating = netRating + marketRating
	}
	report = report + "--------------------" + "\n"
	report = report + "Operations" + "\n"

	operations := allOperations()
	for i := range operations {
		operationsRep := s.Rating(operations[i])

		report = report + operations[i] + ": " + strconv.Itoa(operationsRep) + "\n"
		operationsRating := s.Rating(operations[i])
		netRating = netRating + operationsRating
	}
	report = report + "--------------------" + "\n"
	report = report + "Total Net Rating: " + strconv.Itoa(netRating) + "\n"
	return report
}

func (s *Syndicate) efficiencyTest() {

	fmt.Println("Run " + s.Name + " Efficiency test:")
	efficiency := s.Rating("Fiscal")
	hits, _, _, _ := sr3SimpleTest(efficiency, 4)
	for marketName, marketRating := range s.Market {
		if hits < 1 {
			continue
		}
		if marketRating > 9 {
			continue
		}
		s.increaseRating(marketName)
		hits--
	}
	fmt.Println("")
}

func (s *Syndicate) publicityTest() {
	fmt.Println("Run " + s.Name + " Publicity test:")
	publicity := s.Rating("Reputation")
	hitsP, _, _, _ := sr3SimpleTest(publicity, 8)
	degradeRound := 3
	degradeRound = degradeRound - hitsP
	for marketName, marketRating := range s.Market {
		if degradeRound < 1 {
			break
		}
		if marketRating < 1 {
			continue
		}
		s.decreaseRating(marketName)
		degradeRound--
	}
	fmt.Println("")
}

func (s *Syndicate) blackOpsTest() {
	target := pickTarget(s, AllSyndicates)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println(s.Name+" plotting Shadowrun against", market, "of", target)
	blOps := s.Rating("Enforcement")
	targSecurity := AllSyndicates[target].Rating("Enforcement")
	hits, outcome, _, gl := sr3SimpleTest(blOps, targSecurity)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
	}
	if hits > 1 {
		s.increaseRating(market)
		AllSyndicates[target].decreaseRating(market)
	}
	fmt.Println("")
}

func (s *Syndicate) intelTest() {
	target := pickTarget(s, AllSyndicates)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println(s.Name+" plotting Shadowrun against", market, "of", target)
	intel := s.Rating("Intelligence")
	targManagement := AllSyndicates[target].Rating("Management")
	hits, outcome, _, gl := sr3SimpleTest(intel, targManagement)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
	}
	if hits > 1 {
		s.increaseRating(market)
		AllSyndicates[target].decreaseRating(market)
	}
	fmt.Println("")
}

func (s *Syndicate) naturalCycle() {
	s.efficiencyTest()
	s.publicityTest()
	s.blackOpsTest()
	s.intelTest()
}

func pickTarget(s *Syndicate, AllSyndicates map[string]*Syndicate) (target string) {
	var targetList []string
	for key := range AllSyndicates {
		if key == s.Name {
			continue
		}
		targetList = append(targetList, key)
	}
	r := randInt(1, len(targetList))
	return targetList[r-1]
}

func pickCommonRandomMarket(source, target string) string {
	sin1 := AllSyndicates[source]
	for key, val := range sin1.Market {
		if val != 0 {
			sin2 := AllSyndicates[target]
			val2, ok := sin2.Market[key]
			if val2 != 0 && ok {
				return key
			}
		}
	}
	return ""
}
