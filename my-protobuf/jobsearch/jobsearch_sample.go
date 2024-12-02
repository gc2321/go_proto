package jobsearch

import (
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The Amazing Proto",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software: ", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId:     887,
			ApplicantFullName: "Applicant Name",
			Phone:             "999-999-8888",
			Email:             "someone@example.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Job Application: ", string(jsonBytes))
}
