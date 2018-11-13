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

	llog.Error(rating < -9000, "Error: unknown parameter '"+ratingName+"'")
	// if rating < -9000 {
	// 	return rating, errors.New("Error: unknown parameter '" + ratingName + "'")
	// }
	return rating, nil
}

//SetRating - Изменяет числовое значение рейтинга по имени
func (s *Syndicate) SetRating(ratingName string, newRating int) error {
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
		return errors.New("Error: unknown parametr '" + ratingName + "'")
	}
	return nil
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
	}
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
	}
}

//AdjustedRating -
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
	report := "Syndicate Report: " + s.Name + "\n"
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

func (s *Syndicate) efficiencyTest() {
	efficiency, err := s.Rating("Fiscal")
	handleError(err)
	hits, outcome, _, gl := sr3SimpleTest(efficiency, 4)
	fmt.Println("eff test:", hits, outcome, gl)
	fmt.Println("WILL INCREASE", hits, "ASSETS")
	for marketName, marketRating := range s.Market {
		if hits < 1 {
			continue
		}
		if marketRating > 9 {
			fmt.Println("                      CAN'T INCREACE BEOYND 10", marketName)
			continue
		}
		s.SetRating(marketName, marketRating+1)
		hits--
		fmt.Println(marketName, "increased")
	}
}

func (s *Syndicate) publicityTest() {
	publicity, err := s.Rating("Reputation")
	handleError(err)
	hitsP, outcome, _, gl := sr3SimpleTest(publicity, 8)
	fmt.Println("Pub test:", hitsP, outcome, gl)
	fmt.Println("WILL SAVE", hitsP, "ASSETS")
	degradeRound := 3
	degradeRound = degradeRound - hitsP
	fmt.Println(degradeRound, "WILL DEGRADE")
	for marketName, marketRating := range s.Market {
		if degradeRound < 1 {
			break
		}
		if marketRating < 1 {
			continue
		}
		s.SetRating(marketName, marketRating-1)
		degradeRound--
		fmt.Println(marketName, "Degraded")
	}
}

func (s *Syndicate) blackOpsTest() {
	target := pickTarget(s, AllSyndicates)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println("Plotting Shadowrun against", market, "of", target, "sponsored by", s.Name)
	blOps, _ := s.Rating("Enforcement")
	targSecurity, _ := AllSyndicates[target].Rating("Enforcement")
	hits, outcome, _, gl := sr3SimpleTest(blOps, targSecurity)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
		fmt.Println("          ", s.Name, market, "decreased")
	}
	if hits > 1 {
		s.increaseRating(market)
		fmt.Println("          ", s.Name, market, "increased")
		AllSyndicates[target].decreaseRating(market)
		fmt.Println("          ", target, market, "decreased")
	}
}

func (s *Syndicate) intelTest() {
	target := pickTarget(s, AllSyndicates)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println("Plotting Shadowrun against", market, "of", target, "sponsored by", s.Name)
	blOps, _ := s.Rating("Intelligence")
	targSecurity, _ := AllSyndicates[target].Rating("Management")
	hits, outcome, _, gl := sr3SimpleTest(blOps, targSecurity)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
		fmt.Println("          ", s.Name, market, "decreased")
	}
	if hits > 1 {
		s.increaseRating(market)
		fmt.Println("          ", s.Name, market, "increased")
		AllSyndicates[target].decreaseRating(market)
		fmt.Println("          ", target, market, "decreased")
	}
}

func (s *Syndicate) naturalCycle() {
	fmt.Println("  ", s.Name, "EffTest")
	s.efficiencyTest()
	fmt.Println("  ", s.Name, "PubTest")
	s.publicityTest()
	//ChooseTarget For BlackOpsTest
	s.blackOpsTest()
	//ChooseTarget For Intel
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
				fmt.Println(source, key, val, target, key, val2)
				return key
			}
		}
	}
	return ""
}
