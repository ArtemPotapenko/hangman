package russian

func IsRussianLetter(r rune) bool {
	return (r >= 'А' && r <= 'Я') ||
		(r >= 'а' && r <= 'я') ||
		r == 'Ё' || r == 'ё'
}
