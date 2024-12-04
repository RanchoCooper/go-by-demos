package main

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const mailTableName = "email"
const syncInterval = 30 * time.Second

type MailFirebase struct {
	To      string          `firestore:"to"`
	Message MessageFirebase `firestore:"message"`
}

type MessageFirebase struct {
	HTML    string `firestore:"html"` // use either html or text field to have plain text email body. Both won't work.
	Subject string `firestore:"subject"`
	Text    string `firestore:"text"`
}

func (Mail) TableName() string {
	return mailTableName
}

type Mail struct {
	db *firestore.Client
}

func NewMail(engine *firestore.Client) *Mail {
	return &Mail{
		db: engine,
	}
}

func (m *Mail) Insert(ctx context.Context, mail MailFirebase) error {
	docRef, _, err := m.db.Collection(mailTableName).Add(ctx, mail)
	if err != nil {
		return err
	}
	docSnapshot, err := m.db.Collection(mailTableName).Doc(docRef.ID).Get(ctx)
	data := docSnapshot.Data()
	if _, ok := data["delivery"]; !ok {
		return nil
	}
	fmt.Println(data["delivery"])
	return nil
}

func main() {
	// 初始化 Firebase 应用程序
	file := "./../../serviceAccountKey.json"
	serviceAccountKeyFilePath, err := filepath.Abs(file)
	if err != nil {
		fmt.Println("Unable to load serviceAccountKey.json file", err)
		return
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("Error initializing Firebase app:", err)
		return
	}
	// client, err := app.Auth(context.Background())
	// client.GetUsers(context.Background())
	// client.Users()

	firestoreClient, err := app.Firestore(context.Background())
	email := NewMail(firestoreClient)

	ctx := context.Background()
	docSnapshot, err := email.db.Collection(mailTableName).Doc("CcL4VSsVMtP3A28FYcCl").Get(ctx)
	data := docSnapshot.Data()
	if _, ok := data["delivery"]; ok {
		fmt.Println(data["delivery"])
		if _, ok := data["delivery"].(map[string]interface{})["state"]; ok {
			if data["delivery"].(map[string]interface{})["state"] == "SUCCESS" {
				fmt.Println("current email is Success")
			}
		}
	}
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}
