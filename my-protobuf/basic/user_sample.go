package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
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

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:                   99,
		Username:             "supername",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Emails:               []string{"superman@example.com", "superman@dc.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
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

func randomCommunicationChannel() anypb.Any {
	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "krypton.man",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "krypton.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))

}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()
	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)

	log.Println(string(jsonBytes))

}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)

		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but :", a.TypeUrl)
	}
}
