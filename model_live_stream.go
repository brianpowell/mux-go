// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type LiveStream struct {
	Id               string            `json:"id,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	StreamKey        string            `json:"stream_key,omitempty"`
	ActiveAssetId    string            `json:"active_asset_id,omitempty"`
	RecentAssetIds   []string          `json:"recent_asset_ids,omitempty"`
	Status           string            `json:"status,omitempty"`
	PlaybackIds      []PlaybackId      `json:"playback_ids,omitempty"`
	NewAssetSettings Asset             `json:"new_asset_settings,omitempty"`
	Passthrough      string            `json:"passthrough,omitempty"`
	ReconnectWindow  float32           `json:"reconnect_window,omitempty"`
	ReducedLatency   bool              `json:"reduced_latency,omitempty"`
	SimulcastTargets []SimulcastTarget `json:"simulcast_targets,omitempty"`
	Test             bool              `json:"test,omitempty"`
}
