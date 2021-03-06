package gosnake

import (
	"io"

	"github.com/NBR41/gosnake/assets"
	"github.com/hajimehoshi/ebiten/audio"
)

const (
	// This sample rate doesn't match with wav/vorbis's sample rate,
	// but decoders adjust them.
	sampleRate = 48000
)

type audioReg struct {
	ctx     *audio.Context
	players *audioPlayers
}

type audioPlayers struct {
	Beginning *audio.Player
	Chomp     *audio.Player
	Death     *audio.Player
}

func newAudio() (*audioReg, error) {
	audioCtx, err := audio.NewContext(sampleRate)
	if err != nil {
		return nil, err
	}

	sounds, err := assets.LoadSounds(audioCtx)
	if err != nil {
		return nil, err
	}

	beginning, err := newAudioPlayer(audioCtx, sounds.Beginning)
	if err != nil {
		return nil, err
	}

	chomp, err := newAudioPlayer(audioCtx, sounds.Chomp)
	if err != nil {
		return nil, err
	}

	death, err := newAudioPlayer(audioCtx, sounds.Death)
	if err != nil {
		return nil, err
	}

	return &audioReg{
		ctx: audioCtx,
		players: &audioPlayers{
			Beginning: beginning,
			Chomp:     chomp,
			Death:     death,
		},
	}, nil
}

func newAudioPlayer(ctx *audio.Context, src io.ReadCloser) (*audio.Player, error) {
	player, err := audio.NewPlayer(ctx, src)
	if err != nil {
		return nil, err
	}
	player.SetVolume(0.3)
	return player, nil
}
