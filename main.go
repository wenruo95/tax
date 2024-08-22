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
	tax1List := []float64{
		base * ElderlyCareInsuranceRate[1],
		base * MediacalInsuranceRate[1],
		base * MaternityInsuranceRate[1],
		base * UnemploymentInsuranceRate[1],
		base * WorkInjuryInsuranceRate[1],
		base * HousingProvidentFundRate[1],
	}
	tax1Listsum := tax1List[0] + tax1List[1] + tax1List[2] + tax1List[3] + tax1List[4] + tax1List[5]

	{
		var tax2sum float64
		for i := 0; i < 12; i++ {
			tax1 := tax1Listsum
			tax2 := CalculateTax((msalary-tax1Listsum)*float64(i+1)-tax2sum, TaxRates) - tax2sum
			tax2Return := base * (HousingProvidentFundRate[0] + HousingProvidentFundRate[1])
			left := msalary - tax1 - tax2
			fmt.Printf("[%2d] 月薪:%0.2f 五险一金:%0.2f 扣税:%0.2f 到手:%0.2f 税率:%0.2f%% 额外公积金:%0.2f\n",
				i+1, msalary, tax1, tax2, left, (msalary-left)*100/msalary, tax2Return)
			tax2sum = tax2sum + tax2
		}
	}

	{
		// year
		salary := msalary*(12+bonus) + stock
		tax1 := tax1Listsum * 12
		tax2 := CalculateTax(salary-tax1-TaxFreeSalary, TaxRates)
		tax2Return := base * (HousingProvidentFundRate[0] + HousingProvidentFundRate[1]) * 12
		left := salary - tax1 - tax2 + tax2Return
		fmt.Printf("月薪(%0.2f) + 奖金(%0.2f) + 股票(%0.2f) = 年包(%0.2f)\n",
			msalary, bonus*msalary, stock, salary)
		fmt.Printf("养老保险(%0.2f) + 医疗保险(%0.2f) + 生育保险(%0.2f) + 失业保险(%0.2f) + 工伤保险(%0.2f)"+
			" + 住房公积金(%0.2f) = 五险一金(%0.2f), 一年共计:%0.2f\n",
			tax1List[0], tax1List[1], tax1List[2], tax1List[3], tax1List[4], tax1List[5], tax1Listsum, tax1)
		fmt.Printf("年薪:%0.2f 五险一金:%0.2f 扣税:%0.2f 公积金返还:%0.2f 到手:%0.2f 税率:%0.2f%%\n",
			salary, tax1, tax2, tax2Return, left, (salary-left)*100/salary)
	}

	{
		// company
		ctaxList := []float64{
			base * ElderlyCareInsuranceRate[0],
			base * MediacalInsuranceRate[0],
			base * MaternityInsuranceRate[0],
			base * UnemploymentInsuranceRate[0],
			base * WorkInjuryInsuranceRate[0],
			base * HousingProvidentFundRate[0],
		}
		ctaxListsum := ctaxList[0] + ctaxList[1] + ctaxList[2] + ctaxList[3] + ctaxList[4] + ctaxList[5]
		fmt.Printf("[公司] 养老保险(%0.2f) + 医疗保险(%0.2f) + 生育保险(%0.2f) + 失业保险(%0.2f) + 工伤保险(%0.2f)"+
			" + 住房公积金(%0.2f) = 五险一金(%0.2f), 一年共计:%0.2f\n",
			ctaxList[0], ctaxList[1], ctaxList[2], ctaxList[3], ctaxList[4], ctaxList[5], ctaxListsum, ctaxListsum*12)
	}
}
