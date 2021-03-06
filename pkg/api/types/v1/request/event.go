//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package request

import "time"

type EventOptions struct {
	Builds    map[string]BuildEvent `json:"builds"`
	Resources ResourcesEvent        `json:"resources"`
}

type BuildEvent struct {
	ID         string     `json:"id"`
	Image      string     `json:"image"`
	Source     string     `json:"source"`
	Size       int64      `json:"size"`
	Step       string     `json:"step"`
	Status     string     `json:"status"`
	Message    string     `json:"message"`
	Processing bool       `json:"processing"`
	Done       bool       `json:"done"`
	Error      bool       `json:"error"`
	Canceled   bool       `json:"canceled"`
	Finished   *time.Time `json:"finished"`
	Started    *time.Time `json:"started"`
}

type ResourcesEvent struct {
	Capacity  struct{} `json:"capacity"`
	Allocated struct{} `json:"allocated"`
	Builders  struct{} `json:"builders"`
}
