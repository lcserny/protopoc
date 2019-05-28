package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	leo := &Person{
		Name: "Leonardo",
		Age:  31,
		SocialFollowers: &SocialFollowers{
			Youtube: 1000,
			Twitter: 2500,
		},
	}

	data, err := proto.Marshal(leo)
	if err != nil {
		log.Fatal("Could not marshal person", err)
	}

	fmt.Println(data)
	_ = ioutil.WriteFile("person.data", data, os.ModePerm)

	newLeo := new(Person)
	//err = proto.Unmarshal(data, newLeo)
	bytes, _ := ioutil.ReadFile("person.data")
	err = proto.Unmarshal(bytes, newLeo)
	if err != nil {
		log.Fatal("Could not unmarshal person", err)
	}

	fmt.Println(newLeo.GetName())
	fmt.Println(newLeo.GetAge())
	fmt.Println(newLeo.GetSocialFollowers().GetYoutube())
	fmt.Println(newLeo.GetSocialFollowers().GetTwitter())
}
