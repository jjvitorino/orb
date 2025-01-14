// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package policies

import (
	"context"
	"fmt"
	"github.com/ns1labs/orb/pkg/errors"
	"github.com/ns1labs/orb/policies/backend"
)

var (
	ErrCreatePolicy       = errors.New("failed to create policy")
	ErrCreateDataset      = errors.New("failed to create dataset")
	ErrInactivateDataset  = errors.New("failed to inactivate dataset")
	ErrUpdateEntity       = errors.New("failed to update entity")
	ErrMalformedEntity    = errors.New("malformed entity")
	ErrNotFound           = errors.New("non-existent entity")
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")
)

func (s policiesService) ListPolicies(ctx context.Context, token string, pm PageMetadata) (Page, error) {
	ownerID, err := s.identify(token)
	if err != nil {
		return Page{}, err
	}
	return s.repo.RetrieveAll(ctx, ownerID, pm)
}

func (s policiesService) ListPoliciesByGroupIDInternal(ctx context.Context, groupIDs []string, ownerID string) ([]Policy, error) {
	if len(groupIDs) == 0 || ownerID == "" {
		return nil, ErrMalformedEntity
	}
	return s.repo.RetrievePoliciesByGroupID(ctx, groupIDs, ownerID)
}

func (s policiesService) ViewPolicyByIDInternal(ctx context.Context, policyID string, ownerID string) (Policy, error) {
	if policyID == "" || ownerID == "" {
		return Policy{}, ErrMalformedEntity
	}
	return s.repo.RetrievePolicyByID(ctx, policyID, ownerID)
}

func (s policiesService) AddDataset(ctx context.Context, token string, d Dataset) (Dataset, error) {
	mfOwnerID, err := s.identify(token)
	if err != nil {
		return Dataset{}, err
	}

	d.MFOwnerID = mfOwnerID

	id, err := s.repo.SaveDataset(ctx, d)
	if err != nil {
		return Dataset{}, errors.Wrap(ErrCreateDataset, err)
	}
	d.ID = id
	return d, nil
}

func (s policiesService) InactivateDatasetByGroupID(ctx context.Context, groupID string, token string) error {
	ownerID, err := s.identify(token)
	if err != nil {
		return err
	}

	if groupID == "" {
		return ErrMalformedEntity
	}
	return s.repo.InactivateDatasetByGroupID(ctx, groupID, ownerID)
}

func (s policiesService) AddPolicy(ctx context.Context, token string, p Policy, format string, policyData string) (Policy, error) {

	mfOwnerID, err := s.identify(token)
	if err != nil {
		return Policy{}, err
	}

	if !backend.HaveBackend(p.Backend) {
		return Policy{}, errors.Wrap(ErrCreatePolicy, errors.New(fmt.Sprintf("unsupported backend: '%s'", p.Backend)))
	}

	if p.Policy == nil {
		// if not already in json, make sure the back end can convert it
		if !backend.GetBackend(p.Backend).SupportsFormat(format) {
			return Policy{}, errors.Wrap(ErrCreatePolicy,
				errors.New(fmt.Sprintf("unsupported policy format '%s' for given backend '%s'", format, p.Backend)))
		}

		p.Policy, err = backend.GetBackend(p.Backend).ConvertFromFormat(format, policyData)
		if err != nil {
			return Policy{}, errors.Wrap(ErrCreatePolicy, err)
		}
	}

	err = backend.GetBackend(p.Backend).Validate(p.Policy)
	if err != nil {
		return Policy{}, errors.Wrap(ErrCreatePolicy, err)
	}

	p.MFOwnerID = mfOwnerID

	id, err := s.repo.SavePolicy(ctx, p)
	if err != nil {
		return Policy{}, errors.Wrap(ErrCreatePolicy, err)
	}
	p.ID = id
	return p, nil
}

func (s policiesService) ViewPolicyByID(ctx context.Context, token string, policyID string) (Policy, error) {
	ownerID, err := s.identify(token)
	if err != nil {
		return Policy{}, err
	}

	res, err := s.repo.RetrievePolicyByID(ctx, policyID, ownerID)
	if err != nil {
		return Policy{}, err
	}
	return res, nil
}
