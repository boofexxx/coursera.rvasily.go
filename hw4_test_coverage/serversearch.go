package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type userServer struct {
	Id        int    `xml:"id"`
	Name      string `json:"name"`
	FirstName string `xml:"first_name"`
	LastName  string `xml:"last_name"`
	Age       int    `xml:"age"`
	About     string `xml:"about"`
	Gender    string `xml:"gender"`
}

type dataServer struct {
	Users []userServer `xml:"row"`
}

var data dataServer

const dataPath = "dataset.xml"

func init() {
	dataXML, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = xml.Unmarshal(dataXML, &data)
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := range data.Users {
		data.Users[i].Name = fmt.Sprintf("%s %s", data.Users[i].FirstName, data.Users[i].LastName)
	}
}

func fillSearchRequest(url *url.Values) (*SearchRequest, error) {
	limit, err := strconv.Atoi(url.Get("limit"))
	if err != nil {
		return nil, err
	}
	offset, err := strconv.Atoi(url.Get("offset"))
	if err != nil {
		return nil, err
	}
	query := url.Get("query")
	orderField := url.Get("order_field")
	orderBy, err := strconv.Atoi(url.Get("order_by"))
	if err != nil {
		return nil, err
	}
	return &SearchRequest{
		Limit:      limit,
		Offset:     offset,
		Query:      query,
		OrderField: orderField,
		OrderBy:    orderBy,
	}, nil
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("AccessToken") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	r.ParseForm()
	sr, err := fillSearchRequest(&r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := make([]userServer, 0, len(data.Users))
	var lessFunc func(int, int) bool = nil
	switch sr.OrderField {
	case "", "Name":
		if sr.OrderBy == -1 {
			lessFunc = func(i, j int) bool {
				return result[i].Name < result[j].Name
			}
		} else if sr.OrderBy == 1 {
			lessFunc = func(i, j int) bool {
				return result[i].Name > result[j].Name
			}
		}
	case "Id":
		if sr.OrderBy == -1 {
			lessFunc = func(i, j int) bool {
				return result[i].Id < result[j].Id
			}
		} else if sr.OrderBy == 1 {
			lessFunc = func(i, j int) bool {
				return result[i].Id > result[j].Id
			}
		}
	case "Age":
		if sr.OrderBy == -1 {
			lessFunc = func(i, j int) bool {
				return result[i].Age < result[j].Age
			}
		} else if sr.OrderBy == 1 {
			lessFunc = func(i, j int) bool {
				return result[i].Age > result[j].Age
			}
		}
	default:
		errOrder := SearchErrorResponse{Error: "ErrorBadOrderField"}
		data, err := json.Marshal(errOrder)
		if err != nil {
			log.Print(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}
	if sr.Query != "" {
		for i := range data.Users {
			if strings.Contains(data.Users[i].Name, sr.Query) ||
				strings.Contains(data.Users[i].About, sr.Query) {
				result = append(result, data.Users[i])
			}
		}
	} else {
		result = data.Users
	}
	if lessFunc != nil {
		sort.Slice(result, lessFunc)
	}
	if sr.Offset <= len(result) && sr.Offset >= 0 {
		if len(result) > sr.Limit {
			result = result[sr.Offset:sr.Limit]
		} else {
			result = result[sr.Offset:]
		}
	}
	data, err := json.Marshal(&result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(data))
	w.WriteHeader(http.StatusOK)
}
