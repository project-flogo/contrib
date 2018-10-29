// Copyright (c) 2015 TIBCO Software Inc.
// All Rights Reserved.

//Cors package to validate CORS requests and provide the correct headers in the response
package cors

import (
	"net/http"
	"strings"

	"github.com/project-flogo/core/support/log"
)

const (
	HeaderOrigin                        string = "Origin"
	HeaderAccessControlRequestMethod    string = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   string = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      string = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     string = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     string = "Access-Control-Allow-Headers"
	HeaderAccessControlExposeHeaders    string = "Access-Control-Expose-Headers"
	HeaderAccessControlAllowCredentials string = "Access-Control-Allow-Credentials"
	HeaderAccessControlMaxAge           string = "Access-Control-Max-Age"
)

// CORS interface
type Cors interface {
	// HandlePreflight Handles the preflight OPTIONS request
	HandlePreflight(w http.ResponseWriter, r *http.Request)
	// WriteCorsActualRequestHeaders writes the needed request headers for the CORS support
	WriteCorsActualRequestHeaders(w http.ResponseWriter)
}

type cors struct {
	// Prefix used for the CORS environment variables
	Prefix string

	logger log.Logger
}

// make sure that the cors implements the Cors interface
var _ Cors = (*cors)(nil)

//Cors constructor
func New(prefix string, logger log.Logger) Cors {
	return cors{Prefix: prefix, logger: logger}
}

// HandlePreflight Handles the cors preflight request setting the right headers and responding to the request
func (c cors) HandlePreflight(w http.ResponseWriter, r *http.Request) {
	// Check if it has Origin Header
	hasOrigin := HasOriginHeader(r)
	if hasOrigin == false {
		c.logger.Info("Invalid CORS preflight request, no Origin header found")
		writeInvalidPreflightResponse(w)
		return
	}

	// Check Access-Control-Request-Method header
	requestMethodHeader := r.Header.Get(HeaderAccessControlRequestMethod)
	if isValidAccessControlMethod(requestMethodHeader, c.Prefix, c.logger) != true {
		// Invalid Access Control Method
		writeInvalidPreflightResponse(w)
		return
	}

	// Check Access-Control-Allow-Headers header
	requestHeadersHeader := r.Header.Get(HeaderAccessControlRequestHeaders)
	if isValidAccessControlHeaders(requestHeadersHeader, c.Prefix, c.logger) != true {
		// Invalid Access Control Header
		writeInvalidPreflightResponse(w)
		return
	}

	writeValidPreflightResponse(w, c)
}

// HasOriginHeader returns true if the request has Origin header, false otherwise
func HasOriginHeader(r *http.Request) bool {
	h := r.Header.Get(HeaderOrigin)
	if h == "" {
		return false
	}
	return true
}

// Check if the method name is valid and allowed by the environment variable
func isValidAccessControlMethod(methodName string, prefix string, logger log.Logger) bool {
	if methodName == "" {
		logger.Infof("Invalid Access Control Method for preflight request: '%s'", methodName)
		return false
	}
	allowedMethodsEnv := GetCorsAllowMethods(prefix)
	allowedMethods := strings.Split(allowedMethodsEnv, ",")
	logger.Debugf("Allowed Methods '%s'", allowedMethods)
	for i := range allowedMethods {
		if strings.ToLower(strings.TrimSpace(allowedMethods[i])) == strings.ToLower(strings.TrimSpace(methodName)) {
			return true
		}
	}
	logger.Infof("Invalid Access Control Method for preflight request: '%s'", methodName)
	return false
}

// Check if the headers are valid and allowed by the environment variable
func isValidAccessControlHeaders(headersStr string, prefix string, logger log.Logger) bool {
	if headersStr == "" {
		return true
	}
	allowedHeadersEnv := GetCorsAllowHeaders(prefix)
	allowedHeaders := strings.Split(allowedHeadersEnv, ",")

	// Create a map for faster lookup
	allowedHeadersMap := make(map[string]struct{}, len(allowedHeaders))
	for _, s := range allowedHeaders {
		allowedHeadersMap[strings.ToLower(strings.TrimSpace(s))] = struct{}{}
	}

	headers := strings.Split(headersStr, ",")

	for i := range headers {
		_, ok := allowedHeadersMap[strings.ToLower(strings.TrimSpace(headers[i]))]
		if ok == false {
			logger.Infof("Invalid Access Control Header for pre-flight request: '%s'", strings.TrimSpace(headers[i]))
			return false
		}
	}
	return true
}

// Writes invalid preflight response
func writeInvalidPreflightResponse(w http.ResponseWriter) {
	// Write 200 but no CORS header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Writes valid preflight response
func writeValidPreflightResponse(w http.ResponseWriter, c cors) {
	// Write 200 with CORS headers
	writeCorsPreflightHeaders(w, c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Writes the CORS preflight request headers (origin and credential)
func writeCorsPreflightHeaders(w http.ResponseWriter, c cors) {
	c.WriteCorsActualRequestHeaders(w)
	w.Header().Set(HeaderAccessControlAllowMethods, GetCorsAllowMethods(c.Prefix))
	w.Header().Set(HeaderAccessControlAllowHeaders, GetCorsAllowHeaders(c.Prefix))
	w.Header().Set(HeaderAccessControlExposeHeaders, GetCorsExposeHeaders(c.Prefix))
	maxAge := GetCorsMaxAge(c.Prefix)
	if maxAge != "" {
		w.Header().Set(HeaderAccessControlMaxAge, maxAge)
	}
}

// Writes the CORS actual request headers (origin and credential)
func (c cors) WriteCorsActualRequestHeaders(w http.ResponseWriter) {
	w.Header().Set(HeaderAccessControlAllowOrigin, GetCorsAllowOrigin(c.Prefix))
	allowCredentials := GetCorsAllowCredentials(c.Prefix)
	if strings.TrimSpace(allowCredentials) == "true" {
		w.Header().Set(HeaderAccessControlAllowCredentials, strings.TrimSpace(allowCredentials))
	}
}
