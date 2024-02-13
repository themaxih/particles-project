package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)

// Ce script est relativement le même que le script "createParticle"

func (p *Particle) respawn() {
	// POSITION

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ParticleX int
	var ParticleY int


	if config.General.RandomSpawnX {
		ParticleX = random.Int() % config.General.WindowSizeX
	} else {
		if config.General.SpawnX > config.General.WindowSizeX {
			ParticleX = config.General.WindowSizeX
		} else {
			ParticleX = config.General.SpawnX
		}

		if config.General.SpawnX < 0 {
			ParticleX = 0
		}
	}
	if config.General.RandomSpawnY {
		ParticleY = random.Int() % config.General.WindowSizeY
	} else {
		if config.General.SpawnY > config.General.WindowSizeY {
			ParticleY = config.General.WindowSizeY
		} else {
			ParticleY = config.General.SpawnY
		}

		if config.General.SpawnY < 0 {
			ParticleY = 0
		}
	}

	// SPEED

	var ParticleSpeedX float64 = random.Float64() * config.General.SpeedMultiplier
	var ParticleSpeedY float64 = random.Float64() * config.General.SpeedMultiplier

	if random.Int() % 2 == 0 {
		ParticleSpeedX = -ParticleSpeedX
	}
	if random.Int() % 2 == 0 {
		ParticleSpeedY = -ParticleSpeedY
	}

	p.PositionX = float64(ParticleX)
	p.PositionY = float64(ParticleY)

	if config.General.RandomSpawnOpacity { 
		p.Opacity = float64(random.Int() % 100) / 100 
	} else {
		p.Opacity = 1
	}

	p.KillState = 0
	p.LifeSpan = config.General.LifeSpan
	p.SpeedX = ParticleSpeedX
	p.SpeedY = ParticleSpeedY
}