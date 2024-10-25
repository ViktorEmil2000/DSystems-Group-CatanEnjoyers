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

type ChitChatServer struct {
}

// mustEmbedUnimplementedServicesServer implements ServicesServer.
func (ccs *ChitChatServer) mustEmbedUnimplementedServicesServer() {
	panic("unimplemented")
}

func (ccs *ChitChatServer) ChatService(csi Services_ChatServiceServer) error {
	errch := make(chan error)

	ID_ := rand.Intn(1e6)

	go receiveFromStream(csi, ID_, errch)

	go sendToStream(csi, ID_, errch)
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

func sendToStream(csi Services_ChatServiceServer, ID_ int, errch chan error) {
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

			if ID != ID_ {
				err := csi.Send(&FromServer{Name: senderName4Client, Body: message4Client})

				if err != nil {
					errch <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []messageStruc{}
				}

				messageHandleObject.mu.Unlock()
			}

		}
		time.Sleep(100 * time.Millisecond)

	}
}
