package lcd

func Lcd(number string) string {
	mapNumber := map[string]string{
		"1": "|\n|",
		"2": " _\n _|\n|_",
		"4": "\n|_|\n  |",
		"7": "_\n |\n |",
		"0": " _\n| |\n|_|",
	}
	var lcdNumber string
	for _, digit := range number {
		lcdNumber += mapNumber[string(digit)]
	}
	return lcdNumber
}
