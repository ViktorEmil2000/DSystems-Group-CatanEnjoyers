package main 

import {

}

func main(){

	fmt.Println("Enter your name")
	reader :=bufio.NewReader(os.Stdin)
	username,err := reader.ReadString('\n')

	if err!= nil{
		log.Println("Failed to read username")
	}
	
	conn,err := grpc.Dial("50051",grpc.WithInsecure())
	client := ChitChatServer.NewServicesClient(conn)	
	
	stream, err := client.ChatService(context.Background())
	if err != nil{
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

	bl := make(chan bool)
	<-bl
}

type clienthandle struct{
	stream ChitChatServer.Services_ChatServiceServer
	username string
}
func (ch *clienthandle) clientConfig(string username){
	ch.username = username
}

func (ch *clienthandle) sendMessage(){
	for{
		reader :=bufio.NewReader(os.stdin)

		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console :: %v",err)
		}
		clientMessage = String.Trim(clientMessage, "\r\n")

		clientMessageBox := &ChitChatServer.FromClient{
			name: ch.ClientName,
			body: clientMessage,
		}


		err = ch.stream.Send(clientMessageBox)
		if err != nil{
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}
func (ch *clienthandle) receiveMessage() {
	for {
		msg := ch.stream.Recv()
		
	}

	fmt.Printf("%s :%s \n", msg.names,msg.body)
}