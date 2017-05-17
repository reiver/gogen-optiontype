package shared

const (
	TestTestTmpl =
`
import (
	"math/rand"
	"time"
)

var (
	randomness = rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )
)
`
)
