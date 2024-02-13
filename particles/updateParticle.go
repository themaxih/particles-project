package particles

import (
	"project-particles/config"
)

func (p *Particle) updateParticle() {
	p.PositionX += p.SpeedX

	// Si le système de gravité est activé, multiplier la vitesse en Y par la force de gravité
	if config.General.ParticleGravity == 0 {
		p.PositionY += p.SpeedY
	} else {
		p.PositionY += p.SpeedY
		p.SpeedY *= config.General.ParticleGravity
	}

	if config.General.ToggleVanish {
		// Réduire la taille et l'oppacité de la particule
		p.Opacity = float64(p.LifeSpan) / float64(config.General.LifeSpan)
		p.ScaleX = float64(p.LifeSpan) / float64(config.General.LifeSpan)
		p.ScaleY = float64(p.LifeSpan) / float64(config.General.LifeSpan)
	}

	if config.General.ToggleBounce {
		p.bounceParticle()
	}
}
