// Package datacircle is a client for the Datacircle API: find the best GTM APIs.
// The API is open, with no key: https://api.datacircle.dev/v1
// Its spec: https://api.datacircle.dev/v1/openapi.json. More at https://datacircle.dev
package datacircle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Base is the API's address.
const Base = "https://api.datacircle.dev/v1"

// Get calls one path of the API and decodes its JSON answer.
func Get(path string, query url.Values) (interface{}, error) {
	u := Base + "/" + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("datacircle: %s", res.Status)
	}
	var out interface{}
	err = json.NewDecoder(res.Body).Decode(&out)
	return out, err
}

// Apis lists every API on the sheet.
func Apis() (interface{}, error) { return Get("apis", nil) }

// Best gives who is best for one API.
func Best(slug string) (interface{}, error) { return Get("apis/"+url.PathEscape(slug)+"/best", nil) }

// Plans lists every plan, filtered by the query.
func Plans(query url.Values) (interface{}, error) { return Get("plans", query) }

// Provider gives one provider's ranks.
func Provider(slug string) (interface{}, error) { return Get("providers/"+url.PathEscape(slug), nil) }
