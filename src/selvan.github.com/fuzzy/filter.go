package fuzzy

import "sort"
import "fmt"

type ByScore []map[string]interface{}

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i]["score"].(float32) < a[j]["score"].(float32) }

type FilterSource struct {
	Source []interface{}
	Key string
}

func (filter *FilterSource) Filter(search string, max_result int) ([]map[string]interface{}, error) {

	results := []map[string]interface{}{}

	for _, source := range filter.Source {
		switch t := source.(type) {
			case map[string]interface{}:
				if len(filter.Key) == 0 {
					continue
				}
				s := source.(map[string]interface{})
				str := s[filter.Key].(string)
				results = apply(&str, &search, s, results)
			case string:
				str := source.(string)
				results = apply(&str, &search, source, results)
	 		default:
	        	return results, fmt.Errorf("Type %T not supported", t)			
		}	
	}

	if len(results) == 0 {
		return results, nil
	}

	// Sort decending by score, best matched values come first
	sort.Sort(sort.Reverse(ByScore(results)))

	results_count := len(results)

	if (max_result == -1) || (max_result > results_count) {
		max_result = results_count
	}

	return results[:max_result], nil
}

func apply(source *string, search *string, candidate interface{}, results []map[string]interface{}) []map[string]interface{} {
	score := ComputeScore(*source, *search)
    if(score > 0) {	
    	results = append(results, map[string]interface{}{
			"score" : score,
			"candidate" : candidate,
		})
    }
    return results
}
