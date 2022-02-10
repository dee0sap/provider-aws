/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1beta1

type FilterNameStringType string

const (
	FilterNameStringType_description FilterNameStringType = "description"
	FilterNameStringType_name FilterNameStringType = "name"
	FilterNameStringType_tag_key FilterNameStringType = "tag-key"
	FilterNameStringType_tag_value FilterNameStringType = "tag-value"
	FilterNameStringType_all FilterNameStringType = "all"
)

type SortOrderType string

const (
	SortOrderType_asc SortOrderType = "asc"
	SortOrderType_desc SortOrderType = "desc"
)