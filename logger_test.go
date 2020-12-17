package logger

import "testing"

func TestGroup(t *testing.T) {
	buf := New(nil)
	l := buf.Log("").(*defaultLog)
	l2 := buf.Log("").(*defaultLog)
	if l.group != l2.group {
		t.Errorf("ERROR: %d != %d for group %s", l.group, l2.group, l.Service())
	} else {
		t.Logf("%s: both have %d", l.Service(), l.group)
	}
	l3 := buf.Log("other").(*defaultLog)
	if l.group == l3.group {
		t.Errorf("ERROR: %s and %s have the same group %d", l.Service(), l3.Service(), l.group)
	} else {
		t.Logf("%s and %s have different groups %d and %d", l.Service(), l3.Service(), l.group, l3.group)
	}
	l4 := l3.Fork().(*defaultLog)
	if l4.group == l3.group {
		t.Errorf("ERROR: %s and %s have the same group %d", l4.Service(), l3.Service(), l4.group)
	} else {
		t.Logf("%s and %s have different groups %d and %d", l4.Service(), l3.Service(), l4.group, l3.group)
	}
}

func expectLines(t *testing.T, count int, lines []*Line) {
	t.Helper()
	if len(lines) != count {
		t.Errorf("ERROR: expected %d lines, not %d", count, len(lines))
	} else {
		t.Logf("Got expected number of lines: %d", count)
	}
}

func TestSeq(t *testing.T) {
	buf := New(nil)
	l := buf.Log("")
	l2 := buf.Log("")
	l.Errorf("0")
	l2.Errorf("1")
	lines := buf.Lines(-1)
	expectLines(t, 2, lines)
	if lines[0].Group != lines[1].Group {
		t.Errorf("ERROR: Expected lines to have the same group, not %d and %d", lines[0].Group, lines[1].Group)
	} else {
		t.Logf("Lines have the same group")
	}
	if lines[0].Seq == lines[1].Seq {
		t.Errorf("ERROR: Lines have the same sequence %d, expected them to be different", lines[0].Seq)
	} else {
		t.Logf("Lines have different sequence numbers %d and %d", lines[0].Seq, lines[1].Seq)
	}
}

func TestBuffer(t *testing.T) {
	buf := New(nil)
	l := buf.Log("")
	if buf.MaxLines() != 1000 {
		t.Errorf("ERROR: Expected MaxLines default to be 1000, not %d", buf.MaxLines())
	} else {
		t.Logf("Got expected MaxLines %d", buf.MaxLines())
	}
	l.Errorf("0")
	l.Errorf("1")
	lines := buf.Lines(-1)
	expectLines(t, 2, lines)
	buf.KeepLines(1)
	if buf.MaxLines() != 1 {
		t.Errorf("ERROR: Expected MaxLines to be 1, not %d", buf.MaxLines())
	} else {
		t.Logf("Got expected MaxLines %d", buf.MaxLines())
	}
	lines = buf.Lines(-1)
	expectLines(t, 1, lines)
	if lines[0].Message != "1" {
		t.Errorf("Expected line message to be `1`, not `%s`", lines[0].Message)
	} else {
		t.Logf("Got expected line message `%s`", lines[0].Message)
	}
	l.Errorf("2")
	lines = buf.Lines(-1)
	expectLines(t, 1, lines)
	if lines[0].Message != "2" {
		t.Errorf("Expected line message to be `2`, not `%s`", lines[0].Message)
	} else {
		t.Logf("Got expected line message `%s`", lines[0].Message)
	}
	buf.KeepLines(2)
	l.Errorf("3")
	l.Errorf("4")
	l.Errorf("5")
	lines = buf.Lines(-1)
	expectLines(t, 2, lines)
	if lines[0].Message != "4" || lines[1].Message != "5" {
		t.Errorf("Expected 4 and 5 for Lines, not `%s` and `%s`", lines[0].Message, lines[1].Message)
	} else {
		t.Logf("Got expected message contents `%s` and `%s`", lines[0].Message, lines[1].Message)
	}
}

func logLevels(l Logger) {
	l.Tracef("0")
	l.Debugf("1")
	l.Infof("2")
	l.Warnf("3")
	l.Errorf("4")
	l.Fatalf("5")
	l.Panicf("6")
	l.Auditf("7")
}

func TestLevels(t *testing.T) {
	lpl := map[Level]int{
		Trace: 8,
		Debug: 7,
		Info:  6,
		Warn:  5,
		Error: 4,
		Fatal: 3,
		Panic: 2,
		Audit: 1,
	}
	for lvl, count := range lpl {
		buf := New(nil)
		l := buf.Log("").(*defaultLog)
		l.SetLevel(lvl)
		logLevels(l)
		lines := buf.Lines(-1)
		t.Logf("Level %s, expect %d lines", lvl, count)
		expectLines(t, count, lines)
	}
}
