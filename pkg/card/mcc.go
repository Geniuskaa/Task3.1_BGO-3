package card

func TranslateMCC(code string) string {
	// представил, что я смог написать грамотный Reader и прочёл mcc из файла
	mcc := map[string]string{
		"6550": "Рестораны",
		"6501": "Супермаркеты",
		"5000": "Пополнение",
	}

	for s, s2 := range mcc {
		if(s == code) { // Вопрос вам, как Go программисту со стажем, есть ли в Go что-то типо equal() в Java?
			return s2  // Или сравнивать через == сойдёт?
		}
	}
	return "zero"
}