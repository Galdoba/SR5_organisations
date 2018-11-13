package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	ansi "github.com/k0kubun/go-ansi"
	"github.com/macroblock/imed/pkg/zlog/loglevel"
	"github.com/macroblock/imed/pkg/zlog/zlog"
	"github.com/macroblock/imed/pkg/zlog/zlogger"
)

var llog = zlog.Instance("main")
var Syndicates map[string]*Syndicate

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
		zlogger.Build().
			Writer(ansi.NewAnsiStdout()).
			Styler(zlogger.AnsiStyler).
			Format(zlogger.DefaultFormat).
			LevelFilter(loglevel.Debug.OrLower()).
			Done(),
		zlogger.Build(). //функциональная команда (начало строителя для логера)
					Writer(&wr).                        //Что выводит и куда
					LevelFilter(loglevel.Error.Only()). //условия вывода
					Format("||||||||||||~l~s~x~m~e\n").
					Done(), //функциональная команда (конец для Build())
	)
	r := randInt(1, 5)
	fmt.Println(r)
	llog.Error(r > 0, "r > 2 str")
	err := errors.New("потому что")
	//err = nil
	llog.Warning(err, "где ")

	seed := int64(time.Now().UnixNano())
	rand.Seed(seed)
	fmt.Println(sr3SimpleTest(1, 10))

	//eff = 6 pub = 8 BO = 6 int = 6
	fmt.Println("Effic")
	fmt.Println(sr3SimpleTest(6, 4))
	fmt.Println("Publ")
	fmt.Println(sr3SimpleTest(8, 8))
	fmt.Println("Black Ops")
	fmt.Println(sr3SimpleTest(6, 7))
	fmt.Println("Intel")
	fmt.Println(sr3SimpleTest(6, 5))

	sin := NewSyndicate("Mafia")
	Syndicates["Mafia"] = sin
	fmt.Println("-----------")
	fmt.Println("FullReport:")
	for i := 0; i < 3; i++ {
		for synName, sin := range Syndicates {
			fmt.Println("Go:", synName)
			fmt.Println(sin.FullReport())
			sin.NaturalCycle()
		}
	}

	if llog.State()&loglevel.Error.Only() != 0 {
		llog.Info("some err")

	}

}

type cType struct{}

func (s *cType) Write(p []byte) (n int, err error) {
	fmt.Println("Custom:", string(p))
	return len(p), nil
}
