package linear

import (
	"io/ioutil"
	"testing"
)

func visualizeTempFile(t *testing.T, model Model, info linearizationInfo) {
	file, err := ioutil.TempFile("", "*.html")
	if err != nil {
		t.Fatalf("failed to create temp file")
	}
	err = Visualize(model, info, file)
	if err != nil {
		t.Fatalf("visualization failed")
	}
	t.Logf("wrote visualization to %s", file.Name())
}

func TestRegisterModelReadme(t *testing.T) {
	// basically the code from the README

	events := []Event{
		// C0: Write(100)
		{Kind: CallEvent, Value: registerInput{false, 100}, Id: 0, ClientId: 0},
		// C1: Read()
		{Kind: CallEvent, Value: registerInput{true, 0}, Id: 1, ClientId: 1},
		// C2: Read()
		{Kind: CallEvent, Value: registerInput{true, 0}, Id: 2, ClientId: 2},
		// C2: Completed Read -> 0
		{Kind: ReturnEvent, Value: 0, Id: 2, ClientId: 2},
		// C1: Completed Read -> 100
		{Kind: ReturnEvent, Value: 100, Id: 1, ClientId: 1},
		// C0: Completed Write
		{Kind: ReturnEvent, Value: 0, Id: 0, ClientId: 0},
	}

	res, info := CheckEventsVerbose(registerModel, events, 0)
	// returns true

	if res != Ok {
		t.Fatal("expected operations to be linearizable")
	}

	visualizeTempFile(t, registerModel, info)

	events = []Event{
		// C0: Write(200)
		{Kind: CallEvent, Value: registerInput{false, 200}, Id: 0, ClientId: 0},
		// C1: Read()
		{Kind: CallEvent, Value: registerInput{true, 0}, Id: 1, ClientId: 1},
		// C1: Completed Read -> 200
		{Kind: ReturnEvent, Value: 200, Id: 1, ClientId: 1},
		// C2: Read()
		{Kind: CallEvent, Value: registerInput{true, 0}, Id: 2, ClientId: 2},
		// C2: Completed Read -> 0
		{Kind: ReturnEvent, Value: 0, Id: 2, ClientId: 2},
		// C0: Completed Write
		{Kind: ReturnEvent, Value: 0, Id: 0, ClientId: 0},
	}

	res, info = CheckEventsVerbose(registerModel, events, 0)
	// returns false

	if res != Illegal {
		t.Fatal("expected operations not to be linearizable")
	}

	visualizeTempFile(t, registerModel, info)
}
