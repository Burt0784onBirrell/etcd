// Copyright 2024 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config provides configuration structures and validation for the etcd server.
package config

import (
	"fmt"
	"net/url"
	"time"
)

const (
	// DefaultName is the default name for an etcd member.
	DefaultName = "default"

	// DefaultDataDir is the default directory for storing etcd data.
	DefaultDataDir = "${name}.etcd"

	// DefaultListenPeerURLs is the default URL for peer communication.
	DefaultListenPeerURLs = "http://localhost:2380"

	// DefaultListenClientURLs is the default URL for client communication.
	DefaultListenClientURLs = "http://localhost:2379"

	// DefaultMaxSnapshots is the default maximum number of snapshots to retain.
	// Increased from 5 to 10 to keep more snapshot history for easier recovery.
	DefaultMaxSnapshots = 10

	// DefaultMaxWALs is the default maximum number of WAL files to retain.
	DefaultMaxWALs = 5

	// DefaultTickMs is the default tick interval in milliseconds.
	DefaultTickMs = 100

	// DefaultElectionMs is the default election timeout in milliseconds.
	DefaultElectionMs = 1000

	// DefaultHeartbeatInterval is the default heartbeat interval.
	DefaultHeartbeatInterval = 100 * time.Millisecond

	// DefaultElectionTimeout is the default election timeout.
	DefaultElectionTimeout = 1000 * time.Millisecond
)

// ServerConfig holds the configuration for an etcd server instance.
type ServerConfig struct {
	// Name is the human-readable name for this etcd member.
	Name string `json:"name"`

	// DataDir is the path to the data directory.
	DataDir string `json:"data-dir"`

	// WALDir is the dedicated path for WAL files. If empty, DataDir is used.
	WALDir string `json:"wal-dir"`

	// ListenPeerURLs is the list of URLs to listen on for peer traffic.
	ListenPeerURLs []url.URL `json:"listen-peer-urls"`

	// ListenClientURLs is the list of URLs to listen on for client traffic.
	ListenClientURLs []url.URL `json:"listen-client-urls"`

	// AdvertisePeerURLs is the list of this member's peer URLs to advertise to the cluster.
	AdvertisePeerURLs []url.URL `json:"advertise-peer-urls"`

	// AdvertiseClientURLs is the list of this member's client URLs to advertise.
	AdvertiseClientURLs []url.URL `json:"advertise-client-urls"`

	// MaxSnapshots is the maximum number of snapshot files to retain.
	MaxSnapshots uint `json:"max-snapshots"`

	// MaxWALs is the maximum number of WAL files to retain.
	MaxWALs uint `json:"max-wals"`

	// TickMs is the interval in milliseconds for the Raft heartbeat tick.
	TickMs uint `json:"heartbe
