package main

import (
	"flag"
	"fmt"
	"math"
)

func CalculateTax(salary float64, rates []*TaxRateData) float64 {
	var sum float64
	for i := 0; i < len(rates) && salary > rates[i].L; i++ {
		rate := rates[i]
		tax := math.Min(salary-rate.L, rate.R-rate.L) * rate.Rate
		sum = sum + tax
		//fmt.Printf("[L:%v R:%v Rate:%0.2f] tax:%v sum:%v current_rate:%0.2f%%\n",
		//	rate.L, rate.R, rate.Rate, tax, sum, sum*100/salary)
	}
	return sum
}

func main() {
	var msalary, bonus, stock float64
	flag.Float64Var(&msalary, "salary", 0, "salary-月薪（单位：元）, --salary 100000")
	flag.Float64Var(&bonus, "bonus", 0, "bonus-年终奖（单位：月）, --bonus 3.5")
	flag.Float64Var(&stock, "stock", 0, "stock-股票（单位：元）, --stock 1000000")
	flag.Parse()
	if msalary == 0 {
		flag.Usage()
		return
	}

	base := math.Min(msalary, 3*SocialAverateSalary)
	mtax1 := []float64{
		base * ElderlyCareInsuranceRate[1],
		base * MediacalInsuranceRate[1],
		base * MaternityInsuranceRate[1],
		base * UnemploymentInsuranceRate[1],
		base * WorkInjuryInsuranceRate[1],
		base * HousingProvidentFundRate[1],
	}
	var mtax1sum float64
	for _, v := range mtax1 {
		mtax1sum += v
	}

	salary := msalary*(12+bonus) + stock
	tax1 := mtax1sum * 12
	tax2 := CalculateTax(salary-tax1-TaxFreeSalary, TaxRates)
	tax2Return := base * (HousingProvidentFundRate[0] + HousingProvidentFundRate[1]) * 12
	left := salary - tax1 - tax2 + tax2Return

	fmt.Printf("月薪(%0.2f) + 奖金(%0.2f) + 股票(%0.2f) = 年包(%0.2f)\n",
		msalary, bonus*msalary, stock, salary)
	fmt.Printf("养老保险(%0.2f) + 医疗保险(%0.2f) + 生育保险(%0.2f) + 失业保险(%0.2f) + 工伤保险(%0.2f)"+
		" 住房公积金(%0.2f) = 五险一金(%0.2f), 一年共计:%0.2f\n",
		mtax1[0], mtax1[1], mtax1[2], mtax1[3], mtax1[4], mtax1[5], mtax1sum, tax1)
	fmt.Printf("年薪:%0.2f 五险一金:%0.2f 扣税:%0.2f 公积金返还:%0.2f 到手:%0.2f 税率:%0.2f%%\n",
		salary, tax1, tax2, tax2Return, left, (1-left/salary)*100)
}
