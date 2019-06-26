package otus_lesson6

import (
	"errors"
	"fmt"
	"io"
	"time"
)

var (
	ErrEventIsNil = errors.New("event is nil")
	ErrWriterIsNil = errors.New("writer is nil")
)

type HwAccepted struct {
	Id int
	Grade int
}

func (e *HwAccepted) LogEntry() string {
	return fmt.Sprintf("accepted %d %d", e.Id, e.Grade)
}

type HwSubmitted struct {
	Id int
	Code string
	Comment string
}

func (e *HwSubmitted) LogEntry() string {
	return fmt.Sprintf("submitted %d %s %s", e.Id, e.Code, e.Comment)
}

type OtusEvent interface {
	LogEntry() string
}

func LogOtusEvent(e OtusEvent, w io.Writer) error {
	if err := checkForNil(e, w); err != nil {
		return err
	}

	_, err := fmt.Fprintf(w, "%s %s", time.Now().Format("2006-01-02"), e.LogEntry())
	return err
}

func checkForNil(event OtusEvent, writer io.Writer) error {
	if event == nil {
		return ErrEventIsNil
	}

	if writer == nil {
		return ErrWriterIsNil
	}

	return nil
}
