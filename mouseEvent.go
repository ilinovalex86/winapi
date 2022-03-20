package winapi

import (
	"time"
	"unsafe"
)

//MouseEvent is type for simple use mouse events.
type MouseEvent struct {
}

//MouseInput is winApi struct for SendInput. Use NewMouseInput to create it.
type MouseInput struct {
	Dx        int32
	Dy        int32
	MouseData int32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type mouseInput struct {
	inputType uint32
	mi        MouseInput
}

func (ms *mouseInput) sendInput() (int, error) {
	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(ms)),
		unsafe.Sizeof(*ms),
	)
	return int(ret), err
}

//NewMouseInput create winApi struct for SendInput
func NewMouseInput(m MouseInput) *mouseInput {
	var mi mouseInput
	mi.inputType = 0
	mi.mi = m
	return &mi
}

func (m *MouseEvent) Move(x int, y int) error {
	ret, _, err := procSetCursorPos.Call(uintptr(x), uintptr(y))
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) LClick(x int, y int) error {
	err := m.Move(x, y)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := NewMouseInput(MouseInput{Flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.mi.Flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) RClick(x int, y int) error {
	err := m.Move(x, y)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := NewMouseInput(MouseInput{Flags: 0x0008})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.mi.Flags = 0x0010
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) DoubleClick(x int, y int) error {
	err := m.Move(x, y)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)
	mi := NewMouseInput(MouseInput{Flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.mi.Flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.mi.Flags = 0x0002
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(time.Millisecond)
	mi.mi.Flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) Drop(x int, y int) error {
	mi := NewMouseInput(MouseInput{Flags: 0x0002})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	err = m.Move(x, y)
	if err != nil {
		return err
	}
	time.Sleep(25 * time.Millisecond)
	mi.mi.Flags = 0x0004
	ret, err = mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) WheelUp() error {
	mi := NewMouseInput(MouseInput{Flags: 0x0800, MouseData: 120})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func (m *MouseEvent) WheelDown() error {
	mi := NewMouseInput(MouseInput{Flags: 0x0800, MouseData: -120})
	ret, err := mi.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}
