package utils
import (
	"encoding/json"
	"github.com/revel/revel"
)

func Log(i interface{}) {
	revel.INFO.Println(i)
}

func LogJson(i interface{}) {
	b, _ := json.MarshalIndent(i, "", " ")
	revel.INFO.Println(string(b))
}
