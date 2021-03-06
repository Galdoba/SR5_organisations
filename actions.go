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

func actionChangeColor(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View(v.Name())
	bg := view.BgColor
	bg++
	if bg > 7 {
		bg = 0
	}
	view.BgColor = bg
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

func actionMoveRight(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX+2, voY)

	return err
}

func actionMoveLeft(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX-2, voY)

	return err
}
