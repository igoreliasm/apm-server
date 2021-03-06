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

// Code generated by "modeldecoder/generator". DO NOT EDIT.

package rumv3

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func (val *metadataRoot) IsSet() bool {
	return val.Metadata.IsSet()
}

func (val *metadataRoot) Reset() {
	val.Metadata.Reset()
}

func (val *metadataRoot) validate() error {
	if err := val.Metadata.validate(); err != nil {
		return errors.Wrapf(err, "m")
	}
	if !val.Metadata.IsSet() {
		return fmt.Errorf("'m' required")
	}
	return nil
}

func (val *metadata) IsSet() bool {
	return len(val.Labels) > 0 || val.Service.IsSet() || val.User.IsSet()
}

func (val *metadata) Reset() {
	for k := range val.Labels {
		delete(val.Labels, k)
	}
	val.Service.Reset()
	val.User.Reset()
}

func (val *metadata) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k, v := range val.Labels {
		if k != "" && !regexpNoDotAsteriskQuote.MatchString(k) {
			return fmt.Errorf("'l': validation rule 'patternKeys(regexpNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'l': validation rule 'maxVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'l': validation rule 'typesVals(string;bool;number)' violated for key %s", k)
		}
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	if !val.Service.IsSet() {
		return fmt.Errorf("'se' required")
	}
	if err := val.User.validate(); err != nil {
		return errors.Wrapf(err, "u")
	}
	return nil
}

func (val *metadataService) IsSet() bool {
	return val.Agent.IsSet() || val.Environment.IsSet() || val.Framework.IsSet() || val.Language.IsSet() || val.Name.IsSet() || val.Runtime.IsSet() || val.Version.IsSet()
}

func (val *metadataService) Reset() {
	val.Agent.Reset()
	val.Environment.Reset()
	val.Framework.Reset()
	val.Language.Reset()
	val.Name.Reset()
	val.Runtime.Reset()
	val.Version.Reset()
}

func (val *metadataService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Agent.validate(); err != nil {
		return errors.Wrapf(err, "a")
	}
	if !val.Agent.IsSet() {
		return fmt.Errorf("'a' required")
	}
	if utf8.RuneCountInString(val.Environment.Val) > 1024 {
		return fmt.Errorf("'en': validation rule 'max(1024)' violated")
	}
	if err := val.Framework.validate(); err != nil {
		return errors.Wrapf(err, "fw")
	}
	if err := val.Language.validate(); err != nil {
		return errors.Wrapf(err, "la")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) < 1 {
		return fmt.Errorf("'n': validation rule 'min(1)' violated")
	}
	if val.Name.Val != "" && !regexpAlphaNumericExt.MatchString(val.Name.Val) {
		return fmt.Errorf("'n': validation rule 'pattern(regexpAlphaNumericExt)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if err := val.Runtime.validate(); err != nil {
		return errors.Wrapf(err, "ru")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *metadataServiceAgent) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceAgent) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceAgent) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) < 1 {
		return fmt.Errorf("'n': validation rule 'min(1)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	if !val.Version.IsSet() {
		return fmt.Errorf("'ve' required")
	}
	return nil
}

func (val *metadataServiceFramework) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceFramework) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceFramework) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *metadataServiceLanguage) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceLanguage) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceLanguage) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *metadataServiceRuntime) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceRuntime) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceRuntime) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	if !val.Version.IsSet() {
		return fmt.Errorf("'ve' required")
	}
	return nil
}

func (val *user) IsSet() bool {
	return val.ID.IsSet() || val.Email.IsSet() || val.Name.IsSet()
}

func (val *user) Reset() {
	val.ID.Reset()
	val.Email.Reset()
	val.Name.Reset()
}

func (val *user) validate() error {
	if !val.IsSet() {
		return nil
	}
	switch t := val.ID.Val.(type) {
	case string:
		if utf8.RuneCountInString(t) > 1024 {
			return fmt.Errorf("'id': validation rule 'max(1024)' violated")
		}
	case int:
	case json.Number:
		if _, err := t.Int64(); err != nil {
			return fmt.Errorf("'id': validation rule 'types(string;int)' violated")
		}
	case nil:
	default:
		return fmt.Errorf("'id': validation rule 'types(string;int)' violated ")
	}
	if utf8.RuneCountInString(val.Email.Val) > 1024 {
		return fmt.Errorf("'em': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'un': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *transactionRoot) IsSet() bool {
	return val.Transaction.IsSet()
}

func (val *transactionRoot) Reset() {
	val.Transaction.Reset()
}

func (val *transactionRoot) validate() error {
	if err := val.Transaction.validate(); err != nil {
		return errors.Wrapf(err, "x")
	}
	if !val.Transaction.IsSet() {
		return fmt.Errorf("'x' required")
	}
	return nil
}

func (val *transaction) IsSet() bool {
	return val.Context.IsSet() || val.Duration.IsSet() || val.ID.IsSet() || val.Marks.IsSet() || val.Name.IsSet() || val.Outcome.IsSet() || val.ParentID.IsSet() || val.Result.IsSet() || val.Sampled.IsSet() || val.SampleRate.IsSet() || val.SpanCount.IsSet() || val.TraceID.IsSet() || val.Type.IsSet() || val.UserExperience.IsSet() || val.Experimental.IsSet()
}

func (val *transaction) Reset() {
	val.Context.Reset()
	val.Duration.Reset()
	val.ID.Reset()
	val.Marks.Reset()
	val.Name.Reset()
	val.Outcome.Reset()
	val.ParentID.Reset()
	val.Result.Reset()
	val.Sampled.Reset()
	val.SampleRate.Reset()
	val.SpanCount.Reset()
	val.TraceID.Reset()
	val.Type.Reset()
	val.UserExperience.Reset()
	val.Experimental.Reset()
}

func (val *transaction) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Context.validate(); err != nil {
		return errors.Wrapf(err, "c")
	}
	if val.Duration.Val < 0 {
		return fmt.Errorf("'d': validation rule 'min(0)' violated")
	}
	if !val.Duration.IsSet() {
		return fmt.Errorf("'d' required")
	}
	if utf8.RuneCountInString(val.ID.Val) > 1024 {
		return fmt.Errorf("'id': validation rule 'max(1024)' violated")
	}
	if !val.ID.IsSet() {
		return fmt.Errorf("'id' required")
	}
	if err := val.Marks.validate(); err != nil {
		return errors.Wrapf(err, "k")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if val.Outcome.Val != "" {
		var matchEnum bool
		for _, s := range enumOutcome {
			if val.Outcome.Val == s {
				matchEnum = true
				break
			}
		}
		if !matchEnum {
			return fmt.Errorf("'o': validation rule 'enum(enumOutcome)' violated")
		}
	}
	if utf8.RuneCountInString(val.ParentID.Val) > 1024 {
		return fmt.Errorf("'pid': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Result.Val) > 1024 {
		return fmt.Errorf("'rt': validation rule 'max(1024)' violated")
	}
	if err := val.SpanCount.validate(); err != nil {
		return errors.Wrapf(err, "yc")
	}
	if !val.SpanCount.IsSet() {
		return fmt.Errorf("'yc' required")
	}
	if utf8.RuneCountInString(val.TraceID.Val) > 1024 {
		return fmt.Errorf("'tid': validation rule 'max(1024)' violated")
	}
	if !val.TraceID.IsSet() {
		return fmt.Errorf("'tid' required")
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'max(1024)' violated")
	}
	if !val.Type.IsSet() {
		return fmt.Errorf("'t' required")
	}
	if err := val.UserExperience.validate(); err != nil {
		return errors.Wrapf(err, "exp")
	}
	return nil
}

func (val *context) IsSet() bool {
	return len(val.Custom) > 0 || val.Page.IsSet() || val.Request.IsSet() || val.Response.IsSet() || val.Service.IsSet() || len(val.Tags) > 0 || val.User.IsSet()
}

func (val *context) Reset() {
	for k := range val.Custom {
		delete(val.Custom, k)
	}
	val.Page.Reset()
	val.Request.Reset()
	val.Response.Reset()
	val.Service.Reset()
	for k := range val.Tags {
		delete(val.Tags, k)
	}
	val.User.Reset()
}

func (val *context) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k := range val.Custom {
		if k != "" && !regexpNoDotAsteriskQuote.MatchString(k) {
			return fmt.Errorf("'cu': validation rule 'patternKeys(regexpNoDotAsteriskQuote)' violated")
		}
	}
	if err := val.Page.validate(); err != nil {
		return errors.Wrapf(err, "p")
	}
	if err := val.Request.validate(); err != nil {
		return errors.Wrapf(err, "q")
	}
	if err := val.Response.validate(); err != nil {
		return errors.Wrapf(err, "r")
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	for k, v := range val.Tags {
		if k != "" && !regexpNoDotAsteriskQuote.MatchString(k) {
			return fmt.Errorf("'g': validation rule 'patternKeys(regexpNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'g': validation rule 'maxVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'g': validation rule 'typesVals(string;bool;number)' violated for key %s", k)
		}
	}
	if err := val.User.validate(); err != nil {
		return errors.Wrapf(err, "u")
	}
	return nil
}

func (val *contextPage) IsSet() bool {
	return val.URL.IsSet() || val.Referer.IsSet()
}

func (val *contextPage) Reset() {
	val.URL.Reset()
	val.Referer.Reset()
}

func (val *contextPage) validate() error {
	if !val.IsSet() {
		return nil
	}
	return nil
}

func (val *contextRequest) IsSet() bool {
	return val.Env.IsSet() || val.Headers.IsSet() || val.HTTPVersion.IsSet() || val.Method.IsSet()
}

func (val *contextRequest) Reset() {
	val.Env.Reset()
	val.Headers.Reset()
	val.HTTPVersion.Reset()
	val.Method.Reset()
}

func (val *contextRequest) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.HTTPVersion.Val) > 1024 {
		return fmt.Errorf("'hve': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Method.Val) > 1024 {
		return fmt.Errorf("'mt': validation rule 'max(1024)' violated")
	}
	if !val.Method.IsSet() {
		return fmt.Errorf("'mt' required")
	}
	return nil
}

func (val *contextResponse) IsSet() bool {
	return val.DecodedBodySize.IsSet() || val.EncodedBodySize.IsSet() || val.Headers.IsSet() || val.StatusCode.IsSet() || val.TransferSize.IsSet()
}

func (val *contextResponse) Reset() {
	val.DecodedBodySize.Reset()
	val.EncodedBodySize.Reset()
	val.Headers.Reset()
	val.StatusCode.Reset()
	val.TransferSize.Reset()
}

func (val *contextResponse) validate() error {
	if !val.IsSet() {
		return nil
	}
	return nil
}

func (val *contextService) IsSet() bool {
	return val.Agent.IsSet() || val.Environment.IsSet() || val.Framework.IsSet() || val.Language.IsSet() || val.Name.IsSet() || val.Runtime.IsSet() || val.Version.IsSet()
}

func (val *contextService) Reset() {
	val.Agent.Reset()
	val.Environment.Reset()
	val.Framework.Reset()
	val.Language.Reset()
	val.Name.Reset()
	val.Runtime.Reset()
	val.Version.Reset()
}

func (val *contextService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Agent.validate(); err != nil {
		return errors.Wrapf(err, "a")
	}
	if utf8.RuneCountInString(val.Environment.Val) > 1024 {
		return fmt.Errorf("'en': validation rule 'max(1024)' violated")
	}
	if err := val.Framework.validate(); err != nil {
		return errors.Wrapf(err, "fw")
	}
	if err := val.Language.validate(); err != nil {
		return errors.Wrapf(err, "la")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if val.Name.Val != "" && !regexpAlphaNumericExt.MatchString(val.Name.Val) {
		return fmt.Errorf("'n': validation rule 'pattern(regexpAlphaNumericExt)' violated")
	}
	if err := val.Runtime.validate(); err != nil {
		return errors.Wrapf(err, "ru")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *contextServiceAgent) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceAgent) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceAgent) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *contextServiceFramework) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceFramework) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceFramework) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *contextServiceLanguage) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceLanguage) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceLanguage) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *contextServiceRuntime) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceRuntime) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceRuntime) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'max(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'max(1024)' violated")
	}
	return nil
}

func (val *transactionMarks) IsSet() bool {
	return len(val.Events) > 0
}

func (val *transactionMarks) Reset() {
	for k := range val.Events {
		delete(val.Events, k)
	}
}

func (val *transactionMarks) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k, v := range val.Events {
		if err := v.validate(); err != nil {
			return errors.Wrapf(err, "events")
		}
		if k != "" && !regexpNoDotAsteriskQuote.MatchString(k) {
			return fmt.Errorf("'events': validation rule 'patternKeys(regexpNoDotAsteriskQuote)' violated")
		}
	}
	return nil
}

func (val *transactionMarkEvents) IsSet() bool {
	return len(val.Measurements) > 0
}

func (val *transactionMarkEvents) Reset() {
	for k := range val.Measurements {
		delete(val.Measurements, k)
	}
}

func (val *transactionMarkEvents) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k := range val.Measurements {
		if k != "" && !regexpNoDotAsteriskQuote.MatchString(k) {
			return fmt.Errorf("'measurements': validation rule 'patternKeys(regexpNoDotAsteriskQuote)' violated")
		}
	}
	return nil
}

func (val *transactionSpanCount) IsSet() bool {
	return val.Dropped.IsSet() || val.Started.IsSet()
}

func (val *transactionSpanCount) Reset() {
	val.Dropped.Reset()
	val.Started.Reset()
}

func (val *transactionSpanCount) validate() error {
	if !val.IsSet() {
		return nil
	}
	if !val.Started.IsSet() {
		return fmt.Errorf("'sd' required")
	}
	return nil
}

func (val *transactionUserExperience) IsSet() bool {
	return val.CumulativeLayoutShift.IsSet() || val.FirstInputDelay.IsSet() || val.TotalBlockingTime.IsSet() || val.Longtask.IsSet()
}

func (val *transactionUserExperience) Reset() {
	val.CumulativeLayoutShift.Reset()
	val.FirstInputDelay.Reset()
	val.TotalBlockingTime.Reset()
	val.Longtask.Reset()
}

func (val *transactionUserExperience) validate() error {
	if !val.IsSet() {
		return nil
	}
	if val.CumulativeLayoutShift.Val < 0 {
		return fmt.Errorf("'cls': validation rule 'min(0)' violated")
	}
	if val.FirstInputDelay.Val < 0 {
		return fmt.Errorf("'fid': validation rule 'min(0)' violated")
	}
	if val.TotalBlockingTime.Val < 0 {
		return fmt.Errorf("'tbt': validation rule 'min(0)' violated")
	}
	if err := val.Longtask.validate(); err != nil {
		return errors.Wrapf(err, "lt")
	}
	return nil
}

func (val *longtaskMetrics) IsSet() bool {
	return val.Count.IsSet() || val.Sum.IsSet() || val.Max.IsSet()
}

func (val *longtaskMetrics) Reset() {
	val.Count.Reset()
	val.Sum.Reset()
	val.Max.Reset()
}

func (val *longtaskMetrics) validate() error {
	if !val.IsSet() {
		return nil
	}
	if val.Count.Val < 0 {
		return fmt.Errorf("'count': validation rule 'min(0)' violated")
	}
	if !val.Count.IsSet() {
		return fmt.Errorf("'count' required")
	}
	if val.Sum.Val < 0 {
		return fmt.Errorf("'sum': validation rule 'min(0)' violated")
	}
	if !val.Sum.IsSet() {
		return fmt.Errorf("'sum' required")
	}
	if val.Max.Val < 0 {
		return fmt.Errorf("'max': validation rule 'min(0)' violated")
	}
	if !val.Max.IsSet() {
		return fmt.Errorf("'max' required")
	}
	return nil
}
