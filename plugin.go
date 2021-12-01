// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"github.com/jptosso/coraza-waf/v2"
	"github.com/jptosso/coraza-waf/v2/operators"
)

type sqli struct {
}

func (o *sqli) Init(data string) error {
	return nil
}

func (o *sqli) Evaluate(tx *coraza.Transaction, value string) bool {
	res, capture := isSQLi(value)
	tx.CaptureField(1, capture)
	return res
}

type xss struct {
}

func (o *xss) Init(data string) error {
	return nil
}

func (o *xss) Evaluate(tx *coraza.Transaction, value string) bool {
	return isXSS(value)
}

func init() {
	operators.RegisterPlugin("detectSQLi", func() coraza.RuleOperator { return new(sqli) })
	operators.RegisterPlugin("detectXSS", func() coraza.RuleOperator { return new(xss) })
}

var (
	_ coraza.RuleOperator = &sqli{}
	_ coraza.RuleOperator = &xss{}
)
