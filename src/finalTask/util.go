package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// created a custom type
type tech []string

// assiging the type to the struct
type container struct {
	Docker tech
}

type techStack struct {
	BackEnd tech
	Cloud   tech
	docker  container
}

type listTechnologies interface {
	listBackend() tech
	listDevops() tech
}

func (t tech) print() {
	for _, list := range t {
		fmt.Println(list)
	}
}

func genTechForWeek() techStack {
	techList := techStack{
		BackEnd: tech{"GoLang-Udemy", "GoLang-Book", "CRUD using Go & SQL", "GoLang WebDev"},
		Cloud:   tech{"AWS-Basics", "AWS Security", "ECS & Fargate", "Database & Ec2"},
		docker: container{
			Docker: tech{"Docker-Basics", "Docker-Compose", "CI/CD using docker & Github"},
		},
	}

	return techList
}

func genUrl() tech {
	urlList := tech{
		"http://www.google.com/",
		"http://in.yahoo.com/?p=us",
		"http://stackoverflow.com/",
		"http://www.youtube.com/",
	}

	return urlList
}

func (t techStack) listBackend() tech {
	Btech := tech{}
	for _, temp := range t.BackEnd {
		//fmt.Println("BackEnd Technologies")
		Btech = append(Btech, temp)
	}
	return Btech
}

func (cl techStack) listDevops() tech {
	cloud := tech{}
	for _, temp := range cl.Cloud {
		cloud = append(cloud, temp)

	}
	for _, tep := range cl.docker.Docker {
		cloud = append(cloud, tep)
	}
	//fmt.Println("Devops Technologies")
	return cloud
}

func (ss tech) toString() string {
	return strings.Join([]string(ss), ",")
}

// file writing

//write file
func (t tech) writeFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(t.toString()), 0644)
}

// readFile
func (t tech) readFile(fileName string) tech {
	resp, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File Not Found")
	}
	// passing the byte and converting to string
	s := strings.Split(string(resp), ",") // you will get []strig as output
	return tech(s)
}
