package constant

const (
	ServiceRegistryZnode = "/serviceRegistry"
)
const ZooEphemeral int32 = 1
const ZooSequence int32 = 2
const ZooContainer int32 = 4
const ZooTTL int32 = 8

const ZooPermissionRead int32 = 1
const ZooPermissionWrite int32 = 2
const ZooPermissionCreate int32 = 4
const ZooPermissionDelete int32 = 8
const ZooPermissionAdmin int32 = 16
