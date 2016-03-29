package client

import (
	"encoding/json"
	"path"
)

func (cli *Client) GroupList() ([]string, error) {
	var groups []string
	resp, err := cli.get("/api/groups", nil, nil)
	if err != nil {
		return groups, err
	}

	err = json.NewDecoder(resp.body).Decode(&groups)
	ensureReaderClosed(resp)
	return groups, err
}

func (cli *Client) GroupShow(name string) ([]string, error) {
	var members []string
	resp, err := cli.get("/api/groups/"+name, nil, nil)
	if err != nil {
		return members, err
	}

	err = json.NewDecoder(resp.body).Decode(&members)
	ensureReaderClosed(resp)
	return members, err
}

func (cli *Client) GroupCreate(name string) (string, error) {
	group := map[string]string{"name": name}
	resp, err := cli.post("/api/groups", nil, group, nil)
	ensureReaderClosed(resp)
	return name, err
}

func (cli *Client) GroupDelete(name string) error {
	resp, err := cli.delete("/api/groups/"+name, nil, nil)
	ensureReaderClosed(resp)
	return err
}

func (cli *Client) MemberAdd(name, group string) error {
	member := map[string]string{"member": name}
	resp, err := cli.post("/api/groups/"+group, nil, member, nil)
	ensureReaderClosed(resp)
	return err
}

func (cli *Client) MemberRemove(name, group string) error {
	url := path.Join("/api/groups", group, name)
	resp, err := cli.delete(url, nil, nil)
	ensureReaderClosed(resp)
	return err
}

func (cli *Client) PolicyList() ([]string, error) {
	var policy []string
	resp, err := cli.get("/api/policy", nil, nil)
	if err != nil {
		return policy, err
	}

	err = json.NewDecoder(resp.body).Decode(&policy)
	ensureReaderClosed(resp)
	return policy, err
}

func (cli *Client) PolicyCreate(peer string) error {
	resp, err := cli.post("/api/policy/"+peer, nil, nil, nil)
	ensureReaderClosed(resp)
	return err
}

func (cli *Client) PolicyDelete(peer string) error {
	resp, err := cli.delete("/api/policy/"+peer, nil, nil)
	ensureReaderClosed(resp)
	return err
}

func (cli *Client) FirewallList() ([]Firewall, error) {
	resp, err := cli.get("/api/firewalls", nil, nil)
	if err != nil {
		return nil, err
	}
	firewalls := []Firewall{}
	err = json.NewDecoder(resp.body).Decode(&firewalls)
	ensureReaderClosed(resp)
	return firewalls, err
}

func (cli *Client) FirewallShow(name string) ([]Firewall, error) {
	resp, err := cli.get("/api/firewalls/"+name, nil, nil)
	firewalls := []Firewall{}
	if err != nil {
		return firewalls, err
	}

	err = json.NewDecoder(resp.body).Decode(&firewalls)
	ensureReaderClosed(resp)
	return firewalls, err
}

func (cli *Client) FirewallCreate(fw Firewall) (Firewall, error) {
	var firewall Firewall
	resp, err := cli.post("/api/firewalls", nil, fw, nil)
	if err == nil {
		err = json.NewDecoder(resp.body).Decode(&firewall)
	}
	ensureReaderClosed(resp)
	return firewall, err
}

func (cli *Client) FirewallDelete(name string) error {
	resp, err := cli.delete("/api/firewalls/"+name, nil, nil)
	ensureReaderClosed(resp)
	return err
}
