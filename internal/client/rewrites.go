/*
 * SPDX-FileCopyrightText: The terraform-provider-migadu Authors
 * SPDX-License-Identifier: 0BSD
 */

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/idna"
	"net/http"
)

type Rewrites struct {
	Rewrites []Rewrite `json:"rewrites"`
}

type Rewrite struct {
	DomainName    string   `json:"domain_name"`
	Name          string   `json:"name"`
	LocalPartRule string   `json:"local_part_rule"`
	OrderNum      int64    `json:"order_num"`
	Destinations  []string `json:"destinations"`
}

// GetRewrites - Returns rewrite rules for a single domain
func (c *MigaduClient) GetRewrites(ctx context.Context, domain string) (*Rewrites, error) {
	ascii, err := idna.ToASCII(domain)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/domains/%s/rewrites", c.Endpoint, ascii)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rewrites := Rewrites{}
	err = json.Unmarshal(body, &rewrites)
	if err != nil {
		return nil, err
	}

	return &rewrites, nil
}

// GetRewrite - Returns specific rewrite rule
func (c *MigaduClient) GetRewrite(ctx context.Context, domain string, slug string) (*Rewrite, error) {
	ascii, err := idna.ToASCII(domain)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/domains/%s/rewrites/%s", c.Endpoint, ascii, slug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rewrite := Rewrite{}
	err = json.Unmarshal(body, &rewrite)
	if err != nil {
		return nil, err
	}

	return &rewrite, nil
}
