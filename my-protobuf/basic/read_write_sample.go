package basic

import (
	"fmt"
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) {
	b, err := proto.Marshal(msg)

	if err != nil {
		log.Fatalf("Can not marshal message: %v", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalf("Can not write to file: %v", err)
	}
}

func ReadProtoFromFile(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalf("can not read file: %v", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalf("can not unmarshal %v", err)
	}
}

func dummyUser() basic.User {
	addr := basic.Address{
		Street:  "123 Maple Street",
		City:    "My Town",
		Country: "US",
		Coordinate: &basic.Address_Corrdinate{
			Latitude:  40.709,
			Longitude: -74.011,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{
		"fly":        5,
		"speed":      5,
		"durability": 4,
	}

	return basic.User{
		Id:                   99,
		Username:             "supername",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Emails:               []string{"superman@example.com", "superman@dc.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: comm,
		SkillRating:          sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func ReadFromFileSample() {
	dest := basic.User{}
	ReadProtoFromFile("superman_file.bin", &dest)
	fmt.Println(&dest)
}
