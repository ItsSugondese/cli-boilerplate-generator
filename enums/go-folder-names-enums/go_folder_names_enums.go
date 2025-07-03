package go_folder_names_enums_enums

import "reflect"

var GoFolderNamesEnums = newGoFolderNamesEnums()

func newGoFolderNamesEnums() goFolderNamesEnums {
	return goFolderNamesEnums{
		Repo:       "repo",
		Controller: "controller",
		DTO:        "dto",
		Model:      "model",
		Route:      "route",
		Navigator:  "navigator",
		Service:    "service",
	}
}

type goFolderNamesEnums struct {
	Repo       string
	Controller string
	DTO        string
	Model      string
	Route      string
	Navigator  string
	Service    string
}

func (e goFolderNamesEnums) GetKeys() []string {
	t := reflect.TypeOf(e)
	keys := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		keys = append(keys, t.Field(i).Name)
	}

	return keys
}

func (e goFolderNamesEnums) GetValues() []string {
	v := reflect.ValueOf(e)
	values := make([]string, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values = append(values, v.Field(i).String())
	}

	return values
}
