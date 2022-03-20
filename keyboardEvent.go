package winapi

import (
	"errors"
	"time"
	"unsafe"
)

//KeyboardEvent is type for simple use keyboard events. Use method Launching
type KeyboardEvent struct {
	Ctrl           bool
	Shift          bool
	JavaScriptCode string
}

//KeyboardInput is winApi struct for SendInput. Use NewKeyboardInput to create it.
type KeyboardInput struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type keyboardInput struct {
	inputType uint32
	ki        KeyboardInput
	padding   uint64
}

//NewKeyboardInput create winApi struct for SendInput
func NewKeyboardInput(k KeyboardInput) *keyboardInput {
	var ki keyboardInput
	ki.inputType = 1 //INPUT_KEYBOARD
	ki.ki = k
	return &ki
}

func (ki *keyboardInput) sendInput() (int, error) {
	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(ki)),
		unsafe.Sizeof(*ki),
	)
	return int(ret), err
}

func downKey(key uint16) error {
	ki := NewKeyboardInput(KeyboardInput{Vk: key})
	ret, err := ki.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

func upKey(key uint16) error {
	ki := NewKeyboardInput(KeyboardInput{Vk: key, Flags: keyUp})
	ret, err := ki.sendInput()
	if ret != 1 {
		return err
	}
	return nil
}

//Launching use for KeyboardEvent struct
func (k *KeyboardEvent) Launching() error {
	if key, ok := javaScriptToUint16[k.JavaScriptCode]; ok {
		if k.Ctrl {
			downKey(ctrl)
			defer upKey(ctrl)
		}
		if k.Shift {
			downKey(shift)
			defer upKey(shift)
		}
		err := downKey(key)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond)
		err = upKey(key)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("not in map")
}

func (ki *keyboardInput) ShiftPress()   { _ = downKey(shift) }
func (ki *keyboardInput) ShiftRelease() { _ = upKey(shift) }
func (ki *keyboardInput) CtrlPress()    { _ = downKey(ctrl) }
func (ki *keyboardInput) CtrlRelease()  { _ = upKey(ctrl) }

const (
	shift = 0x10
	ctrl  = 0x11
	keyUp = 0x0002
)

var javaScriptToUint16 = map[string]uint16{
	"Escape": 0x1B,

	"Insert":   0x2D,
	"Delete":   0x2E,
	"Home":     0x24,
	"End":      0x23,
	"PageUp":   0x21,
	"PageDown": 0x22,

	"Backquote": 0xC0,
	"Digit1":    0x31,
	"Digit2":    0x32,
	"Digit3":    0x33,
	"Digit4":    0x34,
	"Digit5":    0x35,
	"Digit6":    0x36,
	"Digit7":    0x37,
	"Digit8":    0x38,
	"Digit9":    0x39,
	"Digit0":    0x30,
	"Minus":     0xBD,
	"Equal":     0xBB,
	"Backspace": 0x08,

	"KeyQ":         0x51,
	"KeyW":         0x57,
	"KeyE":         0x45,
	"KeyR":         0x52,
	"KeyT":         0x54,
	"KeyY":         0x59,
	"KeyU":         0x55,
	"KeyI":         0x49,
	"KeyO":         0x4F,
	"KeyP":         0x50,
	"BracketLeft":  0xDB,
	"BracketRight": 0xDD,
	"Backslash":    0xDC,

	"CapsLock":  0x14,
	"KeyA":      0x41,
	"KeyS":      0x53,
	"KeyD":      0x44,
	"KeyF":      0x46,
	"KeyG":      0x47,
	"KeyH":      0x48,
	"KeyJ":      0x4A,
	"KeyK":      0x4B,
	"KeyL":      0x4C,
	"Semicolon": 0xBA,
	"Quote":     0xDE,
	"Enter":     0x0D,

	"KeyZ":   0x5A,
	"KeyX":   0x58,
	"KeyC":   0x43,
	"KeyV":   0x56,
	"KeyB":   0x42,
	"KeyN":   0x4E,
	"KeyM":   0x4D,
	"Comma":  0xBC,
	"Period": 0xBE,
	"Slash":  0xBF,

	"Space": 0x20,

	"ArrowUp":    0x26,
	"ArrowDown":  0x28,
	"ArrowRight": 0x27,
	"ArrowLeft":  0x25,

	"Numpad1":        0x31,
	"Numpad2":        0x32,
	"Numpad3":        0x33,
	"Numpad4":        0x34,
	"Numpad5":        0x35,
	"Numpad6":        0x36,
	"Numpad7":        0x37,
	"Numpad8":        0x38,
	"Numpad9":        0x39,
	"Numpad0":        0x30,
	"NumpadDecimal":  0x6E,
	"NumpadEnter":    0x0D,
	"NumpadAdd":      0x6B,
	"NumpadSubtract": 0x6D,
	"NumpadMultiply": 0x6A,
	"NumpadDivide":   0x6F,
}
