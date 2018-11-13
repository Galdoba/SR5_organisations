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

//AllSyndicates - сборник всех синдикатов(временно)
var AllSyndicates = make(map[string]*Syndicate)

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

	AllSyndicates = make(map[string]*Syndicate)

	AllSyndicates["Mafia"] = NewSyndicate("Mafia")
	AllSyndicates["Yakuza"] = NewSyndicate("Yakuza")
	AllSyndicates["Triada"] = NewSyndicate("Triada")

	for i := 0; i < 10; i++ {
		fmt.Println(" ")
		fmt.Println("Cycle", i+1)
		for synName, sin := range AllSyndicates {
			fmt.Println(synName + ":")
			fmt.Println(" -----------")
			sin.naturalCycle()
			//fmt.Println(sin.FullReport())
		}
	}

	if llog.State()&loglevel.Error.Only() != 0 {
		llog.Info("some err")

	}

}

type cType struct{}

func (s *cType) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
