package database
import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/utils"
	"strings"
	"gopkg.in/mgo.v2"
)


var Session *mgo.Session
var Users *mgo.Collection

func Init(url, dbname string) {
	ok := true
	config := revel.Config

	if url == "" {
		url, ok = config.String("db.url")
		if !ok {
			url, ok = config.String("db.urlEnv")
			if ok {
				utils.Log("get db conf from urlEnv: " + url)
			}
		} else {
			utils.Log("get db conf from db.url: " + url)
		}

		if ok {
			urls := strings.Split(url, "/")
			dbname = urls[len(urls) - 1]
		}
	}

	if dbname == "" {
		dbname, _ = config.String("db.dbname")
	}

	if !ok {
		host, _ := revel.Config.String("db.host")
		port, _ := config.String("db.port")
		username, _ := config.String("db.username")
		password, _ := config.String("db.password")

		usernameAndPassword := username + ":" + password + "@"

		if username == "" || password == "" {
			usernameAndPassword = ""
		}

		url = "mongodb://" + usernameAndPassword + host + ":" + port + "/" + dbname
	}

	utils.Log(url)

	var err error
	Session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	Session.SetMode(mgo.Monotonic, true)

	Users = Session.DB(dbname).C("users")

}