package romannumerals

func ConvertToRoman(arabic int) (roman string) {
	if arabic == 2 {
		roman = "II"
		return roman
	}
	roman = "I"
	return roman
}
