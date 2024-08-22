package main

// 前开后闭
type TaxRateData struct {
	L, R, Rate float64
}

var TaxRates = []*TaxRateData{
	{0, 36000, P100(3)},
	{36000, 144000, P100(10)},
	{144000, 300000, P100(20)},
	{300000, 420000, P100(25)},
	{420000, 660000, P100(30)},
	{660000, 960000, P100(35)},
	{960000, 10000 * 1e8, P100(45)},
}

func P100(x float64) float64 {
	return x / float64(100)
}

var (
	SocialAverateSalary float64 = 14553 // 深圳3倍社平

	ElderlyCareInsuranceRate = []float64{P100(15), P100(8)} // 养老保险
	//ElderlyCareInsuranceRate  = []float64{P100(13 + 1), P100(5)}  // 养老保险
	MediacalInsuranceRate     = []float64{P100(6 + 0.2), P100(2)} // 医疗保险
	MaternityInsuranceRate    = []float64{P100(0.5), P100(0)}     // 生育保险
	UnemploymentInsuranceRate = []float64{P100(1), P100(0.5)}     // 失业保险
	WorkInjuryInsuranceRate   = []float64{P100(0.14), P100(0)}    // 工伤保险
	HousingProvidentFundRate  = []float64{P100(12), P100(12)}     // 住房公积金
	//HousingProvidentFundRate = []float64{P100(10), P100(10)} // 住房公积金

	HousingProvidentFundReturnRate = 100 / 100 // 住房公积金返回比例

	//TaxFreeSalary float64 = 60000 // 专项附加扣除
	TaxFreeSalary float64 = 0 // 专项附加扣除
)
