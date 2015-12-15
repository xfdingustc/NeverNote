package utils
import (
	"encoding/json"
	"strconv"
	"fmt"
	"regexp"
)


func IsUsername(username string) bool {
	if username == "" {
		return false
	}

	ok, _ := regexp.MatchString(`[^0-9a-zA-Z_\-]`, username)
	return !ok
}

func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.|\-]?)*[_a-z\-A-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.|\-]?)*[a-zA-Z0-9\-]+\.[0-9a-zA-Z]{2,6}$`, email)
	return ok
}

var rulesStr = `{
	"username": [
		{"rule": "required", "msg": "inputUserName"},
		{"rule": "noSpecialChars", "msg": "noSpecialChars"},
		{"rule": "minLength", "data": "8", "msg": "minLength", "msgData": "4"}
	],
	"email": [
		{"rule": "required", "msg": "inputEmail"},
		{"rule": "email", "msg": "errorEmail"}
	],
	"password": [
		{"rule": "required", "msg": "inputPassword"},
		{"rule": "password", "msg": "errorPassword"},
		{"rule": "minLength", "data": "6", "msg": "minLength", "msgData": "6"}
	],
	"subDomain": [
		{"rule": "subDomain", "msg": "errorSubDomain"}
	]

}
`

var rulesMap map[string][]map[string]string

var rules = map[string]func(string, map[string]string) (bool, string) {
	"required": func(value string, rule map[string]string) (ok bool, msg string) {
		if value == "" {
			return
		}
		ok = true
		return
	},

	"minLength": func(value string, rule map[string]string) (ok bool, msg string) {
		if value == "" {
			return
		}

		data := rule["data"]
		dataI, _ := strconv.Atoi(data)
		ok = len(value) >= dataI
		return
	},

	"noSpecialChars": func(value string, rule map[string]string) (ok bool, msg string) {
		if value == "" {
			return
		}
		ok = IsUsername(value)
		return
	},

	"email": func(value string, rule map[string]string) (ok bool, msg string) {
		if value == "" {
			return;
		}

		ok = IsEmail(value)
		return
	},

	"password": func(value string, rule map[string]string) (ok bool, msg string) {
		if value == "" {
			return
		}
		ok = len(value) >= 6
		return
	},
}

func InitValidate() {
	json.Unmarshal([]byte(rulesStr), &rulesMap)
	LogJson(rulesMap)

}

func Validate(name, value string) (ok bool, msg string)  {
	ruleList, _ := rulesMap[name]
	fmt.Printf("name: %s ruleMapsize: %d\n", name, len(rulesMap))

	for _, rule := range ruleList {
		fmt.Print("check one rule")
		ruleFunc, _ := rules[rule["rule"]]

		if ok2, msg2 := ruleFunc(value, rule); !ok2 {
			fmt.Print("Validate error: ")
			ok = false
			if msg2 != "" {
				msg = msg2
			} else {
				msg = rule["msg"]
			}

			msgData := rule["msgData"]
			if msgData != "" {
				msg += "-" + msgData
			}
			return
		}
	}
	ok = true
	return
}
