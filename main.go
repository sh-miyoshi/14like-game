package main

import (
	"runtime"

	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/manager"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/fps"
	"github.com/sh-miyoshi/14like-game/pkg/inputs"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
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

	// WIP: 別の場所で管理
	player := object.Player{}
	player.Init(manager.GetDamageManager().AddDamage)

	enemy1 := object.Enemy1{}
	enemy1.Init(manager.GetDamageManager().AddDamage)

	manager.GetDamageManager().SetInsts([]object.Object{&player, &enemy1})
	player.SetTargetEnemy(&enemy1)

MAIN:
	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		inputs.KeyStateUpdate()

		// Main Game Proc
		player.Update()
		enemy1.Update()
		manager.GetDamageManager().Update()

		player.Draw()
		enemy1.Draw()

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			logger.Info("Game end by escape command")
			break MAIN
		}

		fpsMgr.Wait()
	}

	player.End()
	enemy1.End()

	dxlib.DxLib_End()
}
