package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {

	addr := basic.Address{
		Street:  "123 Maple Street",
		City:    "My Town",
		Country: "US",
		Coordinate: &basic.Address_Corrdinate{
			Latitude:  40.709,
			Longitude: -74.011,
		},
	}

	u := basic.User{
		Id:       99,
		Username: "supername",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Emails:   []string{"superman@example.com", "superman@dc.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@example.com", "wm@dc.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}
	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `{"id":98,"username":"wonderwoman","is_active":true,"password":"d29uZGVyX21hbl9wd2Q=","emails":["wonderwoman@example.com","wm@dc.com"],"gender":"GENDER_FEMALE"}`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatalln("Got error : ", err)
	}
	log.Println(&p)
}
