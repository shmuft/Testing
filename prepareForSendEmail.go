package main

type UserForEmail struct {
	UserName string
	Ank      anketaEmail
	T1       test1Email
	T2       test1Email
	T3       test1Email
	level    byte
}
type test1Email struct {
	Headers []string
	Values  []int
}
type anketaEmail struct {
	Name            string
	Sex             string
	Age             int
	Education       string
	Experience      int
	HowManyPeople   int
	HowLongWorkProf int
	HowLongWork     int
}

func NewUserForEmail(u *user) UserForEmail {
	userEmail := UserForEmail{}
	userEmail.PrepareUserForEmail(u)
	return userEmail
}

func (userEmail *UserForEmail) PrepareUserForEmail(u *user) {
	userEmail.UserName = u.UserName
	userEmail.PrepareAnketa(u)
	userEmail.PrepareTest1(u)
	userEmail.PrepareTest2(u)
	userEmail.PrepareTest3(u)
}

func (userEmail *UserForEmail) PrepareAnketa(u *user) {
	ank := anketaEmail{
		Age:             u.Ank.Age,
		Experience:      u.Ank.Experience,
		HowLongWork:     u.Ank.Experience,
		HowLongWorkProf: u.Ank.HowLongWorkProf,
		HowManyPeople:   u.Ank.HowManyPeople,
		Name:            u.Ank.Name,
	}
	var educationString string

	switch u.Ank.Sex {
	case "M":
		ank.Sex = "Мужской"
	case "F":
		ank.Sex = "Женский"
	}

	switch u.Ank.Education {
	case 1:
		educationString = "Неполное среднее"
	case 2:
		educationString = "Среднее"
	case 3:
		educationString = "Специальное (Училище или Техникум)"
	case 4:
		educationString = "Неполное высшее"
	case 5:
		educationString = "Высшее"
	case 6:
		educationString = "Несколько высших образовании (2 и более)"
	case 7:
		educationString = "Ученая степень"
	}
	ank.Education = educationString
	userEmail.Ank = ank
}

func (userEmail *UserForEmail) PrepareTest3(u *user) {
	sumArrayTest3 := make([]int, 3, 3)
	radioValues := make([]int, 3, 3)
	sumArrayTest3[0] = 0
	sumArrayTest3[1] = 0
	sumArrayTest3[2] = 0

	for i := 0; i < len(u.T3); i++ {
		radioValues[0] = 0
		radioValues[1] = 0
		radioValues[2] = 0

		radioValues[u.T3[i].Answers[0]] = 1
		switch i {
		case 0:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[1] += radioValues[1]
			sumArrayTest3[2] += radioValues[2]
		case 1:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[1] += radioValues[1]
			sumArrayTest3[2] += radioValues[2]
		case 2:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 3:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 4:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 5:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 6:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 7:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 8:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 9:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[1] += radioValues[1]
			sumArrayTest3[0] += radioValues[2]
		case 10:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[1] += radioValues[1]
			sumArrayTest3[0] += radioValues[2]
		case 11:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 12:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 13:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 14:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 15:
			sumArrayTest3[2] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		case 16:
			sumArrayTest3[1] += radioValues[0]
			sumArrayTest3[0] += radioValues[1]
			sumArrayTest3[2] += radioValues[2]
		case 17:
			sumArrayTest3[0] += radioValues[0]
			sumArrayTest3[2] += radioValues[1]
			sumArrayTest3[1] += radioValues[2]
		}
	}

	headers := []string{
		"Демократический",
		"Либеральный (попустительский)",
		"Автократический",
	}
	userEmail.T3 = test1Email{Headers:headers, Values:sumArrayTest3}

}

func (userEmail *UserForEmail) PrepareTest2(u *user) {
	sumArrayTest2 := make([]int, 8, 8)
	for i := 0; i < 8; i++ {
		sumArrayTest2[i] = 0
	}

	for i := 0; i < len(u.T2); i++ {
		switch i {
		case 0:
			sumArrayTest2[4] += u.T2[i].Answers[0]
			sumArrayTest2[6] += u.T2[i].Answers[1]
			sumArrayTest2[3] += u.T2[i].Answers[2]
			sumArrayTest2[1] += u.T2[i].Answers[3]
			sumArrayTest2[7] += u.T2[i].Answers[4]
			sumArrayTest2[2] += u.T2[i].Answers[5]
			sumArrayTest2[0] += u.T2[i].Answers[6]
			sumArrayTest2[5] += u.T2[i].Answers[7]
		case 1:
			sumArrayTest2[0] += u.T2[i].Answers[0]
			sumArrayTest2[1] += u.T2[i].Answers[1]
			sumArrayTest2[4] += u.T2[i].Answers[2]
			sumArrayTest2[5] += u.T2[i].Answers[3]
			sumArrayTest2[2] += u.T2[i].Answers[4]
			sumArrayTest2[6] += u.T2[i].Answers[5]
			sumArrayTest2[3] += u.T2[i].Answers[6]
			sumArrayTest2[7] += u.T2[i].Answers[7]
		case 2:
			sumArrayTest2[1] += u.T2[i].Answers[0]
			sumArrayTest2[7] += u.T2[i].Answers[1]
			sumArrayTest2[2] += u.T2[i].Answers[2]
			sumArrayTest2[3] += u.T2[i].Answers[3]
			sumArrayTest2[6] += u.T2[i].Answers[4]
			sumArrayTest2[4] += u.T2[i].Answers[5]
			sumArrayTest2[5] += u.T2[i].Answers[6]
			sumArrayTest2[0] += u.T2[i].Answers[7]
		case 3:
			sumArrayTest2[6] += u.T2[i].Answers[0]
			sumArrayTest2[2] += u.T2[i].Answers[1]
			sumArrayTest2[5] += u.T2[i].Answers[2]
			sumArrayTest2[0] += u.T2[i].Answers[3]
			sumArrayTest2[3] += u.T2[i].Answers[4]
			sumArrayTest2[7] += u.T2[i].Answers[5]
			sumArrayTest2[4] += u.T2[i].Answers[6]
			sumArrayTest2[1] += u.T2[i].Answers[7]
		case 4:
			sumArrayTest2[5] += u.T2[i].Answers[0]
			sumArrayTest2[0] += u.T2[i].Answers[1]
			sumArrayTest2[6] += u.T2[i].Answers[2]
			sumArrayTest2[2] += u.T2[i].Answers[3]
			sumArrayTest2[4] += u.T2[i].Answers[4]
			sumArrayTest2[1] += u.T2[i].Answers[5]
			sumArrayTest2[7] += u.T2[i].Answers[6]
			sumArrayTest2[3] += u.T2[i].Answers[7]
		case 5:
			sumArrayTest2[3] += u.T2[i].Answers[0]
			sumArrayTest2[6] += u.T2[i].Answers[1]
			sumArrayTest2[1] += u.T2[i].Answers[2]
			sumArrayTest2[7] += u.T2[i].Answers[3]
			sumArrayTest2[5] += u.T2[i].Answers[4]
			sumArrayTest2[0] += u.T2[i].Answers[5]
			sumArrayTest2[2] += u.T2[i].Answers[6]
			sumArrayTest2[4] += u.T2[i].Answers[7]
		case 6:
			sumArrayTest2[2] += u.T2[i].Answers[0]
			sumArrayTest2[5] += u.T2[i].Answers[1]
			sumArrayTest2[7] += u.T2[i].Answers[2]
			sumArrayTest2[4] += u.T2[i].Answers[3]
			sumArrayTest2[0] += u.T2[i].Answers[4]
			sumArrayTest2[3] += u.T2[i].Answers[5]
			sumArrayTest2[1] += u.T2[i].Answers[6]
			sumArrayTest2[6] += u.T2[i].Answers[7]
		}
	}

	headers := []string{
		"Реализатор",
		"Координатор",
		"Творец",
		"Генератор идей",
		"Исследователь",
		"Эксперт",
		"Дипломат",
		"Исполнитель",
	}
	userEmail.T2 = test1Email{Headers:headers, Values:sumArrayTest2}
}

func (userEmail *UserForEmail) PrepareTest1(u *user) {
	sumArrayTest1 := make([]int, 12, 12)
	for i := 0; i < 12; i++ {
		sumArrayTest1[i] = 0
	}
	for i := 0; i < len(u.T1); i++ {
		switch i {
		case 0:
			sumArrayTest1[0] += u.T1[i].Answers[0]
			sumArrayTest1[4] += u.T1[i].Answers[1]
			sumArrayTest1[7] += u.T1[i].Answers[2]
			sumArrayTest1[10] += u.T1[i].Answers[3]
		case 1:
			sumArrayTest1[1] += u.T1[i].Answers[3]
			sumArrayTest1[2] += u.T1[i].Answers[0]
			sumArrayTest1[5] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
		case 2:
			sumArrayTest1[8] += u.T1[i].Answers[0]
			sumArrayTest1[3] += u.T1[i].Answers[1]
			sumArrayTest1[2] += u.T1[i].Answers[2]
			sumArrayTest1[4] += u.T1[i].Answers[3]
		case 3:
			sumArrayTest1[3] += u.T1[i].Answers[0]
			sumArrayTest1[5] += u.T1[i].Answers[1]
			sumArrayTest1[2] += u.T1[i].Answers[2]
			sumArrayTest1[8] += u.T1[i].Answers[3]
		case 4:
			sumArrayTest1[2] += u.T1[i].Answers[0]
			sumArrayTest1[1] += u.T1[i].Answers[1]
			sumArrayTest1[0] += u.T1[i].Answers[2]
			sumArrayTest1[10] += u.T1[i].Answers[3]
		case 5:
			sumArrayTest1[1] += u.T1[i].Answers[0]
			sumArrayTest1[0] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
			sumArrayTest1[5] += u.T1[i].Answers[3]
		case 6:
			sumArrayTest1[2] += u.T1[i].Answers[0]
			sumArrayTest1[4] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
			sumArrayTest1[5] += u.T1[i].Answers[3]
		case 7:
			sumArrayTest1[8] += u.T1[i].Answers[0]
			sumArrayTest1[10] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
			sumArrayTest1[9] += u.T1[i].Answers[3]
		case 8:
			sumArrayTest1[5] += u.T1[i].Answers[0]
			sumArrayTest1[10] += u.T1[i].Answers[1]
			sumArrayTest1[8] += u.T1[i].Answers[2]
			sumArrayTest1[7] += u.T1[i].Answers[3]
		case 9:
			sumArrayTest1[3] += u.T1[i].Answers[0]
			sumArrayTest1[0] += u.T1[i].Answers[1]
			sumArrayTest1[4] += u.T1[i].Answers[2]
			sumArrayTest1[10] += u.T1[i].Answers[3]
		case 10:
			sumArrayTest1[1] += u.T1[i].Answers[0]
			sumArrayTest1[2] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
			sumArrayTest1[5] += u.T1[i].Answers[3]
		case 11:
			sumArrayTest1[6] += u.T1[i].Answers[0]
			sumArrayTest1[10] += u.T1[i].Answers[1]
			sumArrayTest1[5] += u.T1[i].Answers[2]
			sumArrayTest1[1] += u.T1[i].Answers[3]
		case 12:
			sumArrayTest1[4] += u.T1[i].Answers[0]
			sumArrayTest1[6] += u.T1[i].Answers[1]
			sumArrayTest1[9] += u.T1[i].Answers[2]
			sumArrayTest1[1] += u.T1[i].Answers[3]
		case 13:
			sumArrayTest1[10] += u.T1[i].Answers[0]
			sumArrayTest1[9] += u.T1[i].Answers[1]
			sumArrayTest1[3] += u.T1[i].Answers[2]
			sumArrayTest1[0] += u.T1[i].Answers[3]
		case 14:
			sumArrayTest1[7] += u.T1[i].Answers[0]
			sumArrayTest1[8] += u.T1[i].Answers[1]
			sumArrayTest1[5] += u.T1[i].Answers[2]
			sumArrayTest1[2] += u.T1[i].Answers[3]
		case 15:
			sumArrayTest1[6] += u.T1[i].Answers[0]
			sumArrayTest1[2] += u.T1[i].Answers[1]
			sumArrayTest1[0] += u.T1[i].Answers[2]
			sumArrayTest1[7] += u.T1[i].Answers[3]
		case 16:
			sumArrayTest1[2] += u.T1[i].Answers[0]
			sumArrayTest1[4] += u.T1[i].Answers[1]
			sumArrayTest1[6] += u.T1[i].Answers[2]
			sumArrayTest1[8] += u.T1[i].Answers[3]
		case 17:
			sumArrayTest1[9] += u.T1[i].Answers[0]
			sumArrayTest1[3] += u.T1[i].Answers[1]
			sumArrayTest1[10] += u.T1[i].Answers[2]
			sumArrayTest1[7] += u.T1[i].Answers[3]
		case 18:
			sumArrayTest1[5] += u.T1[i].Answers[0]
			sumArrayTest1[7] += u.T1[i].Answers[1]
			sumArrayTest1[0] += u.T1[i].Answers[2]
			sumArrayTest1[9] += u.T1[i].Answers[3]
		case 19:
			sumArrayTest1[8] += u.T1[i].Answers[0]
			sumArrayTest1[7] += u.T1[i].Answers[1]
			sumArrayTest1[0] += u.T1[i].Answers[2]
			sumArrayTest1[1] += u.T1[i].Answers[3]
		case 20:
			sumArrayTest1[2] += u.T1[i].Answers[0]
			sumArrayTest1[9] += u.T1[i].Answers[1]
			sumArrayTest1[3] += u.T1[i].Answers[2]
			sumArrayTest1[11] += u.T1[i].Answers[3]
		case 21:
			sumArrayTest1[0] += u.T1[i].Answers[0]
			sumArrayTest1[1] += u.T1[i].Answers[1]
			sumArrayTest1[7] += u.T1[i].Answers[2]
			sumArrayTest1[9] += u.T1[i].Answers[3]
		case 22:
			sumArrayTest1[3] += u.T1[i].Answers[0]
			sumArrayTest1[6] += u.T1[i].Answers[1]
			sumArrayTest1[7] += u.T1[i].Answers[2]
			sumArrayTest1[0] += u.T1[i].Answers[3]
		case 23:
			sumArrayTest1[1] += u.T1[i].Answers[0]
			sumArrayTest1[7] += u.T1[i].Answers[1]
			sumArrayTest1[6] += u.T1[i].Answers[2]
			sumArrayTest1[9] += u.T1[i].Answers[3]
		case 24:
			sumArrayTest1[1] += u.T1[i].Answers[0]
			sumArrayTest1[10] += u.T1[i].Answers[1]
			sumArrayTest1[8] += u.T1[i].Answers[2]
			sumArrayTest1[3] += u.T1[i].Answers[3]
		case 25:
			sumArrayTest1[1] += u.T1[i].Answers[0]
			sumArrayTest1[3] += u.T1[i].Answers[1]
			sumArrayTest1[11] += u.T1[i].Answers[2]
			sumArrayTest1[8] += u.T1[i].Answers[3]
		case 26:
			sumArrayTest1[5] += u.T1[i].Answers[0]
			sumArrayTest1[10] += u.T1[i].Answers[1]
			sumArrayTest1[6] += u.T1[i].Answers[2]
			sumArrayTest1[4] += u.T1[i].Answers[3]
		case 27:
			sumArrayTest1[11] += u.T1[i].Answers[0]
			sumArrayTest1[8] += u.T1[i].Answers[1]
			sumArrayTest1[4] += u.T1[i].Answers[2]
			sumArrayTest1[6] += u.T1[i].Answers[3]
		case 28:
			sumArrayTest1[11] += u.T1[i].Answers[0]
			sumArrayTest1[4] += u.T1[i].Answers[1]
			sumArrayTest1[3] += u.T1[i].Answers[2]
			sumArrayTest1[7] += u.T1[i].Answers[3]
		case 29:
			sumArrayTest1[3] += u.T1[i].Answers[0]
			sumArrayTest1[5] += u.T1[i].Answers[1]
			sumArrayTest1[4] += u.T1[i].Answers[2]
			sumArrayTest1[8] += u.T1[i].Answers[3]
		case 30:
			sumArrayTest1[0] += u.T1[i].Answers[0]
			sumArrayTest1[2] += u.T1[i].Answers[1]
			sumArrayTest1[6] += u.T1[i].Answers[2]
			sumArrayTest1[9] += u.T1[i].Answers[3]
		case 31:
			sumArrayTest1[6] += u.T1[i].Answers[0]
			sumArrayTest1[9] += u.T1[i].Answers[1]
			sumArrayTest1[10] += u.T1[i].Answers[2]
			sumArrayTest1[11] += u.T1[i].Answers[3]
		case 32:
			sumArrayTest1[9] += u.T1[i].Answers[0]
			sumArrayTest1[11] += u.T1[i].Answers[1]
			sumArrayTest1[4] += u.T1[i].Answers[2]
			sumArrayTest1[6] += u.T1[i].Answers[3]
		}
	}

	headers := []string{
		"Вознаграждение",
		"Условия работы",
		"Cтруктурирование работы",
		"Социальные контакты",
		"Взаимоотношения",
		"Признание",
		"Достижения",
		"Власть и влиятельность",
		"Разнообразие",
		"Креативность",
		"Самосовершенствование",
		"Интересная работа",
	}
	userEmail.T1 = test1Email{Headers:headers, Values:sumArrayTest1}
}
