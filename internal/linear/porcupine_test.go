package linear

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type registerInput struct {
	op    bool // false = put, true = get
	value int
}

// a sequential specification of a register
var registerModel = Model{
	Init: func() interface{} {
		return 0
	},
	// step function: takes a state, input, and output, and returns whether it
	// was a legal operation, along with a new state
	Step: func(state, input, output interface{}) (bool, interface{}) {
		regInput := input.(registerInput)
		if regInput.op == false {
			return true, regInput.value // always ok to execute a put
		} else {
			readCorrectValue := output == state
			return readCorrectValue, state // state is unchanged
		}
	},
	DescribeOperation: func(input, output interface{}) string {
		inp := input.(registerInput)
		switch inp.op {
		case true:
			return fmt.Sprintf("get() -> '%d'", output.(int))
		case false:
			return fmt.Sprintf("put('%d')", inp.value)
		}
		return "<invalid>" // unreachable
	},
}

func TestRegisterModel(t *testing.T) {
	// examples taken from http://nil.csail.mit.edu/6.824/2017/quizzes/q2-17-ans.pdf
	// section VII

	ops := []Operation{
		{0, registerInput{false, 100}, 0, 0, 100},
		{1, registerInput{true, 0}, 25, 100, 75},
		{2, registerInput{true, 0}, 30, 0, 60},
	}
	res := CheckOperations(registerModel, ops)
	if res != true {
		t.Fatal("expected operations to be linearizable")
	}

	// same example as above, but with Event
	events := []Event{
		{0, CallEvent, registerInput{false, 100}, 0},
		{1, CallEvent, registerInput{true, 0}, 1},
		{2, CallEvent, registerInput{true, 0}, 2},
		{2, ReturnEvent, 0, 2},
		{1, ReturnEvent, 100, 1},
		{0, ReturnEvent, 0, 0},
	}
	res = CheckEvents(registerModel, events)
	if res != true {
		t.Fatal("expected operations to be linearizable")
	}

	ops = []Operation{
		{0, registerInput{false, 200}, 0, 0, 100},
		{1, registerInput{true, 0}, 10, 200, 30},
		{2, registerInput{true, 0}, 40, 0, 90},
	}
	res = CheckOperations(registerModel, ops)
	if res != false {
		t.Fatal("expected operations to not be linearizable")
	}

	// same example as above, but with Event
	events = []Event{
		{0, CallEvent, registerInput{false, 200}, 0},
		{1, CallEvent, registerInput{true, 0}, 1},
		{1, ReturnEvent, 200, 1},
		{2, CallEvent, registerInput{true, 0}, 2},
		{2, ReturnEvent, 0, 2},
		{0, ReturnEvent, 0, 0},
	}
	res = CheckEvents(registerModel, events)
	if res != false {
		t.Fatal("expected operations to not be linearizable")
	}
}

func TestZeroDuration(t *testing.T) {
	ops := []Operation{
		{0, registerInput{false, 100}, 0, 0, 100},
		{1, registerInput{true, 0}, 25, 100, 75},
		{2, registerInput{true, 0}, 30, 0, 30},
		{3, registerInput{true, 0}, 30, 0, 30},
	}
	res, info := CheckOperationsVerbose(registerModel, ops, 0)
	if res != Ok {
		t.Fatal("expected operations to be linearizable")
	}

	visualizeTempFile(t, registerModel, info)

	ops = []Operation{
		{0, registerInput{false, 200}, 0, 0, 100},
		{1, registerInput{true, 0}, 10, 200, 10},
		{2, registerInput{true, 0}, 10, 200, 10},
		{3, registerInput{true, 0}, 40, 0, 90},
	}
	res, _ = CheckOperationsVerbose(registerModel, ops, 0)
	if res != Illegal {
		t.Fatal("expected operations to not be linearizable")
	}
}

func TestSetModel(t *testing.T) {

	// Set Model is from Jepsen/Knossos Set.
	// A set supports add and read operations, and we must ensure that
	// each read can't read duplicated or unknown values from the set

	// inputs
	type setInput struct {
		op    bool // false = read, true = write
		value int
	}

	// outputs
	type setOutput struct {
		values  []int // read
		unknown bool  // read
	}

	setModel := Model{
		Init: func() interface{} { return []int{} },
		Step: func(state interface{}, input interface{}, output interface{}) (bool, interface{}) {
			st := state.([]int)
			inp := input.(setInput)
			out := output.(setOutput)

			if inp.op == true {
				// always returns true for write
				index := sort.SearchInts(st, inp.value)
				if index >= len(st) || st[index] != inp.value {
					// value not in the set
					st = append(st, inp.value)
					sort.Ints(st)
				}
				return true, st
			}

			sort.Ints(out.values)
			return out.unknown || reflect.DeepEqual(st, out.values), out.values
		},
		Equal: func(state1, state2 interface{}) bool {
			return reflect.DeepEqual(state1, state2)
		},
	}

	events := []Event{
		{0, CallEvent, setInput{true, 100}, 0},
		{1, CallEvent, setInput{true, 0}, 1},
		{2, CallEvent, setInput{false, 0}, 2},
		{2, ReturnEvent, setOutput{[]int{100}, false}, 2},
		{1, ReturnEvent, setOutput{}, 1},
		{0, ReturnEvent, setOutput{}, 0},
	}
	res := CheckEvents(setModel, events)
	if res != true {
		t.Fatal("expected operations to be linearizable")
	}

	events = []Event{
		{0, CallEvent, setInput{true, 100}, 0},
		{1, CallEvent, setInput{true, 110}, 1},
		{2, CallEvent, setInput{false, 0}, 2},
		{2, ReturnEvent, setOutput{[]int{100, 110}, false}, 2},
		{1, ReturnEvent, setOutput{}, 1},
		{0, ReturnEvent, setOutput{}, 0},
	}
	res = CheckEvents(setModel, events)
	if res != true {
		t.Fatal("expected operations to be linearizable")
	}

	events = []Event{
		{0, CallEvent, setInput{true, 100}, 0},
		{1, CallEvent, setInput{true, 110}, 1},
		{2, CallEvent, setInput{false, 0}, 2},
		{2, ReturnEvent, setOutput{[]int{}, true}, 2},
		{1, ReturnEvent, setOutput{}, 1},
		{0, ReturnEvent, setOutput{}, 0},
	}
	res = CheckEvents(setModel, events)
	if res != true {
		t.Fatal("expected operations to be linearizable")
	}

	events = []Event{
		{0, CallEvent, setInput{true, 100}, 0},
		{1, CallEvent, setInput{true, 110}, 1},
		{2, CallEvent, setInput{false, 0}, 2},
		{2, ReturnEvent, setOutput{[]int{100, 100, 110}, false}, 2},
		{1, ReturnEvent, setOutput{}, 1},
		{0, ReturnEvent, setOutput{}, 0},
	}
	res = CheckEvents(setModel, events)
	if res == true {
		t.Fatal("expected operations not to be linearizable")
	}
}
