/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliec.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"encoding/json"
	"fmt"
)

type MachineIdentityUniversalAuthRefreshRequest struct {
	AccessToken string `json:"accessToken"`
}

type InfisicalAPIError struct {
	StatusCode int
	Err        any
	Message    any
	Details    any
}

func (e *InfisicalAPIError) Error() string {
	if e.Details != nil {
		detailsJSON, _ := json.Marshal(e.Details)
		return fmt.Sprintf("API error (%d): error=%v message=%v, details=%s", e.StatusCode, e.Err, e.Message, string(detailsJSON))
	}
	return fmt.Sprintf("API error (%d): error=%v message=%v", e.StatusCode, e.Err, e.Message)
}

type MachineIdentityDetailsResponse struct {
	AccessToken       string `json:"accessToken"`
	ExpiresIn         int    `json:"expiresIn"`
	AccessTokenMaxTTL int    `json:"accessTokenMaxTTL"`
	TokenType         string `json:"tokenType"`
}

type RevokeMachineIdentityAccessTokenResponse struct {
	Message string `json:"message"`
}

type GetSecretByKeyV3Response struct {
	Secret SecretsV3 `json:"secret"`
}

type GetSecretsV3Response struct {
	Secrets         []SecretsV3        `json:"secrets"`
	ImportedSecrets []ImportedSecretV3 `json:"imports,omitempty"`
	Modified        bool               `json:"modified,omitempty"`
	ETag            string             `json:"ETag,omitempty"`
}

type SecretsV3 struct {
	ID            string `json:"id"`
	Workspace     string `json:"workspace"`
	Environment   string `json:"environment"`
	Version       int    `json:"version"`
	Type          string `json:"string"`
	SecretKey     string `json:"secretKey"`
	SecretValue   string `json:"secretValue"`
	SecretComment string `json:"secretComment"`
}

type ImportedSecretV3 struct {
	Environment string      `json:"environment"`
	FolderID    string      `json:"folderId"`
	SecretPath  string      `json:"secretPath"`
	Secrets     []SecretsV3 `json:"secrets"`
}

type InfisicalAPIErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	// According to Infisical's API docs, `details` are only returned for 403 errors.
	Details any `json:"details,omitempty"`
}
