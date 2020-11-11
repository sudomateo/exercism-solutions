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
	namesInUse = make(map[string]struct{})
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
	name, err := genName()
	if err != nil {
		return "", err
	}
	for {
		if len(namesInUse) >= maxNumNames {
			return "", errors.New("exhausted namespace")
		}
		if _, ok := namesInUse[name]; ok {
			name, err = genName()
			if err != nil {
				return "", err
			}
			continue
		}
		namesInUse[name] = struct{}{}
		r.name = name
		break
	}
	return r.name, nil
}

// Reset clears out the name for a robot.
func (r *Robot) Reset() {
	if r.name != "" {
		r.name = ""
	}
}

// genName generates a unique name for a robot in the format
// of two uppercase letters followed by three digits.
func genName() (string, error) {
	rand.Seed(time.Now().UnixNano())
	id := genId()
	str := genStr()
	name := str + id
	return name, nil
}

// genId generates a random string of three digits.
func genId() string {
	n := rand.Intn(999 + 1)
	id := fmt.Sprintf("%03d", n)
	return id
}

// genStr generates a random string of two uppercase letters.
func genStr() string {
	r := make([]rune, 2)
	for i := range r {
		r[i] = rune(rand.Intn('Z'-'A'+1) + 'A')
	}
	return string(r)
}
