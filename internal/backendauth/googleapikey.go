// Copyright Envoy AI Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package backendauth

import (
	"context"

	"github.com/envoyproxy/ai-gateway/internal/filterapi"
	"github.com/envoyproxy/ai-gateway/internal/internalapi"
)

type googleAIKeyHandler struct{ apiKey string }

func newGoogleAIKeyHandler(auth *filterapi.GoogleAIKeyAuth) (filterapi.BackendAuthHandler, error) {
	return &googleAIKeyHandler{apiKey: auth.Key}, nil
}

// Do injects x-goog-api-key for Google AI Studio (generativelanguage.googleapis.com).
// Google AI Studio uses x-goog-api-key header auth, not Authorization: Bearer.
// https://ai.google.dev/api/rest
func (g *googleAIKeyHandler) Do(_ context.Context, requestHeaders map[string]string, _ []byte) ([]internalapi.Header, error) {
	requestHeaders["x-goog-api-key"] = g.apiKey
	return []internalapi.Header{{"x-goog-api-key", g.apiKey}}, nil
}
