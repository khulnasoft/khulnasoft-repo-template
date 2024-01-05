package broker

var (
	Request          ObjectType = "request-payload"
	MeshSync         ObjectType = "meshsync-data"
	LogStreamObject  ObjectType = "log-stream"
	SMI              ObjectType = "smi-data"
	ErrorObject      ObjectType = "error"
	ExecInputObject  ObjectType = "exec-input"
	ExecOutputObject ObjectType = "exec-output"
	ActiveExecObject ObjectType = "active-exec"

	Add        EventType = "ADDED"
	Update     EventType = "MODIFIED"
	Delete     EventType = "DELETED"
	ErrorEvent EventType = "ERROR"
	ReSync     EventType = "RESYNC"

	LogRequestEntity      RequestEntity = "log-stream"
	ReSyncDiscoveryEntity RequestEntity = "resync-discovery"
	ExecRequestEntity     RequestEntity = "exec-request"
	ActiveExecEntity      RequestEntity = "active-exec"
)

type ObjectType string
type EventType string
type RequestEntity string

type Message struct {
	ObjectType ObjectType
	EventType  EventType
	Request    *RequestObject
	Object     interface{}
}

type RequestObject struct {
	Entity  RequestEntity
	Payload interface{}
}
