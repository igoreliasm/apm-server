// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package generator

import (
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func generateNullableInterfaceValidation(w io.Writer, fields []structField, f structField, _ bool) error {
	rules, err := validationRules(f.tag)
	if err != nil {
		return errors.Wrap(err, "nullableInterface")
	}
	for _, rule := range rules {
		switch rule.name {
		case tagMin, tagMax:
			//handled in switch statement for string types
		case tagRequired:
			ruleNullableRequired(w, f)
		case tagTypes:
			nullableInterfaceRuleTypes(w, f, rules, rule)
		default:
			errors.Wrap(errUnhandledTagRule(rule), "nullableInterface")
		}
	}
	return nil
}

func nullableInterfaceRuleTypes(w io.Writer, f structField, rules []validationRule, rule validationRule) {
	var isRequired bool
	var maxRule validationRule
	for _, r := range rules {
		if r.name == tagRequired {
			isRequired = true
			continue
		}
		if r.name == tagMax {
			maxRule = r
			continue
		}
	}

	var switchStmt string
	if maxRule != (validationRule{}) {
		switchStmt = `switch t := val.%s.Val.(type){`
	} else {
		switchStmt = `switch val.%s.Val.(type){`
	}
	fmt.Fprintf(w, switchStmt, f.Name())

	for _, typ := range strings.Split(rule.value, ";") {
		switch typ {
		case "int":
			fmt.Fprintf(w, `
case int:
case json.Number:
	if _, err := t.Int64(); err != nil{
		return fmt.Errorf("'%s': validation rule '%s(%s)' violated")
	}
`[1:], jsonName(f), rule.name, rule.value)
		case "string":
			fmt.Fprintf(w, `
case %s:
`[1:], typ)
			if maxRule != (validationRule{}) {
				fmt.Fprintf(w, `
if utf8.RuneCountInString(t) %s %s{
	return fmt.Errorf("'%s': validation rule '%s(%s)' violated")
}
`[1:], ruleMinMaxOperator(maxRule.name), maxRule.value, jsonName(f), maxRule.name, maxRule.value)
			}
		case "interface":
			fmt.Fprint(w, `
case interface{}:
`[1:])
		case "map[string]interface":
			fmt.Fprint(w, `
case map[string]interface{}:
`[1:])
		default:
			fmt.Fprintf(w, `
case %s:
`[1:], typ)
		}
	}
	if !isRequired {
		fmt.Fprintf(w, `
case nil:
`[1:])
	}
	fmt.Fprintf(w, `
default:
	return fmt.Errorf("'%s': validation rule '%s(%s)' violated ")
}
`[1:], jsonName(f), rule.name, rule.value)
}
