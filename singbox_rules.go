package common

import (
	"strings"
	"time"

	"github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badoption"
)

// ToOptions converts SmartRoutingRules to sing-box options, including required outbounds, rules, and rulesets.
func (srs SmartRoutingRules) ToOptions(urlTestInterval, urlTestIdleTimeout time.Duration) ([]option.Outbound, []option.Rule, []option.RuleSet) {
	outbounds := []option.Outbound{}
	rules := []option.Rule{}
	rulesets := []option.RuleSet{}
	for _, sr := range srs {
		rs := sr.RuleSets.ToOptions()
		tags := sr.RuleSets.Tags()
		detour := sr.Category
		if len(sr.Outbounds) == 1 && sr.Outbounds[0] == "direct" {
			detour = "direct"
		}
		rule := option.Rule{
			Type: constant.RuleTypeDefault,
			DefaultOptions: option.DefaultRule{
				RawDefaultRule: option.RawDefaultRule{
					RuleSet: tags,
				},
				RuleAction: option.RuleAction{
					Action: constant.RuleActionTypeRoute,
					RouteOptions: option.RouteActionOptions{
						Outbound: detour,
					},
				},
			},
		}
		rules = append(rules, rule)
		rulesets = append(rulesets, rs...)

		if detour != "direct" {
			outbounds = append(outbounds, option.Outbound{
				Type: constant.TypeURLTest,
				Tag:  "sr-" + sr.Category,
				Options: &option.URLTestOutboundOptions{
					Outbounds:   sr.Outbounds,
					URL:         "https://google.com/generate_204",
					Interval:    badoption.Duration(urlTestInterval),
					IdleTimeout: badoption.Duration(urlTestIdleTimeout),
				},
			})
		}
	}
	return outbounds, rules, rulesets
}

// ToOptions converts AdBlockRules to sing-box options, including the rule and rulesets.
func (ab AdBlockRules) ToOptions() (option.Rule, []option.RuleSet) {
	rulesets := (RuleSets)(ab).ToOptions()
	tags := (RuleSets)(ab).Tags()
	rule := option.Rule{
		Type: constant.RuleTypeDefault,
		DefaultOptions: option.DefaultRule{
			RawDefaultRule: option.RawDefaultRule{
				RuleSet: tags,
			},
			RuleAction: option.RuleAction{
				Action: constant.RuleActionTypeReject,
			},
		},
	}
	return rule, rulesets
}

// Tags returns the tags of all RuleSets.
func (rs RuleSets) Tags() []string {
	tags := []string{}
	for _, rule := range rs {
		tags = append(tags, rule.Tag)
	}
	return tags
}

// ToOptions converts RuleSets to sing-box RuleSet options.
func (rs RuleSets) ToOptions() []option.RuleSet {
	rulesets := []option.RuleSet{}
	for _, rule := range rs {
		format := rule.Format
		if format == "" {
			if strings.HasSuffix(rule.URL, ".srs") {
				format = constant.RuleSetFormatBinary
			} else {
				format = constant.RuleSetFormatSource
			}
		}
		detour := rule.DownloadDetour
		if detour == "" {
			detour = "direct"
		}
		rulesets = append(rulesets, option.RuleSet{
			Type:   constant.RuleSetTypeRemote,
			Tag:    rule.Tag,
			Format: format,
			RemoteOptions: option.RemoteRuleSet{
				URL:            rule.URL,
				DownloadDetour: detour,
				UpdateInterval: badoption.Duration(24 * time.Hour),
			},
		})
	}
	return rulesets
}
