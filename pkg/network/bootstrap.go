package network

import (
	"github.com/openshift/cluster-network-operator/pkg/bootstrap"
	cnoclient "github.com/openshift/cluster-network-operator/pkg/client"
	"github.com/openshift/cluster-network-operator/pkg/platform"

	operv1 "github.com/openshift/api/operator/v1"
)

// Bootstrap creates resources required by SDN on the cloud.
func Bootstrap(conf *operv1.Network, client cnoclient.Client) (*bootstrap.BootstrapResult, error) {
	out := &bootstrap.BootstrapResult{}

	infraStatus, err := platform.InfraStatus(client)
	if err != nil {
		return nil, err
	}
	out.Infra = *infraStatus

	if conf.Spec.DefaultNetwork.Type == operv1.NetworkTypeOVNKubernetes || (conf.Spec.Migration != nil && conf.Spec.Migration.Mode == operv1.LiveNetworkMigrationMode) {
		o, err := bootstrapOVN(conf, client, infraStatus)
		if err != nil {
			return nil, err
		}
		out.OVN = *o
	}

	return out, nil
}
