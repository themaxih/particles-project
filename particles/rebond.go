package particles

import (
	"project-particles/config"
)

func (p *Particle) bounceParticle() {
	// Pouvoir rebondire sur les côtés de la fenêtre
	if int(p.PositionX) >= config.General.WindowSizeX-10 {
		p.SpeedX = -p.SpeedX
	}

	if int(p.PositionX) <= 0 {
		p.SpeedX = -p.SpeedX
	}

	if int(p.PositionY) >= config.General.WindowSizeY-10 {
		p.SpeedY = -p.SpeedY
	}

	if int(p.PositionY) <= 0 {
		p.SpeedY = -p.SpeedY
	}
}
