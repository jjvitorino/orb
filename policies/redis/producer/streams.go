// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package producer

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ns1labs/orb/policies"
	"go.uber.org/zap"
)

const (
	streamID  = "orb.policies"
	streamLen = 1000
)

var _ policies.Service = (*eventStore)(nil)

type eventStore struct {
	svc    policies.Service
	client *redis.Client
	logger *zap.Logger
}

func (e eventStore) AddPolicy(ctx context.Context, token string, p policies.Policy, format string, policyData string) (policies.Policy, error) {
	return e.svc.AddPolicy(ctx, token, p, format, policyData)
}

func (e eventStore) ViewPolicyByID(ctx context.Context, token string, policyID string) (policies.Policy, error) {
	return e.svc.ViewPolicyByID(ctx, token, policyID)
}

func (e eventStore) ListPolicies(ctx context.Context, token string, pm policies.PageMetadata) (policies.Page, error) {
	return e.svc.ListPolicies(ctx, token, pm)
}

func (e eventStore) ViewPolicyByIDInternal(ctx context.Context, policyID string, ownerID string) (policies.Policy, error) {
	return e.svc.ViewPolicyByIDInternal(ctx, policyID, ownerID)
}

func (e eventStore) ListPoliciesByGroupIDInternal(ctx context.Context, groupIDs []string, ownerID string) ([]policies.Policy, error) {
	return e.svc.ListPoliciesByGroupIDInternal(ctx, groupIDs, ownerID)
}

func (e eventStore) AddDataset(ctx context.Context, token string, d policies.Dataset) (policies.Dataset, error) {
	return e.svc.AddDataset(ctx, token, d)
}

func (e eventStore) CreateDataset(ctx context.Context, token string, d policies.Dataset) (policies.Dataset, error) {
	ds, err := e.svc.AddDataset(ctx, token, d)
	if err != nil {
		return ds, err
	}

	event := createDatasetEvent{
		id:           ds.ID,
		ownerID:      ds.MFOwnerID,
		name:         ds.Name.String(),
		agentGroupID: ds.AgentGroupID,
		policyID:     ds.PolicyID,
		sinkID:       ds.SinkID,
	}
	record := &redis.XAddArgs{
		Stream:       streamID,
		MaxLenApprox: streamLen,
		Values:       event.Encode(),
	}
	err = e.client.XAdd(ctx, record).Err()
	if err != nil {
		e.logger.Error("error sending event to event store", zap.Error(err))
		return ds, err
	}

	return ds, nil
}

func (e eventStore) InactivateDatasetByGroupID(ctx context.Context, groupID string, ownerID string) error {
	return e.svc.InactivateDatasetByGroupID(ctx, groupID, ownerID)
}

// NewEventStoreMiddleware returns wrapper around policies service that sends
// events to event store.
func NewEventStoreMiddleware(svc policies.Service, client *redis.Client, logger *zap.Logger) policies.Service {
	return eventStore{
		logger: logger,
		svc:    svc,
		client: client,
	}
}
