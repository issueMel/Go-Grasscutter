package pros

type EnterReason uint32

const (
	None  EnterReason = 0
	Login EnterReason = 1
)

var EnterReason_name = map[EnterReason]string{
	0: "None",
	1: "Login",
}

var EnterReason_value = map[string]uint32{
	"None":  0,
	"Login": 1,
}
