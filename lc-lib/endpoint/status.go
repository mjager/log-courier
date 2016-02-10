/*
 * Copyright 2014-2015 Jason Woods.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package endpoint

// status holds an Endpoint status
type status int

// Endpoint statuses
const (
	// Not yet ready
	endpointStatusIdle status = iota

	// Ready to receive events
	endpointStatusReady

	// Busy
	endpointStatusBusy

	// Could receive events but too many are oustanding
	endpointStatusFull

	// Do not use this endpoint, it has failed
	endpointStatusFailed

	// The endpoint is about to shutdown once pending payloads are complete
	endpointStatusClosing
)