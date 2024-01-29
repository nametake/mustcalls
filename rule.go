package mustcalls

import (
	"fmt"
	"regexp"
	"strings"
)

type Rule struct {
	Funcs              FuncRules `yaml:"funcs"`
	FilePatterns       []string  `yaml:"file_patterns,omitempty"`
	FuncPatterns       []string  `yaml:"func_patterns,omitempty"`
	RecvPatterns       []string  `yaml:"recv_patterns,omitempty"`
	IgnoreFilePatterns []string  `yaml:"ignore_file_patterns,omitempty"`
	IgnoreFuncPatterns []string  `yaml:"ignore_func_patterns,omitempty"`
	IgnoreRecvPatterns []string  `yaml:"ignore_recv_patterns,omitempty"`
}

func (rule *Rule) IsTargetFile(filename string) (bool, error) {
	if len(rule.FilePatterns) == 0 {
		return true, nil
	}
	for _, pattern := range rule.FilePatterns {
		matched, err := regexp.MatchString(pattern, filename)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func (rule *Rule) IsIgnoreFile(fileName string) (bool, error) {
	if len(rule.IgnoreFilePatterns) == 0 {
		return false, nil
	}
	for _, pattern := range rule.IgnoreFilePatterns {
		matched, err := regexp.MatchString(pattern, fileName)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func (rule *Rule) IsTargetFunc(funcName string) (bool, error) {
	if len(rule.FuncPatterns) == 0 {
		return true, nil
	}
	for _, pattern := range rule.FuncPatterns {
		matched, err := regexp.MatchString(pattern, funcName)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func (rule *Rule) IsIgnoreFunc(funcName string) (bool, error) {
	if len(rule.IgnoreFuncPatterns) == 0 {
		return false, nil
	}
	for _, pattern := range rule.IgnoreFuncPatterns {
		matched, err := regexp.MatchString(pattern, funcName)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func (rule *Rule) IsTargetRecv(recvName string) (bool, error) {
	if len(rule.RecvPatterns) == 0 {
		return true, nil
	}
	for _, pattern := range rule.RecvPatterns {
		matched, err := regexp.MatchString(pattern, recvName)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func (rule *Rule) IsIgnoreRecv(recvName string) (bool, error) {
	if len(rule.IgnoreRecvPatterns) == 0 {
		return false, nil
	}
	for _, pattern := range rule.IgnoreRecvPatterns {
		matched, err := regexp.MatchString(pattern, recvName)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

type FuncRule struct {
	Name string `yaml:"name"`
}

func (rule *FuncRule) Match(calledFuncs []*astCalledFunc) bool {
	for _, calledFunc := range calledFuncs {
		if rule.matchName(calledFunc) {
			return true
		}
	}
	return false
}

func (rule *FuncRule) matchName(calledFunc *astCalledFunc) bool {
	return rule.Name == calledFunc.Name
}

type FuncRules []*FuncRule

func (funcRules FuncRules) Match(calledFuncs []*astCalledFunc) FuncRules {
	unmatchedRules := make(FuncRules, 0, len(funcRules))
	for _, funcRule := range funcRules {
		if !funcRule.Match(calledFuncs) {
			unmatchedRules = append(unmatchedRules, funcRule)
		}
	}
	return unmatchedRules
}

func (funcRules FuncRules) ErrorMsg(funcName string) string {
	ruleFuncNames := []string{}

	for _, funcRule := range funcRules {
		ruleFuncNames = append(ruleFuncNames, funcRule.Name)
	}

	return fmt.Sprintf("%s is not calling %s", funcName, strings.Join(ruleFuncNames, ", "))
}
