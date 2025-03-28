package vstr_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestEqual(t *testing.T) {
	vtest.Equal(t, true, vstr.Equal("a", "a"))
	vtest.Equal(t, false, vstr.Equal("a", "a "))
}

func TestRepeat(t *testing.T) {
	vtest.Equal(t, "a", vstr.Repeat("a", 1))
	vtest.Equal(t, "aa", vstr.Repeat("a", 2))
	vtest.Equal(t, "abcabc", vstr.Repeat("abc", 2))
}

func TestS(t *testing.T) {
	vtest.Equal(t, "t", vstr.S("%s", "t"))
	vtest.Equal(t, "t", vstr.S("%v", "t"))
}

func TestStrLimit(t *testing.T) {
	vtest.Equal(t, "ab...", vstr.StrLimit("abcdef123", 2))
	vtest.Equal(t, "ab---", vstr.StrLimit("abcdef123", 2, "---"))
	vtest.Equal(t, "...", vstr.StrLimit("abcdef123", 0))
	vtest.Equal(t, "", vstr.StrLimit("abcdef123", 0, ""))
	vtest.Equal(t, "abcdef123", vstr.StrLimit("abcdef123", 100))
	vtest.Equal(t, "abc中国...", vstr.StrLimit("abc中国(%123", 5))
}

func TestSubStr(t *testing.T) {
	vtest.Equal(t, "cdef123", vstr.SubStr("abcdef123", 2))
	vtest.Equal(t, "3", vstr.SubStr("abcdef123", -1))
	vtest.Equal(t, "23", vstr.SubStr("abcdef123", -2))
	vtest.Equal(t, "abcdef123", vstr.SubStr("abcdef123", 0))
	vtest.Equal(t, "abc", vstr.SubStr("abcdef123", 0, 3))
	vtest.Equal(t, "abc中国", vstr.SubStr("abc中国(%123", 0, 5))
	vtest.Equal(t, "c中国(%", vstr.SubStr("abc中国(%123", 2, 5))
}

func TestTrim(t *testing.T) {
	vtest.Equal(t, "abcdef123", vstr.Trim("    abcdef123    "))
	vtest.Equal(t, "abcdef123", vstr.Trim("    abcdef123"))
	vtest.Equal(t, "abcdef123", vstr.Trim("abcdef123    "))
}
