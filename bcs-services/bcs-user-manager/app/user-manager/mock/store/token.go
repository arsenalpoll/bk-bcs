/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package store xxx
package store

import (
	"github.com/stretchr/testify/mock"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/user-manager/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/user-manager/storages/sqlstore"
)

// MockTokenStore is a mock of TokenStore
type MockTokenStore struct {
	mock.Mock
}

// GetTokenByCondition mock get token by condition
func (m *MockTokenStore) GetTokenByCondition(cond *models.BcsUser) *models.BcsUser {
	args := m.Called(cond)
	return args.Get(0).(*models.BcsUser)
}

// GetUserTokensByName mock get user tokens by name
func (m *MockTokenStore) GetUserTokensByName(name string) []models.BcsUser {
	args := m.Called(name)
	return args.Get(0).([]models.BcsUser)
}

// CreateToken mock create token
func (m *MockTokenStore) CreateToken(token *models.BcsUser) error {
	args := m.Called(token)
	return args.Error(0)
}

// UpdateToken mock update token
func (m *MockTokenStore) UpdateToken(token, updatedToken *models.BcsUser) error {
	args := m.Called(token, updatedToken)
	return args.Error(0)
}

// DeleteToken mock delete token
func (m *MockTokenStore) DeleteToken(token string) error {
	args := m.Called(token)
	return args.Error(0)
}

// CreateTemporaryToken mock create temporary token
func (m *MockTokenStore) CreateTemporaryToken(token *models.BcsTempToken) error {
	args := m.Called(token)
	return args.Error(0)
}

// GetTempTokenByCondition mock create temporary token
func (m *MockTokenStore) GetTempTokenByCondition(cond *models.BcsTempToken) *models.BcsTempToken {
	args := m.Called(cond)
	return args.Get(0).(*models.BcsTempToken)
}

// GetAllNotExpiredTokens mock create temporary token
func (m *MockTokenStore) GetAllNotExpiredTokens() []models.BcsUser {
	args := m.Called()
	return args.Get(0).([]models.BcsUser)
}

// GetAllTokens mock create temporary token
func (m *MockTokenStore) GetAllTokens() []models.BcsUser {
	args := m.Called()
	return args.Get(0).([]models.BcsUser)
}

var _ sqlstore.TokenStore = &MockTokenStore{}
