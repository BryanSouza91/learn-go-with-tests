package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()
		want := `3
				2
				1
				Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
