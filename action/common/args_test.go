package common

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var argsTestCases = []struct {
	input string
	args  []string
	valid bool
}{
	{
		`-a xxx -b yyy -c ' { "a": "b", "c": "d" }' -d -e`,
		[]string{"-a", "xxx", "-b", "yyy", "-c", `{ "a": "b", "c": "d" }`, "-d", "-e"},
		true,
	},
	{
		`-a xxx -b yyy \
-c \
' { \
    "a": "b", \
    "c": "d" \
}' \
-d \
-e`,
		[]string{"-a", "xxx", "-b", "yyy", "-c", `{  "a": "b",  "c": "d"  }`, "-d", "-e"},
		true,
	},
	{
		`-a xxx -b yyy
-c \
' { \
    "a": "b", \
    "c": "d" \
}' \
-d \
-e`,
		[]string{"-a", "xxx", "-b", "yyy"},
		true,
	},
	{
		`-a \
' { \
    "a": "b" \
-b`,
		[]string{},
		false,
	},
}

var dictArgTestCases = []struct {
	dataStr string
	data    map[string]string
	valid   bool
}{
	{
		`{"a": "b", "c": "d"}`,
		map[string]string{"a": "b", "c": "d"},
		true,
	},
	{
		`{"a": "b", "c"}`,
		map[string]string{},
		false,
	},
	{
		`["a", "b", "c", "d"]`,
		map[string]string{},
		false,
	},
	{
		`{ VALID STRING }`,
		map[string]string{},
		false,
	},
}

var arrayArgTestCases = []struct {
	dataStr string
	data    []string
	valid   bool
}{
	{
		`["a", "b", "c", "d"]`,
		[]string{"a", "b", "c", "d"},
		true,
	},
	{
		`{"a": "b", "c": "d"}`,
		[]string{},
		false,
	},
	{
		`{"a": "b", "c"}`,
		[]string{},
		false,
	},

	{
		`{ VALID STRING }`,
		[]string{},
		false,
	},
}

func TestReadArgs(t *testing.T) {
	var stdin bytes.Buffer
	for _, testCase := range argsTestCases {
		stdin.Reset()
		stdin.Write([]byte(testCase.input))
		if !testCase.valid {
			assert.NotNil(t, ReadArgs(&stdin))
			continue
		}
		assert.Nil(t, ReadArgs(&stdin))
		assert.Equal(t, len(os.Args), len(testCase.args))
		for i := 0; i < len(os.Args); i++ {
			assert.Equal(t, os.Args[i], testCase.args[i])
		}
	}
}

func TestParseDictArg(t *testing.T) {
	for _, testCase := range dictArgTestCases {
		if !testCase.valid {
			_, err := ParseDictArg(testCase.dataStr)
			assert.NotNil(t, err)
			continue
		}
		data, err := ParseDictArg(testCase.dataStr)
		assert.Nil(t, err)
		assert.Equal(t, len(data), len(testCase.data))
		for key, value := range data {
			assert.Equal(t, value, testCase.data[key])
		}
	}
}

func TestParseArrayArg(t *testing.T) {
	for _, testCase := range arrayArgTestCases {
		if !testCase.valid {
			_, err := ParseArrayArg(testCase.dataStr)
			assert.NotNil(t, err)
			continue
		}
		data, err := ParseArrayArg(testCase.dataStr)
		assert.Nil(t, err)
		assert.Equal(t, len(data), len(testCase.data))
		for i := 0; i < len(data); i++ {
			assert.Equal(t, data[i], testCase.data[i])
		}
	}
}
