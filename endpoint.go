/*
The MIT License (MIT)

Copyright © 2020 Neil Hemming

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package okta provides the OKTA endpoints for oauth2 access.
// see https://godoc.org/golang.org/x/oauth2 for more detailson use
package okta

import (
	"fmt"

	"golang.org/x/oauth2"
)

// NewEndpoint creates a new endpoint for Okta authentication
// oktaDomain and authServerID are part of the OKTA tenant
func NewEndpoint(oktaDomain string, authServerID string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  fmt.Sprintf("https://%s.okta.com/oauth2/%s/v1/authorize", oktaDomain, authServerID),
		TokenURL: fmt.Sprintf("https://%s.okta.com/oauth2/%s/v1/token", oktaDomain, authServerID),
	}
}
