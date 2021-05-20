package router

import "cdn/util"

type JsonObject struct {

	Key string
	Value string

}

func Stringify(objects ...JsonObject) string {
	ret := "{\n"
	for _, object := range objects {
		ret += "\t" + object.Stringify() + ",\n"
	}
	ret = ret[:len(ret)-2]
	ret += "\n}"
	return ret
}

func (json *JsonObject) Stringify() string {
	return util.Format("\"{}\":\"{}\"", json.Key, json.Value)
}
