package otus_lesson6

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestLogOtusEvent(t *testing.T) {
	var buffer strings.Builder

	if LogOtusEvent(nil, &buffer) == nil || LogOtusEvent(&HwAccepted{}, nil) == nil {
		t.Fatal("no nil check")
	}

	currentDate := time.Now().Format("2006-01-02")
	t.Logf("Current date: %s\n", currentDate)

	t.Log("\n")
	t.Log("Testing HwAccepted:")

	msg1 := &HwAccepted{
		Id: 1,
		Grade: 2,
	}

	expectedString := fmt.Sprintf("%s accepted %d %d", currentDate, msg1.Id, msg1.Grade)
	testWriteEvent(t, expectedString, &buffer, msg1)

	t.Log("\n")
	t.Log("Testing HwSubmitted:")

	msg2 := &HwSubmitted{
		Id: 1,
		Code: "code",
		Comment: "comment",
	}

	buffer.Reset()
	expectedString = fmt.Sprintf("%s submitted %d %s %s", currentDate, msg2.Id, msg2.Code, msg2.Comment)
	testWriteEvent(t, expectedString, &buffer, msg2)
}

func testWriteEvent(t *testing.T, expected string, buffer *strings.Builder, event OtusEvent) {
	err := LogOtusEvent(event, buffer)

	if err != nil {
		t.Fatal(err)
	}


	actualString := buffer.String()

	t.Logf("Expected string: \"%s\".\n", expected)
	t.Logf("Actual string: \"%s\".\n", actualString)

	if expected != actualString {
		t.Fatalf("Incorrect result. Expected = \"%s\", actual = \"%s\".", expected, actualString)
	}
}
