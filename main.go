package main

import (
	"flag"
	"fmt"
)

type taxScale struct {
	min      float64
	rate     float64
	discount float64
}

var scales = []taxScale{
	{0, 0, 0},
	{2501, 10, 250},
	{4167, 20, 666.67},
	{5001, 30, 1166.67},
	{6667, 34, 1433.33},
	{15001, 38, 2033.33},
}

var cnss_rate float64 = 4.52
var amo_rate float64 = 1.85
var grace_rate float64 = 20

func main() {
	salaire_brut := getSalaire()

	if salaire_brut < 0 {
		println("[ERR] salaire brut is required")
		flag.Usage()
		return
	}

	amo := salaire_brut * amo_rate / 100
	cnss := salaire_brut * cnss_rate / 100
	rni := salaire_brut*(100-grace_rate)/100 - amo - cnss
	tax_rate, discount := getTaxes(rni)
	ir := rni*tax_rate/100 - discount
	salaire_net := salaire_brut - ir - cnss - amo

	printfln("- Salaire brut : %.2f MAD", salaire_brut)
	printfln("\t- Cnss : %.2f MAD", cnss)
	printfln("\t- Amo : %.2f MAD", amo)
	printfln("\t- RNI : %2.f MAD", rni)
	printfln("\t- IR : %2.f MAD", ir)
	println("==============================")
	printfln("[*] Salaire net : %2.f MAD", salaire_net)
}

func printfln(format string, a ...any) (int, error) {
	return fmt.Printf(format+"\n", a...)
}

func getTaxes(rni float64) (float64, float64) {
	var tax_rate float64 = -1
	var discount float64 = -1

	for i := range scales {
		if tax_rate < 0 || rni > scales[i].min {
			tax_rate = scales[i].rate
			discount = scales[i].discount
		}
	}

	return tax_rate, discount
}

func getSalaire() float64 {
	sal := flag.Float64("s", -1, "Salaire brut")
	flag.Parse()

	return *sal
}
