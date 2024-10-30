package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

const (
	TestChannel      = "C04HFJ7Q5GR"
	JavisTestChannel = "C076KTC9SD7"
	JavisProdChannel = "C076ASN3HA9"
)

func main() {
	api := slack.New("SLACK-TOKEN")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	// groups, err := api.GetUserGroups(slack.GetUserGroupsOptionIncludeUsers(false))
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return
	// }
	// for _, group := range groups {
	// 	fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	// }

	// user, err := api.GetUserByEmail("rancho941110@gmail.com")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user.Name)
	a, b, err := api.PostMessage(JavisTestChannel, slack.MsgOptionText("hello, it's a testing message", false))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a, b)

}
