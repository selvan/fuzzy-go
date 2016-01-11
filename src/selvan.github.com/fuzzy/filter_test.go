package fuzzy_test

import "selvan.github.com/fuzzy"
import "testing"

func TestFilter(t *testing.T) {

	var data []interface{}

	data = append(data, 
		map[string]interface{}{"place" : "bangalore" , "id" : 1},
		map[string]interface{}{"place" : "chennai" , "id" : 1},
		map[string]interface{}{"place" : "bombay" , "id" : 1},
		)
	
	filter_source := new(fuzzy.FilterSource)
	filter_source.Source = data
	filter_source.Key = "place"

	results, _ := filter_source.Filter("bo", -1)
	t.Logf("Result length %d", len(results))
	r := map[string]interface{}(results[0])
	t.Logf("Score is %f", r["score"].(float32))

	v := r["candidate"].(map[string]interface{})
	t.Logf(" place is %s", v["place"])
}