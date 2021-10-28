package topology

import (
	"github.com/zhongzc/ng_monitoring/config"
	"go.etcd.io/etcd/clientv3"
)

var (
	discover *TopologyDiscoverer
	syncer   *TopologySyncer
)

func Init() error {
	var err error
	discover, err = NewTopologyDiscoverer(config.GetGlobalConfig())
	if err != nil {
		return err
	}
	syncer = NewTopologySyncer(discover.etcdCli)
	syncer.Start()
	discover.Start()
	return err
}

func GetCurrentComponent() []Component {
	if discover == nil {
		return nil
	}
	components := make([]Component, 0, len(discover.components))
	for _, comp := range discover.components {
		components = append(components, comp)
	}
	return components
}

func GetEtcdClient() *clientv3.Client {
	if discover == nil {
		return nil
	}
	return discover.etcdCli
}

func Subscribe() Subscriber {
	if discover == nil {
		return nil
	}
	return discover.Subscribe()
}

func Stop() {
	if syncer == nil {
		return
	}
	syncer.Stop()
}
