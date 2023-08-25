package main

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
	"golang.org/x/net/context"
)

func main() {

	TenantID := "c524e684-89b9-41e5-89c5-546aa8741088"
	ClientID := "02f84e6f-f960-4543-a681-300ad761401e"
	Scrit := "Q9Q8Q~7KVCqMBoquHU8tJ1Nl7hn5yu-HVMFCqaNa"

	cred, err := azidentity.NewClientSecretCredential(TenantID, ClientID, Scrit, nil)

	if err != nil {
		fmt.Printf("Error creating credentials: %v\n", err)
		return
	}

	socps := []string{"https://servicebus.azure.net//.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, socps)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	requestParameters := &graphusers.ItemMessagesRequestBuilderGetQueryParameters{
		Select: []string{"sender", "subject"},
	}
	configuration := &graphusers.ItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	r, err := client.Me().Messages().Get(context.TODO(), configuration)
	if err != nil {
		fmt.Printf("Error getting the messge: %v\n", err)
		return
	}

	fmt.Printf("結果 : %v", r.GetOdataCount())

	// var topValue int32 = 25
	// query := &graphusers.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
	// 	// Only request specific properties
	// 	Select: []string{"from", "isRead", "receivedDateTime", "subject"},
	// 	// Get at most 25 results
	// 	Top: &topValue,
	// 	// Sort by received time, newest first
	// 	Orderby: []string{"receivedDateTime DESC"},
	// }
	//
	// messages, err := client.Me().MailFolders().ByMailFolderId("inbox").
	// 	Messages().
	// 	Get(context.Background(),
	// 		&graphusers.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
	// 			QueryParameters: query,
	// 		})
	//
	// if err != nil {
	// 	fmt.Printf("Error getting the messge: %v\n", err)
	// 	return
	// }
	//
	// location, err := time.LoadLocation("Local")
	// if err != nil {
	// 	log.Panicf("Error getting local timezone: %v", err)
	// }
	//
	// // Output each message's details
	// for _, message := range messages.GetValue() {
	// 	fmt.Printf("Message: %s\n", *message.GetSubject())
	// 	fmt.Printf("  From: %s\n", *message.GetFrom().GetEmailAddress().GetName())
	//
	// 	status := "Unknown"
	// 	if *message.GetIsRead() {
	// 		status = "Read"
	// 	} else {
	// 		status = "Unread"
	// 	}
	// 	fmt.Printf("  Status: %s\n", status)
	// 	fmt.Printf("  Received: %s\n", (*message.GetReceivedDateTime()).In(location))
	// }
	//
	// // If GetOdataNextLink does not return nil,
	// // there are more messages available on the server
	// nextLink := messages.GetOdataNextLink()
	//
	// fmt.Println()
	// fmt.Printf("More messages available? %t\n", nextLink != nil)
	// fmt.Println()
}

//

// 	fmt.Printf("結果 : %v", r.GetOdataCount())
// }
