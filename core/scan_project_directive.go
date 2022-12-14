package core

import (
	"github.com/jsightapi/jsight-api-core/jerr"
)

func (core *JApiCore) processCurrentDirective() *jerr.JApiError {
	if core.currentDirective == nil {
		return nil // directive already processed, i.e. by ContextClose before next keyword or when includes happen
	}

	if je := core.processContext(core.currentDirective, &core.directives); je != nil {
		return je
	}

	core.currentDirective = nil

	return nil
}
