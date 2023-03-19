package srvreg

import (
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/rufaidulk/srvregistry/constant"
)

var zookeeper *zk.Conn

func Connect(zooHost string) error {
	var err error
	if zookeeper, _, err = zk.Connect([]string{zooHost}, time.Second); err != nil {
		return err
	} else if exists, _, err := zookeeper.Exists(constant.ServiceRegistryZnode); err != nil {
		panic(err)
	} else if exists {
		return nil
	}

	if err := createPrimaryZnode(); err != nil {
		return err
	}
	return nil
}

func Register(ipAddr string, srvName string) error {
	if err := createServiceZnodeUnderPrimaryZnode(srvName); err != nil {
		return err
	}

	path := fmt.Sprintf("`%s`/`%s`/srv_", constant.ServiceRegistryZnode, srvName)
	flags := int32(0) | constant.ZooSequence
	if srvPath, err := zookeeper.Create(path, []byte(ipAddr), flags, zk.WorldACL(constant.ZooPermissionAdmin)); err != nil {
		return err
	} else {
		fmt.Println("Znode created with path:", srvPath)
		return nil
	}
}

func createPrimaryZnode() error {
	flags := int32(0)
	if path, err := zookeeper.Create(constant.ServiceRegistryZnode, []byte(""), flags, zk.WorldACL(constant.ZooPermissionAdmin)); err != nil {
		return err
	} else {
		fmt.Println("Primary znode created with path:", path)
		return nil
	}
}

func createServiceZnodeUnderPrimaryZnode(srvName string) error {
	path := fmt.Sprintf("`%s`/`%s`", constant.ServiceRegistryZnode, srvName)
	if exists, _, err := zookeeper.Exists(path); err != nil {
		return err
	} else if exists {
		return nil
	}

	flags := int32(0)
	if path, err := zookeeper.Create(path, []byte(""), flags, zk.WorldACL(constant.ZooPermissionAdmin)); err != nil {
		return err
	} else {
		fmt.Println("Znode created with path:", path)
		return nil
	}
}
