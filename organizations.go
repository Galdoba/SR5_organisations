package main

import (
	"errors"
	"fmt"
	"strconv"
)

//type AllOrganizations map[string]*Organization

//Organization - представляет собой ОРГАНИЗАЦИЮ среднего порядка с экономикой построенной на криминале. имеет 2 типа параметров Operation и Market
type Organization struct {
	Name      string
	orgType   string
	Operation map[string]int
	Market    map[string]int
	//Adjustment map[string]int
	err error
}

func allMarkets(orgType string) []string {
	var markets []string
	switch orgType {
	case "Syndicate":
		markets = []string{
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
	case "Gang":
		markets = []string{
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
	case "Corporation":
		markets = []string{
			"Aerospace",
			"Agriculture",
			"Biotech",
			"Chemicals",
			"Comp Engineering",
			"Comp Science",
			"Consumer Goods",
			"Cybernetics",
			"Entertainment",
			"Finance",
			"Heavy Industry",
			"Light Industry",
			"Military Technology",
			"Mystic Goods",
			"Transport",
			"Services",
		}
	}

	return markets
}

func allAttributes() []string {
	attributes := []string{
		"Intelligence",
		"Counter-Intelligence",
		"Reputation",
		"Managment",
		"BlackOps",
		"Security",
	}
	return attributes
}

//NewOrganization - создает рандомный синдикат со случайными параметрами.
//TODO: придумать схему контролируемого создания организаций.
func NewOrganization(name string, orgType string) *Organization {
	s := Organization{}
	s.Name = name
	s.orgType = orgType
	s.Market = make(map[string]int)
	markets := allMarkets(s.orgType)
	var limit int
	switch s.orgType {
	case "Corporation":
		limit = 14
	case "Syndicate":
		limit = 10
	case "Gang":
		limit = 6
	}
	for i := range markets {
		s.Market[markets[i]] = randInt(0, limit)
	}
	s.Operation = make(map[string]int)
	operations := allAttributes()
	for i := range operations {
		s.Operation[operations[i]] = randInt(1, limit)
	}
	return &s
}

//Rating - Возвращает числовое значение рейтинга по имени
func (s *Organization) Rating(ratingName string) int {
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
func (s *Organization) SetRating(ratingName string, newRating int) {
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

func (s *Organization) increaseRating(ratingName string) {
	var limit int
	switch s.orgType {
	case "Corporation":
		limit = 14
	case "Syndicate":
		limit = 10
	case "Gang":
		limit = 6
	}
	found := false
	if _, ok := s.Market[ratingName]; ok {
		found = true
		if s.Market[ratingName] < limit {
			s.Market[ratingName] = s.Market[ratingName] + 1
		}
	}
	if _, ok := s.Operation[ratingName]; ok {
		found = true
		if s.Operation[ratingName] < limit {
			s.Operation[ratingName] = s.Operation[ratingName] + 1
		}
	}
	if !found {
		s.err = errors.New("Error: unknown parametr '" + ratingName + "' for '" + s.orgType + "' organization")
		return
	}
	fmt.Println(s.Name, ratingName, "increased...")
}

func (s *Organization) decreaseRating(ratingName string) {
	found := false
	if _, ok := s.Market[ratingName]; ok {
		found = true
		if s.Market[ratingName] > 0 {
			s.Market[ratingName] = s.Market[ratingName] - 1
		}
	}
	if _, ok := s.Operation[ratingName]; ok {
		found = true
		if s.Operation[ratingName] > 1 {
			s.Operation[ratingName] = s.Operation[ratingName] - 1
		}
	}
	if !found {
		s.err = errors.New("Error: unknown parametr '" + ratingName + "'")
		return
	}
	fmt.Println(s.Name, ratingName, "decreased...")
}

//FullReport - Возвращает форматированную стену текста с описанием всех характеристик.
func (s *Organization) FullReport() string {
	netRating := 0
	report := "Organization Report: " + s.Name + "\n"
	report = report + "Markets" + "\n"

	markets := allMarkets(s.orgType)
	for i := range markets {
		marketRep := s.Rating(markets[i])

		report = report + markets[i] + ": " + strconv.Itoa(marketRep) + "\n"
		marketRating := s.Rating(markets[i])
		netRating = netRating + marketRating
	}
	report = report + "--------------------" + "\n"
	report = report + "Operations" + "\n"

	operations := allAttributes()
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

func (s *Organization) efficiencyTest() {

	fmt.Println("Run " + s.Name + " Efficiency test:")
	efficiency := s.Rating("Managment")
	hits, _, _, _ := sr3SimpleTest(efficiency, 4)
	var asset []string
	for marketName, _ := range s.Market {
		asset = append(asset, marketName)
	}
	for marketName, _ := range s.Operation {
		asset = append(asset, marketName)
	}
	asset = shuffleStringSlice(asset)
	for hits > 0 {
		s.increaseRating(asset[hits])
		hits--
	}
	fmt.Println("")
}

func (s *Organization) publicityTest() {
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

func (s *Organization) blackOpsTest() {
	target := pickTarget(s, AllOrganizations)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println(s.Name+" plotting Shadowrun against", market, "of", target)
	blOps := s.Rating("BlackOps")
	bonuses, _, _, _ := sr3SimpleTest(s.Rating(market), 5)
	flaws, _, _, _ := sr3SimpleTest(AllOrganizations[target].Rating(market), 5)
	var plannedDays float64
	intelHits, _, _, _ := sr3SimpleTest(s.Rating("Intelligence"), 4)
	randDays := randInt(0, 21)
	plannedDays = float64(randDays - intelHits + 1)
	if plannedDays < 0 {
		plannedDays = 1.0
	}
	var mcGuffinValue float64
	mcGuffinValue = float64(AllOrganizations[target].Rating(market) * 1000)
	run := PlotShadowrun(s, AllOrganizations[target], bonuses, flaws, plannedDays, mcGuffinValue)
	fmt.Println("Shadowrun parameters:")
	fmt.Println("Mission:", run.plot)
	fmt.Println("Proposed Payment:", run.payment)
	fmt.Println("Sponsor:", run.Sponsor.Name)
	fmt.Println("Optimal mission time:", plannedDays, "days")
	for i := bonuses; i > 0; i-- {
		bon := bonusList()
		r := randInt(1, len(bon)-1)
		fmt.Println("  " + bon[r])
	}
	for i := flaws; i > 0; i-- {
		flw := flawsList()
		r := randInt(1, len(flw)-1)
		fmt.Println("  " + flw[r])
	}
	targSecurity := AllOrganizations[target].Rating("Security")
	hits, outcome, _, gl := sr3SimpleTest(blOps, targSecurity)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
	}
	if hits > 1 {
		attr := false
		allAtr := allAttributes()
		for i := range allAtr {
			if allAtr[i] == market {
				attr = true
			}
		}
		if !attr {
			s.increaseRating(market)
		}
		AllOrganizations[target].decreaseRating(market)
	}
	fmt.Println("")
}

func (s *Organization) intelTest() {
	target := pickTarget(s, AllOrganizations)
	market := pickCommonRandomMarket(s.Name, target)
	fmt.Println(s.Name+" plotting Shadowrun against", market, "of", target)
	intel := s.Rating("Intelligence")
	targManagement := AllOrganizations[target].Rating("Counter-intelligence")
	hits, outcome, _, gl := sr3SimpleTest(intel, targManagement)
	fmt.Println("Run was a", outcome, gl)
	if hits == 0 {
		s.decreaseRating(market)
	}
	if hits > 1 {
		s.increaseRating(market)
		AllOrganizations[target].decreaseRating(market)
	}
	fmt.Println("")
}

func (s *Organization) naturalCycle() {
	s.efficiencyTest()
	s.publicityTest()
	s.blackOpsTest()
	s.intelTest()
}

func pickTarget(s *Organization, AllOrganizations map[string]*Organization) (target string) {
	var targetList []string
	for key := range AllOrganizations {
		if key == s.Name {
			continue
		}
		targetList = append(targetList, key)
	}
	r := randInt(1, len(targetList))
	return targetList[r-1]
}

func pickCommonRandomMarket(source, target string) string {
	sin1 := AllOrganizations[source]
	for key, val := range sin1.Market {
		if val != 0 {
			sin2 := AllOrganizations[target]
			val2, ok := sin2.Market[key]
			if val2 != 0 && ok {
				return key
			}
		}
	}
	for key, val := range sin1.Operation {
		if val != 0 {
			sin2 := AllOrganizations[target]
			val2, ok := sin2.Operation[key]
			if val2 != 0 && ok {
				return key
			}
		}
	}
	return ""
}
