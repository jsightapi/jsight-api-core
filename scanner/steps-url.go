package scanner

import (
	"github.com/jsightapi/jsight-api-core/jerr"
)

func stateU(s *Scanner, c byte) *jerr.JApiError {
	switch c {
	case 'R':
		s.step = stateUR
		return nil
	default:
		return s.japiErrorUnexpectedChar("in keyword URL", "R")
	}
}

func stateUR(s *Scanner, c byte) *jerr.JApiError {
	switch c {
	case 'L':
		s.found(KeywordEnd)
		s.stepStack.Push(stateExpectKeyword)
		s.step = stateParameterOrAnnotation
		return nil
	default:
		return s.japiErrorUnexpectedChar("in keyword URL", "L")
	}
}
