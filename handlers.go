package main

import (
	"net/http"
	"time"
	"strconv"
	"log"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	//"encoding/json"
	"encoding/json"
)

func Test3Handle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	if u.Level != 4 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {

		r.ParseForm()
		arrAnswers := make([]testAnswer, 0, 0)
		for i := 0; i < len(dbTest3); i++ {
			s := "name" + strconv.Itoa(i)
			tAnswer := testAnswer{Id: i, Answers:make([]int, 1, 1) }
			s1 := r.FormValue(s)
			if s1 == "" {
				s1 = "0"
			}
			//log.Println(s1)
			value, err := strconv.Atoi(s1)
			if err != nil {
				log.Println(err)
				return
			}
			tAnswer.Answers[0] = value

			arrAnswers = append(arrAnswers, tAnswer)
		}
		u.T3 = arrAnswers
		u.Level++
		app.dbMongo.UpdateUser(&u)
	}
	UserEmail := NewUserForEmail(&u)
	sendEMail(&UserEmail)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func Test2Handle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	if u.Level != 3 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {

		r.ParseForm()
		arrAnswers := make([]testAnswer, 0, 0)
		for i := 0; i < len(dbTest2); i++ {
			s := "name" + strconv.Itoa(i) + "-"
			tAnswer := testAnswer{Id: i, Answers:make([]int, 8, 8) }
			sum := 0
			countAnswers := 0
			for j := 0; j < 8; j++ {
				s1 := r.FormValue(s + strconv.Itoa(j))
				if s1 == "" {
					s1 = "0"
				}
				value, err := strconv.Atoi(s1)
				if err != nil {
					log.Println(err)
					return
				}
				if value != 0 {
					countAnswers++
					sum += value
				}
				tAnswer.Answers[j] = value
			}
			if sum != 10 || countAnswers > 4 {
				w.Write([]byte("Ошибка с количеством в ответе"))
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			arrAnswers = append(arrAnswers, tAnswer)
		}
		u.T2 = arrAnswers
		u.Level++
		app.dbMongo.UpdateUser(&u)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func Test1Handle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	if u.Level != 2 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		arrAnswers := make([]testAnswer, 0, 0)
		for i := 0; i < len(dbTest1); i++ {
			s := "name" + strconv.Itoa(i) + "-"
			t1Answer := testAnswer{Id: i, Answers:make([]int, 4, 4) }
			sum := 0
			for j := 0; j < 4; j++ {
				s1 := r.FormValue(s + strconv.Itoa(j))
				if s1 == "" {
					s1 = "0"
				}
				value, err := strconv.Atoi(s1)
				if err != nil {
					log.Println(err)
					return
				}
				sum += value
				t1Answer.Answers[j] = value
			}
			if sum != 11 {
				log.Println("Ошибка с количеством в ответе")
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			arrAnswers = append(arrAnswers, t1Answer)
		}
		u.T1 = arrAnswers
		u.Level++
		//log.Println(u)
		app.dbMongo.UpdateUser(&u)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

type testSt struct {
	string
	float32
}

func GetDataTest1(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if u.Level != 5 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		data := make([]testSt, 0, 0)

		dataPrepare := NewUserForEmail(&u)

		for i := 0; i < 12; i++ {
			data = append(data, testSt{dataPrepare.T1.Headers[i], float32(dataPrepare.T1.Values[i])})
		}
		data2, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data2))
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func GetDataTest2(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if u.Level != 5 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		data := make([]testSt, 0, 0)

		dataPrepare := NewUserForEmail(&u)

		for i := 0; i < 8; i++ {
			data = append(data, testSt{dataPrepare.T2.Headers[i], float32(dataPrepare.T2.Values[i])})
		}
		data2, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data2))
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func GetDataTest3(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if u.Level != 5 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		data := make([]testSt, 0, 0)

		dataPrepare := NewUserForEmail(&u)

		for i := 0; i < 3; i++ {
			data = append(data, testSt{dataPrepare.T3.Headers[i], float32(dataPrepare.T3.Values[i])})
		}
		data2, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data2))
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}


func AnketaHandle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	if u.Level != 1 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()

		var ank anketa
		ank.Name = r.FormValue("name")
		temp, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		ank.Age = temp

		temp, err = strconv.Atoi(r.FormValue("education"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		ank.Education = temp

		temp, err = strconv.Atoi(r.FormValue("experience"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		ank.Experience = temp

		temp, err = strconv.Atoi(r.FormValue("howlongwork"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		ank.HowLongWork = temp

		temp, err = strconv.Atoi(r.FormValue("howmanypeople"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		ank.HowManyPeople = temp

		temp, err = strconv.Atoi(r.FormValue("howlongworkprof"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		ank.Sex = r.FormValue("sex")

		ank.HowLongWorkProf = temp
		u.Ank = ank
		u.Level++
		app.dbMongo.UpdateUser(&u)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func AboutHandle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	if u.Level != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		u.Level++
		app.dbMongo.UpdateUser(&u)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//showSessions()
	//log.Println(u.Level)
	//log.Println(u)
	switch {
	case u.Level == 0:
		level0(w)
	case u.Level == 1:
		level1(w)
	case u.Level == 2:
		level2(w)
	case u.Level == 3:
		level3(w)
	case u.Level == 4:
		level4(w)
	case u.Level == 5:
		level5(w, &u)
	default:
		tpl.ExecuteTemplate(w, "index.gohtml", u)
	}
}

func level0(w http.ResponseWriter) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func level1(w http.ResponseWriter) {
	tpl.ExecuteTemplate(w, "anketa.gohtml", nil)
}

func level2(w http.ResponseWriter) {
	tpl.ExecuteTemplate(w, "test1.gohtml", dbTest1)
}

func level3(w http.ResponseWriter) {
	tpl.ExecuteTemplate(w, "test2.gohtml", dbTest2)
}

func level4(w http.ResponseWriter) {
	tpl.ExecuteTemplate(w, "test3.gohtml", dbTest3)
}

func level5(w http.ResponseWriter, u *user) {
	tpl.ExecuteTemplate(w, "graph.gohtml", u)
}

func SignupHandle(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		_, err := app.dbMongo.GetUser(un)
		if err == nil {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		password := req.FormValue("password")
		if password == "" {
			http.Error(w, "Empty password", http.StatusForbidden)
			return
		}
		//create session
		sID := uuid.NewV4()

		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		u := user{un, anketa{}, []testAnswer{}, []testAnswer{}, []testAnswer{}, 0, cryptPassword}

		err = app.dbMongo.AddUser(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		u, err := app.dbMongo.GetUser(un)
		if err != nil {
			http.Error(w, "Username and/or password do not mutch", http.StatusForbidden)
			return
		}
		password := r.FormValue("password")
		err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and/or password do not mutch", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	//showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")

	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > time.Hour*2 {
		go cleanSessions()
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
