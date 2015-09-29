package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

/*
import "code.google.com/p/gcfg"
*/

/*
Порядок инсталляции:
- проверяем темплэйты
- создаем окружение
- создаем БД, БД-юзеров и расширение постгис
- заливаем схему из дампа БД (одной большой транзакцией)
- заливаем скрипты: 3 крона
- правим schema_initialization_timestamp
- заливаем скрипты: бэкап / рестор (рестор онли экзампл комментс)
- заливаем кронтаб на кроны + бэкап


Objects: methods
- application: reads config, gets templates, creates environment
- dbHelper: creates database, creates users, installs postgis, installs schema
- scriptHelper: checks file existence, creates file, writes file, appends file
- connectHelper: make connects, parse results
*/

const (
	host          = ""
	platformAppid = ""
	myLogin       = ""
	myPassword    = ""
)

func main() {

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	session, err := login(myLogin, myPassword, cookieJar)
	if err != nil {
		panic(err)
	}

	fmt.Println(session)

	err = logout(session, cookieJar)
	if err != nil {
		panic(err)
	}

	return
}

/*----------------------------------------------------------------------------*/
func login(login, password string, cookieJar *cookiejar.Jar) (session string, err error) {

	// prepare request
	rest := "/users/authentication/rest/signin"
	data := make(map[string]string)
	data["appid"] = platformAppid
	data["login"] = login
	data["password"] = password

	// make request
	res, err := postHttps(host+rest, data, cookieJar)
	if err != nil {
		return "", err
	}

	// decode result
	var v map[string]interface{}
	err = json.Unmarshal(res, &v)
	if err != nil {
		return "", err
	}
	// extract session
	for k := range v {
		if k == "session" {
			session = v[k].(string)
		}
	}

	return session, nil
}

/*----------------------------------------------------------------------------*/
func logout(session string, cookieJar *cookiejar.Jar) error {
	// prepare request
	rest := "/users/authentication/rest/signout"
	data := make(map[string]string)
	data["appid"] = platformAppid
	data["session"] = session

	// make request
	res, err := postHttps(host+rest, data, cookieJar)
	if err != nil {
		return err
	}

	// decode result
	var v map[string]interface{}
	err = json.Unmarshal(res, &v)
	if err != nil {
		return err
	}

	// extract result
	for k := range v {
		if k == "error" {
			return errors.New(v[k].(string))
		}
		if k == "result" {
			if v[k].(float64) == 0 {
				return nil //all OK
			}
		}

	}

	return nil
}

/*----------------------------------------------------------------------------*/
func postHttps(request string, data map[string]string, cookieJar *cookiejar.Jar) ([]byte, error) {

	client := &http.Client{
		Jar: cookieJar,
	}

	postData := url.Values{}
	for k, v := range data {
		postData.Add(k, v)
	}

	resp, err := client.PostForm(request, postData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
