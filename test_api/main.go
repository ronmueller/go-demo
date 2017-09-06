package main

import (
	"api/controllers"
	"api/models"
	"fmt"
)

func main(){
	cbo := new(controllers.CBase)
	_ = cbo.CBConnect()
	var story models.Story
	_, err := cbo.Bucket["content"].Get("story_30868966", &story)
	
	if err == nil {
		fmt.Println(story)
	}
}
