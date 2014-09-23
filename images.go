package main

import (
	"time"
)

type Image struct {
	Id              string      `json:"id"`
	Parent          string      `json:"parent,omitempty"`
	Created         time.Time   `json:"created"`
	Container       string      `json:"container,omitempty"`
	ContainerConfig ImageConfig `json:"containerconfig,omitempty"`
	DockerVersion   string      `json:"docker_version,omitempty"`
	Config          ImageConfig `json:"config,omitempty"`
	Architecture    string      `json:"architecture,omitempty"`
	OS              string      `json:"linux,omitempty"`
	Size            int64
}

type ImageConfig struct {
	Hostname        string
	Domainname      string
	User            string
	Memory          int
	MemorySwap      int
	CpuShares       int
	Cpuset          string
	AttachStdin     bool
	AttachStdout    bool
	AttachStderr    bool
	PortSpecs       []string
	ExposedPorts    map[string]struct{}
	Tty             bool
	OpenStdin       bool
	StdinOnce       bool
	Env             []string
	Cmd             []string
	Image           string
	Volumes         map[string]struct{}
	WorkingDir      string
	Entrypoint      []string
	NetworkDisabled bool
	OnBuild         []string
}
