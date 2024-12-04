package basic

import (
	"fmt"
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalf("Can not marshal message: %v", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalf("Can not write to file: %v", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalf("can not read file: %v", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalf("can not unmarshal %v", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}
	ReadProtoFromJSON("superman_file.json", &dest)
	fmt.Println(&dest)
}
