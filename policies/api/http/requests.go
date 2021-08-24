// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package http

import (
	"github.com/ns1labs/orb/fleet"
	"github.com/ns1labs/orb/pkg/errors"
	"github.com/ns1labs/orb/pkg/types"
)

const (
	maxLimitSize = 100
	maxNameSize  = 1024
	nameOrder    = "name"
	idOrder      = "id"
	ascDir       = "asc"
	descDir      = "desc"
)

type addPolicyReq struct {
	Name       string         `json:"name"`
	Backend    string         `json:"backend"`
	Policy     types.Metadata `json:"policy,omitempty"`
	Format     string         `json:"format,omitempty"`
	PolicyData string         `json:"policy_data,omitempty"`
	token      string
}

func (req addPolicyReq) validate() error {
	if req.token == "" {
		return errors.ErrUnauthorizedAccess
	}

	if req.Name == "" {
		return errors.ErrMalformedEntity
	}
	if req.Backend == "" {
		return errors.ErrMalformedEntity
	}

	if req.Policy == nil {
		// passing policy data blob in the specified format
		if req.Format == "" || req.PolicyData == "" {
			return errors.ErrMalformedEntity
		}
	} else {
		// policy is in json, verified by the back ends later
		if req.Format != "" || req.PolicyData != "" {
			return errors.ErrMalformedEntity
		}
	}

	_, err := types.NewIdentifier(req.Name)
	if err != nil {
		return errors.Wrap(errors.ErrMalformedEntity, err)
	}

	return nil
}

type viewResourceReq struct {
	token string
	id    string
}

func (req viewResourceReq) validate() error {
	if req.token == "" {
		return errors.ErrUnauthorizedAccess
	}
	if req.id == "" {
		return errors.ErrMalformedEntity
	}
	return nil
}

type addDatasetReq struct {
	Name         string `json:"name"`
	AgentGroupID string `json:"agent_group_id"`
	PolicyID     string `json:"agent_policy_id"`
	SinkID       string `json:"sink_id"`
	token        string
}

func (req addDatasetReq) validate() error {
	if req.token == "" {
		return errors.ErrUnauthorizedAccess
	}

	if req.Name == "" || req.AgentGroupID == "" || req.PolicyID == "" || req.SinkID == "" {
		return errors.ErrMalformedEntity
	}

	_, err := types.NewIdentifier(req.Name)
	if err != nil {
		return errors.Wrap(errors.ErrMalformedEntity, err)
	}

	return nil
}

type listResourcesReq struct {
	token        string
	pageMetadata fleet.PageMetadata
}

func (req *listResourcesReq) validate() error {
	if req.token == "" {
		return errors.ErrUnauthorizedAccess
	}

	if req.pageMetadata.Limit == 0 {
		req.pageMetadata.Limit = defLimit
	}

	if req.pageMetadata.Limit > maxLimitSize {
		return errors.ErrMalformedEntity
	}

	if len(req.pageMetadata.Name) > maxNameSize {
		return errors.ErrMalformedEntity
	}

	if req.pageMetadata.Order != "" &&
		req.pageMetadata.Order != nameOrder && req.pageMetadata.Order != idOrder {
		return errors.ErrMalformedEntity
	}

	if req.pageMetadata.Dir != "" &&
		req.pageMetadata.Dir != ascDir && req.pageMetadata.Dir != descDir {
		return errors.ErrMalformedEntity
	}

	return nil
}
