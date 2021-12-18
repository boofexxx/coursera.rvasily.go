package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	_ "reflect"
	"testing"
)

// код писать тут
type TestData struct {
	Client   SearchClient
	Request  SearchRequest
	Response *SearchResponse
	IsError  bool
}

func TestFindUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	cases := []TestData{
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL},
			Request: SearchRequest{
				Limit:      26,
				Offset:     0,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: &SearchResponse{
				Users: []User{
					{
						Id:     1,
						Name:   "Hilda Mayer",
						Age:    21,
						About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n",
						Gender: "female",
					},
					{
						Id:     10,
						Name:   "Henderson Maxwell",
						Age:    30,
						About:  "Ex et excepteur anim in eiusmod. Cupidatat sunt aliquip exercitation velit minim aliqua ad ipsum cillum dolor do sit dolore cillum. Exercitation eu in ex qui voluptate fugiat amet.\n",
						Gender: "male",
					},
					{
						Id:     18,
						Name:   "Terrell Hall",
						Age:    27,
						About:  "Ut nostrud est est elit incididunt consequat sunt ut aliqua sunt sunt. Quis consectetur amet occaecat nostrud duis. Fugiat in irure consequat laborum ipsum tempor non deserunt laboris id ullamco cupidatat sit. Officia cupidatat aliqua veniam et ipsum labore eu do aliquip elit cillum. Labore culpa exercitation sint sint.\n",
						Gender: "male",
					},
					{
						Id:     28,
						Name:   "Cohen Hines",
						Age:    32,
						About:  "Deserunt deserunt dolor ex pariatur dolore sunt labore minim deserunt. Tempor non et officia sint culpa quis consectetur pariatur elit sunt. Anim consequat velit exercitation eiusmod aute elit minim velit. Excepteur nulla excepteur duis eiusmod anim reprehenderit officia est ea aliqua nisi deserunt officia eiusmod. Officia enim adipisicing mollit et enim quis magna ea. Officia velit deserunt minim qui. Commodo culpa pariatur eu aliquip voluptate culpa ullamco sit minim laboris fugiat sit.\n",
						Gender: "male",
					},
					{
						Id:     29,
						Name:   "Clarissa Henry",
						Age:    34,
						About:  "Nostrud enim ea ad reprehenderit tempor ullamco exercitation. Elit in voluptate pariatur sit nisi occaecat laboris esse ipsum. Mollit elit et deserunt ea laboris sunt est amet culpa laboris occaecat ipsum sunt sunt.\n",
						Gender: "female",
					},
				},
				NextPage: false,
			},
			IsError: false,
		},
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL},
			Request: SearchRequest{
				Limit:      2,
				Offset:     1,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: &SearchResponse{
				Users: []User{
					{
						Id:     10,
						Name:   "Henderson Maxwell",
						Age:    30,
						About:  "Ex et excepteur anim in eiusmod. Cupidatat sunt aliquip exercitation velit minim aliqua ad ipsum cillum dolor do sit dolore cillum. Exercitation eu in ex qui voluptate fugiat amet.\n",
						Gender: "male",
					},
					{
						Id:     18,
						Name:   "Terrell Hall",
						Age:    27,
						About:  "Ut nostrud est est elit incididunt consequat sunt ut aliqua sunt sunt. Quis consectetur amet occaecat nostrud duis. Fugiat in irure consequat laborum ipsum tempor non deserunt laboris id ullamco cupidatat sit. Officia cupidatat aliqua veniam et ipsum labore eu do aliquip elit cillum. Labore culpa exercitation sint sint.\n",
						Gender: "male",
					},
				},
				NextPage: false,
			},
			IsError: false,
		},
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL},
			Request: SearchRequest{
				Limit:      10,
				Offset:     1,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: &SearchResponse{
				Users: []User{
					{
						Id:     10,
						Name:   "Henderson Maxwell",
						Age:    30,
						About:  "Ex et excepteur anim in eiusmod. Cupidatat sunt aliquip exercitation velit minim aliqua ad ipsum cillum dolor do sit dolore cillum. Exercitation eu in ex qui voluptate fugiat amet.\n",
						Gender: "male",
					},
					{
						Id:     18,
						Name:   "Terrell Hall",
						Age:    27,
						About:  "Ut nostrud est est elit incididunt consequat sunt ut aliqua sunt sunt. Quis consectetur amet occaecat nostrud duis. Fugiat in irure consequat laborum ipsum tempor non deserunt laboris id ullamco cupidatat sit. Officia cupidatat aliqua veniam et ipsum labore eu do aliquip elit cillum. Labore culpa exercitation sint sint.\n",
						Gender: "male",
					},
					{
						Id:     28,
						Name:   "Cohen Hines",
						Age:    32,
						About:  "Deserunt deserunt dolor ex pariatur dolore sunt labore minim deserunt. Tempor non et officia sint culpa quis consectetur pariatur elit sunt. Anim consequat velit exercitation eiusmod aute elit minim velit. Excepteur nulla excepteur duis eiusmod anim reprehenderit officia est ea aliqua nisi deserunt officia eiusmod. Officia enim adipisicing mollit et enim quis magna ea. Officia velit deserunt minim qui. Commodo culpa pariatur eu aliquip voluptate culpa ullamco sit minim laboris fugiat sit.\n",
						Gender: "male",
					},
					{
						Id:     29,
						Name:   "Clarissa Henry",
						Age:    34,
						About:  "Nostrud enim ea ad reprehenderit tempor ullamco exercitation. Elit in voluptate pariatur sit nisi occaecat laboris esse ipsum. Mollit elit et deserunt ea laboris sunt est amet culpa laboris occaecat ipsum sunt sunt.\n",
						Gender: "female",
					},
				},
				NextPage: false,
			},
			IsError: false,
		},
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL},
			Request: SearchRequest{
				Limit:      1,
				Offset:     0,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: &SearchResponse{
				Users: []User{
					{
						Id:     1,
						Name:   "Hilda Mayer",
						Age:    21,
						About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n",
						Gender: "female",
					},
				},
				NextPage: true,
			},
			IsError: false,
		},
		// ERRORS
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL}, // Limit is less than 0
			Request: SearchRequest{
				Limit:      -1,
				Offset:     0,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: nil,
			IsError:  true,
		},
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL}, // Offset is less than 0
			Request: SearchRequest{
				Limit:      1,
				Offset:     -1,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: nil,
			IsError:  true,
		},
		{
			Client: SearchClient{AccessToken: "", URL: ts.URL}, // AccessToken is wrong
			Request: SearchRequest{
				Limit:      1,
				Offset:     0,
				Query:      "H",
				OrderField: "",
				OrderBy:    0,
			},
			Response: nil,
			IsError:  true,
		},
		{
			Client: SearchClient{AccessToken: "SomeToken", URL: ts.URL}, // OrderField is wrong
			Request: SearchRequest{
				Limit:      1,
				Offset:     0,
				Query:      "H",
				OrderField: "Wrong",
				OrderBy:    0,
			},
			Response: nil,
			IsError:  true,
		},
	}
	for i := range cases {
		result, err := cases[i].Client.FindUsers(cases[i].Request)
		if err != nil && !cases[i].IsError {
			t.Errorf("[%d] unexpected error: %#v", i, err)
		}
		if err == nil && cases[i].IsError {
			t.Errorf("[%d] expected error, got nil", i)
		}
		if !checkResponses(cases[i].Response, result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", i, cases[i].Response, result)
		}
	}
	ts.Close()
}

func checkUsers(l, r *User) bool {
	if l.Id != r.Id {
		fmt.Printf("Id: %d != %d\n", l.Id, r.Id)
		return false
	} else if l.Age != r.Age {
		fmt.Printf("Age: %d != %d\n", l.Age, r.Age)
		return false
	} else if l.Name != r.Name {
		fmt.Printf("Name: %s != %s\n", l.Name, r.Name)
		return false
	} else if l.Gender != r.Gender {
		fmt.Printf("Gender: %s != %s\n", l.Gender, r.Gender)
		return false
	} else if l.About != r.About {
		fmt.Printf("About:\n%s\n!=\n%s\n", l.About, r.About)
		return false
	}
	return true
}

func checkResponses(l, r *SearchResponse) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil {
		return false
	} else if r == nil {
		return false
	}
	if l.NextPage != r.NextPage {
		fmt.Printf("NextPage: %t\n!=\n%t\n", l.NextPage, r.NextPage)
		return false
	}
	if len(l.Users) != len(r.Users) {
		fmt.Printf("len(Users): %d\n!=\n%d\n", len(l.Users), len(r.Users))
		return false
	}
	for i := range l.Users {
		return checkUsers(&l.Users[i], &r.Users[i])
	}
	return true
}
