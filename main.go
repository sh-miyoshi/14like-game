package main

import (
	"runtime"

	"github.com/sh-miyoshi/14like-game/pkg/app/background"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/manager"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/result"
	"github.com/sh-miyoshi/14like-game/pkg/app/title"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/fps"
	"github.com/sh-miyoshi/14like-game/pkg/inputs"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/sound"
)

var (
	state = 3
	count = 0
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fps.FPS = 60
	fpsMgr := fps.Fps{}

	dxlib.Init("data/DxLib.dll")
	dxlib.SetDoubleStartValidFlag(dxlib.TRUE)
	dxlib.SetAlwaysRunFlag(dxlib.TRUE)
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)
	dxlib.ChangeWindowMode(dxlib.TRUE)
	logger.InitLogger(true, "application.log")
	dxlib.SetGraphMode(config.ScreenSizeX, config.ScreenSizeY)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	inputs.Init(inputs.DeviceTypeKeyboard)
	config.SkillNumberFontHandle = dxlib.CreateFontToHandle(dxlib.CreateFontToHandleOption{
		Size: dxlib.Int32Ptr(10),
	})

	config.Init()
	sound.SEInit()

	mgr := manager.Manager{}

	bg := background.BackGround{}

	titleInst := title.NewTitle()
	resultInst := result.NewResult()

MAIN:
	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		inputs.KeyStateUpdate()

		// Main Game Proc
		switch state {
		case 0:
			if titleInst.Update() {
				stateChange(1)
				sound.BGMPlay()
				continue
			}
			titleInst.Draw()
		case 1:
			if count == 0 {
				mgr.Init()
				mgr.AddObject(models.ObjectInstPlayer, nil)
				mgr.AddObject(models.ObjectInstCloudOfDarkness, nil)
				bg.Init(config.Phase1)
			}

			bg.Update()
			mgr.Update()
			if mgr.IsEnd() {
				stateChange(2)
				sound.BGMStop()
				resultInst.SetValues(mgr.GetResult().Hits)
				continue
			}

			bg.Draw()
			mgr.Draw()
		case 2:
			if resultInst.Update() {
				stateChange(0)
				continue
			}
			resultInst.Draw()
		case 3: // Debug
			if count == 0 {
				mgr.Init()
				mgr.AddObject(models.ObjectInstPlayer, nil)
				mgr.AddObject(models.ObjectInstStygianShadow, nil)
				bg.Init(config.Phase2A)
			}

			bg.Update()
			mgr.Update()

			bg.Draw()
			mgr.Draw()
		}
		count++

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			logger.Info("Game end by escape command")
			break MAIN
		}

		fpsMgr.Wait()
	}

	sound.BGMStop()

	dxlib.DxLib_End()
}

func stateChange(next int) {
	state = next
	count = 0
}
