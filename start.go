package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/macroblock/imed/pkg/zlog/loglevel"
	"github.com/macroblock/imed/pkg/zlog/zlog"
	"github.com/macroblock/imed/pkg/zlog/zlogger"
)

var llog = zlog.Instance("main")

//AllOrganizations - сборник всех синдикатов(временно)
var AllOrganizations = make(map[string]*Organization)

func main() {
	wr := cType{}

	// newLogger := misc.NewSimpleLogger
	// if misc.IsTerminal() {
	// 	newLogger = misc.NewAnsiLogger
	// }
	// llog.Add(
	// 	misc.NewAnsiLogger(loglevel.Warning.OrLower(), ""),
	// 	misc.NewAnsiLogger(loglevel.Info.Only().Include(loglevel.Notice.Only()), "~x\n"),
	// )

	llog.Add(
		// zlogger.Build().
		// 	Writer(ansi.NewAnsiStdout()).
		// 	Styler(zlogger.AnsiStyler).
		// 	Format(zlogger.DefaultFormat).
		// 	LevelFilter(loglevel.Debug.OrLower()).
		// 	Done(),
		zlogger.Build(). //функциональная команда (начало строителя для логера)
					Writer(&wr).                        //Что выводит и куда
					LevelFilter(loglevel.Error.Only()). //условия вывода
					Format("||||||||||||~l~s~x~m~e\n").
					Done(), //функциональная команда (конец для Build())
		zlogger.Build(). //функциональная команда (начало строителя для логера)
					Writer(&wr).                       //Что выводит и куда
					LevelFilter(loglevel.Info.Only()). //условия вывода
					Format("~x\n").
					Done(), //функциональная команда (конец для Build())
	)
	// r := randInt(1, 5)
	// fmt.Println(r)
	// llog.Error(r > 0, "r > 2 str")
	// err := errors.New("потому что")
	// //err = nil
	// llog.Warning(err, "где ")

	seed := int64(time.Now().UnixNano())
	rand.Seed(seed)
	fmt.Println(sr3SimpleTest(1, 10))

	// AllOrganizations = make(map[string]*Organization)

	// AllOrganizations["Mafia"] = NewOrganization("Mafia", "Syndicate")
	// AllOrganizations["Yakuza"] = NewOrganization("Yakuza", "Syndicate")
	// AllOrganizations["Ares"] = NewOrganization("Ares", "Corporation")
	// AllOrganizations["Renraku"] = NewOrganization("Renraku", "Corporation")
	// AllOrganizations["Haloweeners"] = NewOrganization("Haloweeners", "Gang")
	// AllOrganizations["Killers"] = NewOrganization("Killers", "Gang")

	// for i := 0; i < 1; i++ {
	// 	fmt.Println(" ")
	// 	fmt.Println("Cycle", i+1)
	// 	for synName, sin := range AllOrganizations {
	// 		fmt.Println(synName + ":")
	// 		fmt.Println(" -----------")
	// 		sin.naturalCycle()
	// 		fmt.Println(sin.FullReport())
	// 	}
	// }

	if llog.State()&loglevel.Error.Only() != 0 {
		llog.Info("some err")

	}
	fl := InputFloat64("float1 =")
	fl2 := fl + 0.5
	fmt.Println(fl, fl2)
	for i := 0; i < 1; i++ {
		npc := NewNPC(InputString("Enter NPC Name: "))
		fmt.Println(npc)
	}

}

type cType struct{}

func (s *cType) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func bonusList() []string {
	bonuses := []string{
		"Mr. Johnson reveal Sponsor before contract is made",
		"Mr. Johnson reveal his target origin before contract is made",
		"Mr. Johnson provide full info on security of the target (Physical)",
		"Mr. Johnson provide full info on security of the target (Magic)",
		"Mr. Johnson provide full info on security of the target (Matrix)",
		"Mr. Johnson provide Magic Support",
		"Mr. Johnson provide Matrix Support",
		"Mr. Johnson provide Logistical Support (can buy stuff for the run with +6 avalability bonus)",
		"Mr. Johnson can provide Distraction",
		"Mr. Johnson give Contact after end of the run",
		"This is light hearted run (+1 karma reward)",
		"This is light hearted run (+1 karma reward)",
		"This is light hearted run (+1 karma reward)",
		"This Run has no restrictions on collateral damage",
		"This Run has no time restrictions",
		"Third party have benefit from runners success",
		"Mr. Johnson willing to pay a bonus for a successful job (payment multiplier +0.2)",
		"Mission present opportunities to get high value loot",
		"Mission present opportunities to rare gear",
	}
	return bonuses
}

func flawsList() []string {
	flaws := []string{
		"Mr. Johnson is unaware of the sponsor",
		"Mission target is unknown/not real",
		"Mr. Johnson provide inacurate information",
		"Mr. Johnson provide inacurate information",
		"Mr. Johnson provide inacurate information",
		"Mission takes place in high Background Area",
		"Mission takes place in high Matrix Noice Area",
		"Mission takes place in highly remote Area (usual transport is un available)",
		"Target is on high Alert (not for the runners)",
		"Mission involves interest of runner's Friend/Contact (reduce loyalty upon success of the mission)",
		"This is hard hearted run (-1 karma reward)",
		"This is hard hearted run (-1 karma reward)",
		"This is hard hearted run (-1 karma reward)",
		"This Run has high restrictions on collateral damage",
		"This Run has tight time restrictions",
		"Third party have benefit from runner's failure",
		"Mr. Johnson will not be able to pay agreed payment (payment multiplier -0.2)",
		"Mission present high risk for runner's gear",
		"Mission present high risk for runner's gear",
	}
	return flaws
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/jroimartin/gocui"
// )

// var counter int
// var ticker int

// func main() {
// 	g, err := gocui.NewGui(gocui.OutputNormal)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer g.Close()

// 	go func() {
// 		for {
// 			time.Sleep(1 * time.Second)
// 			g.Update(layout)
// 			ticker++

// 		}
// 	}()

// 	g.SetManagerFunc(layout)

// 	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
// 		log.Panicln(err)
// 	}

// 	if err := g.SetKeybinding("", rune(113), gocui.ModNone, increaseCounter); err != nil {
// 		log.Panicln(err)
// 	}

// 	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
// 		log.Panicln(err)
// 	}
// }

// func layout(g *gocui.Gui) error {
// 	maxX, maxY := g.Size()
// 	v, err := g.SetView("size", maxX/2-7, maxY/2, maxX/2+10, maxY/2+3)
// 	if err != nil && err != gocui.ErrUnknownView {
// 		return err
// 	}
// 	v.Clear()
// 	fmt.Fprintf(v, "%d, %d, %d, %d", maxX, maxY, ticker, counter)
// 	return nil
// }

// func quit(g *gocui.Gui, v *gocui.View) error {
// 	return gocui.ErrQuit
// }

// func increaseCounter(g *gocui.Gui, v *gocui.View) error {
// 	counter++
// 	if counter > 50 {
// 		return gocui.ErrQuit
// 	}

// 	return nil
// }
