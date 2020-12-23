package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type State struct {
	PIR1 bool
	PIR2 bool
	PIR3 bool
	PIR4 bool
}

func (s State) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *State) Deserialize(b []byte) error {
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, s)
	if err != nil {
		return err
	}
	return nil
}

func (s State) String() string {
	return fmt.Sprintf("1: %t, 2: %t, 3: %t, 4: %t", s.PIR1, s.PIR2, s.PIR3, s.PIR4)
}
