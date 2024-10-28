package proto

import (
	"log"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type messageStruc struct {
	ID       int
	Username string
	Message  string
}

type messageHandle struct {
	MQue []messageStruc
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

type ConnectedUser struct {
	ID  int
	csi Services_ChatServiceServer
}

var UList = []ConnectedUser{}

var sendToStreamCheck = 0

type ChitChatServer struct {
}

// mustEmbedUnimplementedServicesServer implements ServicesServer.
func (ccs *ChitChatServer) mustEmbedUnimplementedServicesServer() {
	panic("unimplemented")
}

func (ccs *ChitChatServer) ChatService(csi Services_ChatServiceServer) error {
	errch := make(chan error)

	ID_ := rand.Intn(1e6)

	UList = append(UList, ConnectedUser{
		ID:  ID_,
		csi: csi,
	})

	go receiveFromStream(csi, ID_, errch)
	if sendToStreamCheck == 0 {
		go sendToStream()
		sendToStreamCheck = 1
	}

	return <-errch
}

// Recieves messages from stream
func receiveFromStream(csi Services_ChatServiceServer, ID_ int, errch chan error) {
	kill := true
	name := "User"
	for kill == true {
		msg, err := csi.Recv()
		if err != nil {
			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageStruc{
				ID:       ID_,
				Username: name,
				Message:  "Disconnected",
			})
			messageHandleObject.mu.Unlock()
			for i, element := range UList {
				if element.ID == ID_ {
					UList = append(UList[:i], UList[i+1:]...)
				}

			}

			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
			kill = false
		} else {
			if name == "User" {
				name = msg.Name
			}
			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageStruc{
				ID:       ID_,
				Username: msg.Name,
				Message:  msg.Body,
			})

			messageHandleObject.mu.Unlock()

			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
		}
	}
}

func sendToStream() {
	for {

		for {
			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}
			ID := messageHandleObject.MQue[0].ID
			senderName4Client := messageHandleObject.MQue[0].Username
			message4Client := messageHandleObject.MQue[0].Message

			messageHandleObject.mu.Unlock()

			for _, element := range UList {
				if element.ID != ID {
					err := element.csi.Send(&FromServer{Name: senderName4Client, Body: message4Client})
					if err != nil {
						log.Printf("couldn't send to user of: %d ", element.ID)
					}
				}
			}

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) > 1 {
				messageHandleObject.MQue = messageHandleObject.MQue[1:]
			} else {
				messageHandleObject.MQue = []messageStruc{}
			}

			messageHandleObject.mu.Unlock()

		}
		time.Sleep(100 * time.Millisecond)

	}
}
