package controllers

import "testing"

// test hitting websocket endpoint
// test add old messages
// test handle messages

// test sending messgae to cleints
// test unsafe

// test handle messages
func TestHandleMessages(test *testing.T){

	go HandleMessages()

	// send a message to be recivcved 


}

// test request for room id that does not exist 