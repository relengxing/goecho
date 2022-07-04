package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			YouTube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's Go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for Good measure
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYouTube())
}
