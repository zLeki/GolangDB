package core

import (
	"GopherDB/build"
	"GopherDB/input"
	DropDown2 "GopherDB/select"
	"GopherDB/table"
	"GopherDB/types"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func PrevPage() {
	Clear()
	var i = -1
	for {
		if types.Page < 0 {
			types.Page = 0
			i = -1 // RESET-
		} else if types.Page >= 3 {
			types.Page--
		}
		if types.Page == 2 && i != 2 {
			clear[runtime.GOOS]()
			var questions []input.Questions
			questions = append(questions, input.Questions{
				"Which value are you changing?",
				32,
				false,
			})
			questions = append(questions, input.Questions{
				"What is the new value?",
				32,
				false,
			})
			input.Input(questions)
			i = 2
		} else if types.Page == 0 && i != 0 {
			clear[runtime.GOOS]()
			DropDown2.DropDown(DropDown2.GetTableNames(types.Db), "Tables")
			i = 0
		} else if types.Page == 1 && i != 1 {
			clear[runtime.GOOS]()
			column, rows := build.BuildTable(types.Db)
			table.Table(column, rows)
			i = 1
		} else {
			types.Page = 0
			i = -1
		}
		time.Sleep(time.Millisecond * 100)
	}

}

var clear map[string]func()

func Clear() {
	clear = make(map[string]func())
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Mac example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
