package lcd

func Lcd(number string) string {
	high := map[string]string{
		"1": "",
		"2": " _ ",
		"3": " _ ",
		"4": "   ",
		"5": " _ ",
		"6": " _ ",
		"7": "_ ",
		"8": " _ ",
		"9": " _ ",
		"0": " _ ",
	}
	middle := map[string]string{
		"1": "|",
		"2": " _|",
		"3": " |",
		"4": "|_|",
		"5": "|_",
		"6": "|_",
		"7": " |",
		"8": "|_|",
		"9": "|_|",
		"0": "| |",
	}
	low := map[string]string{
		"1": "|",
		"2": "|_",
		"3": "_|",
		"4": "  |",
		"5": " _|",
		"6": "|_|",
		"7": " |",
		"8": "|_|",
		"9": " _|",
		"0": "|_|",
	}
	var lcdHigh, lcdMiddle, lcdLow string
	for _, digit := range number {
		lcdHigh += high[string(digit)]
		lcdMiddle += middle[string(digit)]
		lcdLow += low[string(digit)]
	}
	return lcdHigh + "\n" + lcdMiddle + "\n" + lcdLow
}
