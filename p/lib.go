package p

var (
	name    = "c"
	version = "v1.1.0"
)

func Version() string {
	return name + ": " + version
}
