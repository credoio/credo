package megadodo

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/Sirupsen/logrus.v0"
)

type Log struct {
	Facility  int
	Severity  int
	Origin    string
	Timestamp string
	Message   string
}

var LR = logrus.New()

func init() {
	LR.Formatter = new(logrus.TextFormatter)
	LR.Level = logrus.DebugLevel
}

func ParseLog(logStr string) (*Log, error) {
	log := []rune(logStr)
	LR.WithFields(logrus.Fields{
		"log": logStr,
		"len": len(log),
	}).Debug("start")
	if len(log) < 1 {
		LR.Debug("input too short")
		return nil, fmt.Errorf("input too short")
	}
	if log[0] != '<' {
		return nil, fmt.Errorf("invalid starting char")
	}
	numDigits := 0
	bracketCloseIndex := 0
	for i := 1; i < len(log); i++ {
		if log[i] == '>' {
			bracketCloseIndex = i
			break
		}
		if !unicode.IsDigit(log[i]) {
			return nil, fmt.Errorf("invalid number in facility")
		}
		numDigits++
	}
	if bracketCloseIndex == 0 {
		return nil, fmt.Errorf("missing bracket")
	}
	if numDigits == 0 {
		return nil, fmt.Errorf("missing facility")
	}
	LR.WithFields(logrus.Fields{
		"numDigits": numDigits,
	}).Debug("digit check")
	numStr := string(log[1 : numDigits+1])
	val, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, fmt.Errorf("non numeric facility")
	}
	msg := logStr[bracketCloseIndex+1:]
	if len(msg) == 0 {
		return nil, fmt.Errorf("missing message")
	}
	fac := val / 8
	if fac < 0 || fac > 23 {
		return nil, fmt.Errorf("facility out of range")
	}
	sev := val - 8*fac
	if sev < 0 || sev > 7 {
		return nil, fmt.Errorf("severity out of range")
	}
	flds := strings.Fields(msg)
	if len(flds) < 6 {
		return nil, fmt.Errorf("not enough fields in msg")
	}
	ts := flds[0] + " " + flds[1] + " " + flds[2]
	host := flds[3]
	tag := flds[4]
	payload := flds[5:]
	if strings.HasSuffix(host, ":") {
		tag = host
		host = "?"
		payload = flds[4:]
	}
	LR.WithFields(logrus.Fields{
		"facsevval": val,
		"fac":       fac,
		"sev":       sev,
		"ts":        ts,
		"host":      host,
		"tag":       tag,
		"payload":   payload,
		"msg":       msg,
	}).Debug("facility/severity")
	return &Log{Facility: fac, Severity: sev}, nil
}
