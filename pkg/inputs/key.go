package inputs

import "github.com/sh-miyoshi/14like-game/pkg/dxlib"

type keyboard struct {
	keyState [256]int
	keyBind  [keyMax]int
}

func (k *keyboard) Init() error {
	k.keyBind[KeyEnter] = dxlib.KEY_INPUT_SPACE
	k.keyBind[KeyLeft] = dxlib.KEY_INPUT_LEFT
	k.keyBind[KeyRight] = dxlib.KEY_INPUT_RIGHT
	k.keyBind[KeyUp] = dxlib.KEY_INPUT_UP
	k.keyBind[KeyDown] = dxlib.KEY_INPUT_DOWN
	k.keyBind[KeyAnotherLeft] = dxlib.KEY_INPUT_A
	k.keyBind[KeyAnotherRight] = dxlib.KEY_INPUT_D
	k.keyBind[KeyAnotherUp] = dxlib.KEY_INPUT_W
	k.keyBind[KeyAnotherDown] = dxlib.KEY_INPUT_S

	return nil
}

func (k *keyboard) KeyStateUpdate() {
	tmp := make([]byte, 256)
	dxlib.GetHitKeyStateAll(tmp)
	for i := 0; i < 256; i++ {
		if tmp[i] == 1 {
			k.keyState[i]++
		} else {
			k.keyState[i] = 0
		}
	}
}

func (k *keyboard) CheckKey(key KeyType) int {
	switch key {
	case KeyLeft:
		return k.keyState[k.keyBind[key]] + k.keyState[k.keyBind[KeyAnotherLeft]]
	case KeyRight:
		return k.keyState[k.keyBind[key]] + k.keyState[k.keyBind[KeyAnotherRight]]
	case KeyUp:
		return k.keyState[k.keyBind[key]] + k.keyState[k.keyBind[KeyAnotherUp]]
	case KeyDown:
		return k.keyState[k.keyBind[key]] + k.keyState[k.keyBind[KeyAnotherDown]]
	}

	return k.keyState[k.keyBind[key]]
}
