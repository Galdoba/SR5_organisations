package main

import (
	"fmt"
)

type Shadowrun struct {
	plot      string
	payment   float64
	payMod    float64
	mrJohnson string
	Sponsor   *Organization
	Target    *Organization
}

func PlotShadowrun(sponsor, target *Organization, edges, flaws int, plannedDays, mcGuffinValue float64) *Shadowrun {
	sRun := Shadowrun{}
	sRun.Sponsor = sponsor
	sRun.Target = target
	sRun.plot = randomPlot()
	basePayment := plotBasePayment(sRun.plot, plannedDays, mcGuffinValue)
	maxOppRating := 1
	switch target.orgType {
	case "Corpotation":
		maxOppRating = 6
	case "Syndicate":
		maxOppRating = 4
	case "Gang":
		maxOppRating = 3
	default:
	}
	complications := randInt(0, maxOppRating)
	paymentMod := paymentMod(edges, flaws, complications)
	sRun.payment = basePayment * paymentMod

	return &sRun
}

func randomPlot() string {
	plotList := []string{
		"Assassination",
		"Blackmail",
		"Bodyguard",
		"Courier",
		"Smuggling",
		"Datasteal",
		"Distraction",
		"Destruction",
		"Encryption",
		"Decryption",
		"Enforcement",
		"Hoax",
		"Counterfeit",
		"Investigation",
		"Extraction",
		"Plant",
		"Retrieval of object",
		"Security",
		"Tailchaser",
		"War",
		"Wild things",
	}
	r := randInt(0, len(plotList)-1)
	return plotList[r]
}

func plotBasePayment(plot string, plannedDays float64, mcGuffinValue float64) float64 {
	paymentMap := make(map[string]float64)
	paymentMap["Assassination"] = 5000.0
	paymentMap["Blackmail"] = (400.0 * plannedDays) + 1000 //(Investigation * days) + Enforcement
	paymentMap["Bodyguard"] = 200.0 * plannedDays          // 200*15
	paymentMap["Courier"] = 1000.0
	paymentMap["Smuggling"] = 1500.0
	paymentMap["Datasteal"] = mcGuffinValue * 0.2
	paymentMap["Distraction"] = 1000.0
	paymentMap["Destruction"] = 5000.0
	paymentMap["Enforcement"] = 1000.0
	paymentMap["Encryption"] = mcGuffinValue
	paymentMap["Decryption"] = mcGuffinValue
	paymentMap["Enforcement"] = 1000.0
	paymentMap["Hoax"] = mcGuffinValue * 0.1
	paymentMap["Counterfeit"] = mcGuffinValue * 0.2
	paymentMap["Investigation"] = 300.0 * plannedDays
	paymentMap["Extraction"] = 20000.0
	paymentMap["Plant"] = mcGuffinValue * 0.2
	paymentMap["Retrieval of object"] = mcGuffinValue * 0.2
	paymentMap["Security"] = 300.0 * plannedDays
	paymentMap["Tailchaser"] = 400.0*plannedDays + 1000.0
	paymentMap["War"] = 10000.0 + (200.0 * plannedDays)
	paymentMap["Wild things"] = mcGuffinValue * 0.2
	fmt.Println("mcGuffinValue", mcGuffinValue)
	return paymentMap[plot]
}

func paymentMod(edges int, flaws int, complications int) float64 {
	var payMod float64
	payMod = 1.0
	payMod = payMod + (float64(edges) * -0.1)
	payMod = payMod + (float64(flaws) * 0.1)
	payMod = payMod + (float64(complications) * 1)
	fmt.Println("paymentMod", payMod)
	fmt.Println("Opposition Rating", complications)
	return payMod
}
