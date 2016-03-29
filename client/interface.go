package client

import ()

// APIClient is an interface that clients that talk with a daolinet server must implement.
type APIClient interface {
	GroupList() ([]string, error)
	GroupShow(string) ([]string, error)
	GroupCreate(string) (string, error)
	GroupDelete(string) error
	MemberAdd(string, string) error
	MemberRemove(string, string) error
	PolicyList() ([]string, error)
	PolicyCreate(string) error
	PolicyDelete(string) error
	FirewallList() ([]Firewall, error)
	FirewallShow(string) ([]Firewall, error)
	FirewallCreate(Firewall) (Firewall, error)
	FirewallDelete(string) error
}

// Ensure that Client always implements APIClient.
var _ APIClient = &Client{}
