package particles

import (
	"fmt"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	l := s.Content // On récupère la liste de particules
	s.Tick++

	s.CursorX, s.CursorY = ebiten.CursorPosition()
	if s.CursorX < 0 {
		s.CursorX = 0
	}
	if s.CursorY < 0 {
		s.CursorY = 0
	}
	if s.CursorX > config.General.WindowSizeX {
		s.CursorX = config.General.WindowSizeX
	}
	if s.CursorY > config.General.WindowSizeY {
		s.CursorY = config.General.WindowSizeY
	}

	if config.General.EventOnClick {
		s.onClick()
	}

	if l.Len() < config.General.MaxParticles || config.General.MaxParticles == -1 { // MaxParticles = -1 veut dire qu'il n'y a pas ed limite de particule
		if config.General.SpawnRate < 1 && config.General.SpawnRate > 0 { // Pouvoir générer par exemple 1 particule tout les 2 Ticks avec "SpawnRate = 0.5"
			if s.Tick%int(1/config.General.SpawnRate) == 0 {
				s.CreateParticle()
			}
		} else {
			for i := 0; i < int(config.General.SpawnRate) && (l.Len() < config.General.MaxParticles || config.General.MaxParticles == -1); i++ {
				s.CreateParticle()
			}
		}
	}

	// Actualiser l'affichage du nombre de TPS toute les secondes
	if s.Tick%60 == 0 {
		ebiten.SetWindowTitle("Project particles - Particules: " + fmt.Sprint(l.Len()) + " - TPS: " + fmt.Sprint(int(ebiten.CurrentTPS())))
	}

	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)

		// Update
		p.updateParticle()

		// KILL PARTICLES

		// Réduire la durée de vie de la particule tout les Ticks
		p.LifeSpan--
		if p.LifeSpan == 0 {
			p.hideParticle()
			p.KillState = 1
			continue
		// SI les particules sortent de l'écran, on appel la fonction "HideParticle()" et le KIllState de la particule passe à 1
		} else if p.PositionX < -20 || p.PositionY < -20 {
			p.hideParticle()
			p.KillState = 1
			continue
		} else if int(p.PositionY) > config.General.WindowSizeY+20 || int(p.PositionX) > config.General.WindowSizeX+20 {
			p.hideParticle()
			p.KillState = 1
			continue
		}
	}

	// On parcour la liste de toutes les particules et si une particule à un KIllState de 1, elle se supprime ( appel de la fonction "Remove()" )
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.KillState == 1 {
			if config.General.ToggleRespawn {
				p.respawn()
			} else {
				l.Remove(e)
			}

		}
	}

	if config.General.RockPaperScissors {
		// Chaque particule de la liste de particules parcours la liste de particules pour savoir quand elles entrent en collsion avec d'autres particules ( appel de la fonction "Collides()")
		for e := l.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			for e2 := l.Front(); e2 != nil; e2 = e2.Next() {
				p2 := e2.Value.(*Particle)
				if p.Collides(p2) && p2 != p {
					// Si un particule rouge touche une particule bleu, la particule rouge devient bleu 
					if p.ColorRed == 1 && p.ColorBlue == 0 && p.ColorGreen == 0 && p2.ColorRed == 0 && p2.ColorBlue == 1 && p2.ColorGreen == 0 {
						p.ColorRed = 0
						p.ColorBlue = 1
						p.ColorGreen = 0
					// Si un particule bleu touche une particule verte, la particule bleu devient verte
					} else if p.ColorRed == 0 && p.ColorBlue == 1 && p.ColorGreen == 0 && p2.ColorRed == 0 && p2.ColorBlue == 0 && p2.ColorGreen == 1 {
						p.ColorRed = 0
						p.ColorBlue = 0
						p.ColorGreen = 1
					// Si un particule verte touche une particule rouge, la particule verte devient rouge 
					} else if p.ColorRed == 0 && p.ColorBlue == 0 && p.ColorGreen == 1 && p2.ColorRed == 1 && p2.ColorBlue == 0 && p2.ColorGreen == 0 {
						p.ColorRed = 1
						p.ColorBlue = 0
						p.ColorGreen = 0
					}
				}
			}
		}
	}

	s.UpdateMode()
}
