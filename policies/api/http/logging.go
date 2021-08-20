/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package http

import (
	"context"
	"github.com/ns1labs/orb/policies"
	"go.uber.org/zap"
	"time"
)

var _ policies.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *zap.Logger
	svc    policies.Service
}

func (l loggingMiddleware) RetrievePoliciesByGroupIDInternal(ctx context.Context, groupIDs []string, ownerID string) (_ []policies.Policy, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: retrieve_policies_by_groups",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: retrieve_policies_by_groups",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.RetrievePoliciesByGroupIDInternal(ctx, groupIDs, ownerID)
}

func (l loggingMiddleware) RetrievePolicyByIDInternal(ctx context.Context, policyID string, ownerID string) (_ policies.Policy, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: retrieve_policy_by_id",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: retrieve_policy_by_id",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.RetrievePolicyByIDInternal(ctx, policyID, ownerID)
}

func (l loggingMiddleware) CreateDataset(ctx context.Context, token string, d policies.Dataset) (_ policies.Dataset, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: create_dataset",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: create_dataset",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.CreateDataset(ctx, token, d)
}

func (l loggingMiddleware) CreatePolicy(ctx context.Context, token string, p policies.Policy, format string, policyData string) (_ policies.Policy, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: create_policy",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: create_policy",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.CreatePolicy(ctx, token, p, format, policyData)
}

func (l loggingMiddleware) InactivateDatasetByGroupID(ctx context.Context, groupID string, ownerID string) (err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: inactivate_dataset",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: inactivate_dataset",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.InactivateDatasetByGroupID(ctx, groupID, ownerID)
}

func NewLoggingMiddleware(svc policies.Service, logger *zap.Logger) policies.Service {
	return &loggingMiddleware{logger, svc}
}
