package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

var counter int
var ticker int
var tickerGo bool
var appErr error
var runStart time.Time

func main() {
	runStart = time.Now()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, actionQuit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", rune(113) /*q*/, gocui.ModNone, actionIncreaseCounter); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", rune(109) /*m*/, gocui.ModNone, actionToggleTicker); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.MouseMiddle, gocui.ModNone, actionButtonClick); err != nil {
		log.Panicln(err)
	}

	counter = 1

	go func() {
		for {
			time.Sleep(40 * time.Millisecond)
			g.Update(layout)
			if tickerGo {
				ticker = ticker + counter
			}

		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

//Создает и отрисовывает все окна - к этому моменту программа должна иметь
//представление что где и в каком окне должно быть.
//Запускается при каждом обновлении экрана
//TODO: прощупать стоит ли хранить содержимое окна где-либо вне его.
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v1, v1Err := newPanelInfo(g, "Size", 0, 0, maxX/2, maxY-3)
	if v1Err != nil && v1Err != gocui.ErrUnknownView {
		return v1Err
	}
	v2, err := newPanelInfo(g, "Info", maxX/2+1, 0, maxX/2-5, maxY-4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	fillPanel(v1)
	fillPanel(v2)
	return nil
}

func newPanelInfo(g *gocui.Gui, panelName string, pX, pY, pW, pH int) (*gocui.View, error) {

	v, err := g.SetView(panelName, pX, pY, pX+pW, pY+pH)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}
	v.Title = panelName
	if panelName == "Info" {
		v.Highlight = true

	}
	return v, nil
}

func fillPanel(v *gocui.View) {
	switch v.Title {
	case "Size":
		v.Clear()
		t := time.Now().Format("2006-Jan-02 15:04:05")
		ts := runStart.Format("2006-Jan-02 15:04:05")
		fmt.Fprintf(v, "Current Real Time: %s \n", t)
		fmt.Fprintf(v, "RunStart: %s\n", ts)
		s := time.Since(runStart).Round(time.Millisecond)
		pureSeconds := float64(s/time.Millisecond) + 567
		fmt.Fprintf(v, "Program working: %s\n Sec: %d\n", s, pureSeconds/1000)
		fmt.Fprintf(v, "%d, %d\n", ticker, counter)
		fmt.Fprintf(v, "rume 'm' = %d", string(rune(109)))
		// if tickerGo {
		// 	fmt.Fprintf(v, "tickerGo is active")
		// }
	case "Info":
		v.Clear()
		if tickerGo {
			fmt.Fprintf(v, "tickerGo is active")
		}
	}

}
