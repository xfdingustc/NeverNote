package utils
import (
	"encoding/json"
	"github.com/revel/revel"
)

func LogJson(i interface{}) {
	b, _ := json.MarshalIndent(i, "", " ")
	revel.INFO.Println(string(b))
}
