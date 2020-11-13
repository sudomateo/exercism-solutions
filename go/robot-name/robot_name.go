package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	maxNumNames = 26 * 26 * 10 * 10 * 10
)

var (
	namesInUse = make(map[string]bool)
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Robot represents a factory robot.
type Robot struct {
	name string
}

// Name generates a unique name for a robot and sets the robot's name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(namesInUse) >= maxNumNames {
		return "", errors.New("exhausted namespace")
	}
	r.name = genName()
	for namesInUse[r.name] {
		r.name = genName()
	}
	namesInUse[r.name] = true
	return r.name, nil
}

// Reset clears out the name for a robot.
func (r *Robot) Reset() {
	r.name = ""
}

// genName generates a unique name for a robot in the format
// of two uppercase letters followed by three digits.
func genName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
