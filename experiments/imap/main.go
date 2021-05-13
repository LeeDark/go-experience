package main

import (
	"io/ioutil"
	"log"
	"net/mail"
	"net/textproto"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS("mail.horisen.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	log.Println("Client state: ", c.State())

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("sergey@horisen.com", "************"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	log.Println("Client state: ", c.State())

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Client state: ", c.State())

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	//// Get the last 4 messages
	//from := uint32(1)
	//to := mbox.Messages
	//if mbox.Messages > 3 {
	//	// We're using unsigned integers here, only subtract if the result is > 0
	//	from = mbox.Messages - 3
	//}
	//seqset := new(imap.SeqSet)
	//seqset.AddRange(from, to)
	//
	//messages := make(chan *imap.Message, 10)
	//done = make(chan error, 1)
	//go func() {
	//	done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	//}()
	//
	//log.Println("Last 4 messages:")
	//for msg := range messages {
	//	log.Println("* " + msg.Envelope.Subject)
	//}
	//
	//if err := <-done; err != nil {
	//	log.Fatal(err)
	//}

	log.Println("Client state: ", c.State())

	dateSince, _ := time.Parse(imap.DateLayout, "1-Nov-2020")
	dateBefore, _ := time.Parse(imap.DateLayout, "1-Dec-2020")

	criteria := &imap.SearchCriteria{
		Header: textproto.MIMEHeader{
			"From":    {"noreply-reportmailer-prod@tnsi.com"},
			"Subject": {"Master Carrier ID Report"},
		},
		Since:  dateSince,
		Before: dateBefore,
	}

	done = make(chan error, 1)
	var results []uint32
	go func() {
		var err error
		results, err = c.Search(criteria)
		done <- err
	}()

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Results: ", results)

	if len(results) == 1 {
		seqset := new(imap.SeqSet)
		seqset.AddNum(results[0])

		// Get the whole message body
		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem(), imap.FetchBodyStructure}

		messages := make(chan *imap.Message, 1)
		done = make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, items, messages)
		}()

		log.Println("Message:")
		msg := <-messages

		//log.Println(msg.BodyStructure)
		//log.Println(msg.BodyStructure.MIMEType)
		//log.Println(msg.BodyStructure.MIMESubType)
		//
		//log.Println(msg.BodyStructure.BodyStructure)
		//log.Println(msg.BodyStructure.Parts)
		//log.Println(msg.BodyStructure.Parts[0])
		//log.Println(msg.BodyStructure.Parts[0].MIMEType)
		//log.Println(msg.BodyStructure.Parts[0].MIMESubType)
		//log.Println(msg.BodyStructure.Parts[0].Parts)
		//log.Println(msg.BodyStructure.Parts[0].Parts[0])
		//log.Println(msg.BodyStructure.Parts[0].Parts[1])
		//log.Println(msg.BodyStructure.Parts[0].Parts[2])
		//log.Println(msg.BodyStructure.Parts[1])
		//log.Println(msg.BodyStructure.Parts[1].MIMEType)
		//log.Println(msg.BodyStructure.Parts[1].MIMESubType)
		////file, err := msg.BodyStructure.MIMEType

		r := msg.GetBody(section)
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		m, err := mail.ReadMessage(r)
		if err != nil {
			log.Fatal(err)
		}

		header := m.Header
		log.Println("Date:", header.Get("Date"))
		log.Println("From:", header.Get("From"))
		log.Println("To:", header.Get("To"))
		log.Println("Subject:", header.Get("Subject"))

		body, err := ioutil.ReadAll(m.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(len(body))
	}

	log.Println("Done!")
}
