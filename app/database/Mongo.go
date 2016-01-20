package database
import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/utils"
	"strings"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)


var Session *mgo.Session
var Users *mgo.Collection
var Notebooks *mgo.Collection

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

	Notebooks = Session.DB(dbname).C("notebooks")

}


// DAO


func Get(collection *mgo.Collection, id string, i interface{}) {
	collection.FindId(bson.ObjectIdHex(id)).One(i)
}

func Insert(collection *mgo.Collection, i interface{}) bool {
	err := collection.Insert(i)
	return Err(err)
}


func GetByQ(collection *mgo.Collection, q interface{}, i interface{}) {
	collection.Find(q).One(i)
}


func Err(err error) bool {
	if err != nil {
		fmt.Println(err)

		if err.Error() == "not found" {
			return true
		}

		return false
	}
	return true
}