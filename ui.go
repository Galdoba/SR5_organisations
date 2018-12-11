package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	viewArr    = []string{"User Input", "Program Output", "v3", "v4"}
	active     = 0
	userString = ""
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("Program Output")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func printInView(g *gocui.Gui, v string, text string) error {
	view, callErr := g.View(v)
	if callErr != nil {
		return callErr
	}
	fmt.Fprint(view, text)
	return nil
}

func userInput(g *gocui.Gui, v *gocui.View) error {
	inputView, _ := g.View("User Input")
	printInView(g, "v3", "test")

	cache := inputView.ViewBuffer()

	bt := []byte(cache)
	if len(bt) > 0 {
		bt = bt[:len(bt)-1]
	}
	if len(bt) > 0 {
		cache = string(bt)
	} else {
		cache = ""
	}
	log, err := g.View("Program Output")
	if err != nil {
		return err
	}
	if len(cache) > 0 || cache != "" {
		fmt.Fprintln(log, "User String:", cache)
		userString = cache
	}
	inputView.Clear()
	inputView.SetCursor(0, 0)
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("User Input", maxX/2, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "User Input (editable)"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "User Input"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("Program Output", maxX/2-1, 0, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Program Output"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v3", 0, 0, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "v3"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "v4 (editable)"
		v.Editable = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main0() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("User Input", gocui.KeyEnter, gocui.ModNone, userInput); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
