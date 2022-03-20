package winapi

type inputEvent interface {
	sendInput() (int, error)
}

//SendInput uses struct from NewKeyboardInput or NewMouseInput.
func SendInput(i inputEvent) (int, error) {
	ret, err := i.sendInput()
	return ret, err
}
