//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package auth

const (
	defaultPort       = 8089
	defaultDistroHost = "127.0.0.1"
)

type Config struct {
	Port       int
	DistroHost string
}

var cfg Config

func GetDefaultConfig() Config {
	return Config{
		Port:       defaultPort,
		DistroHost: defaultDistroHost,
	}
}
