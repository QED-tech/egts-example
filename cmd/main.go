package main

import (
	"egts-demo/internal/message"
	"egts-demo/internal/network"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/kuznetsovin/egts-protocol/libs/egts"
	"log"
	"time"
)

func main() {
	address := flag.String("address", "localhost:9090", "EGTS server address")
	idleTimeout := flag.Duration("idle_timeout", time.Minute, "Idle timeout for connection")
	flag.Parse()

	client, err := network.NewTCPClient(*address, *idleTimeout)
	if err != nil {
		log.Fatal(err)
	}

	sendAuthMessage(client)

	sendTelematicDataMessage(client, 2)
	sendTelematicDataMessage(client, 3)
}

func sendAuthMessage(client *network.Client) {
	authMessage, err := message.CreateAuthMessage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PID: %d, Encoded package: %s\n", message.AuthMessagePID, hex.EncodeToString(authMessage))

	ackAuth, err := client.Send(authMessage)
	if err != nil {
		log.Fatal(err)
	}

	responsePackage := egts.Package{}
	_, err = responsePackage.Decode(ackAuth)
	if err != nil {
		log.Fatal(err)
	}

	b, err := responsePackage.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PID: %d, Response: %s\n", message.AuthMessagePID, b)
}

func sendTelematicDataMessage(client *network.Client, pid int) {
	telematicDataMessage, err := message.CreateTelematicDataMessage(pid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PID: %d, Encoded package: %s\n", pid, hex.EncodeToString(telematicDataMessage))

	ackTelematicData, err := client.Send(telematicDataMessage)
	if err != nil {
		log.Fatal(err)
	}

	responsePackage := egts.Package{}
	_, err = responsePackage.Decode(ackTelematicData)
	if err != nil {
		log.Fatal(err)
	}

	b, err := responsePackage.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PID: %d, Response: %s\n", pid, b)
}
