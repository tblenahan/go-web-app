package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var tpl *template.Template

type DataContainer struct {
	strArr []string
	sages []sage
}

var sages = []string{"Gandhi", "Jesus", "Muhammad", "Buddha"}

type sage struct {
	Name string
	Motto string
}

func getSages() []sage {
	g := sage{
		Name: "Gandhi",
		Motto: "be the change",
	}

	j := sage{
		Name: "Jesus",
		Motto: "love all",
	}

	m := sage{
		Name: "Muhammad",
		Motto: "to overcome evil with good is good",
	}

	mlk := sage{
		Name: "Martin Luther King",
		Motto: "Hatred never ceases with hatred but is healed with love alone.",
	}

	return []sage{g, j, m, mlk}
}

var practiceData = DataContainer{
	strArr: sages,
	sages: getSages(),
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"fdateMDY": monthDayYear,
	"fdbl": double,
	"fsq": square,
	"fsqrt": sqRoot,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-20-1992")
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("./*.gohtml"))
}

func main() {
	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "make money!!!")

	//err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", practiceData.strArr)

	//err := tpl.ExecuteTemplate(os.Stdout,"tpl2.gohtml", practiceData.sages)

	//err := tpl.ExecuteTemplate(os.Stdout,"tpl3.gohtml", time.Now())

	err := tpl.ExecuteTemplate(os.Stdout,"tpl4.gohtml", 3)

	if err != nil {
		log.Fatalln(err)
	}
}
