package main

import "github.com/jroimartin/gocui"

func actionQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//////

func actionIncreaseCounter(g *gocui.Gui, v *gocui.View) error {
	counter++
	if counter > 50 {
		return gocui.ErrQuit
	}
	//toggleTicker()
	return nil
}

func actionToggleTicker(g *gocui.Gui, v *gocui.View) error {
	toggleTicker()
	return nil
}

func toggleTicker() {
	if tickerGo {
		tickerGo = false
	} else {
		tickerGo = true
	}

}

//////

func actionButtonClick(g *gocui.Gui, v *gocui.View) error {
	err := executeClick(v)
	return err
}

func executeClick(v *gocui.View) error {
	switch v.Title {
	case "Info":

		toggleTicker()
		fillPanel(v)
	default:

	}

	return nil
}
