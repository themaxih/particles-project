package particles

import (
	"project-particles/config"
)

func (s *System) UpdateMode() {
	var mode string = config.General.Preset

	switch {
	case mode == "fr":
		// Quand le mode est "fr", on va récupérer la position en X de chaque particules et modifier leur couleur en bleu, blanc ou rouge en fonction du tier où ils se trouvent
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			
			if int(p.PositionX) <= config.General.WindowSizeX/3 { // Bleu
				p.ColorRed = 0
				p.ColorGreen = 0
				p.ColorBlue = 1
			} else if int(p.PositionX) <= (config.General.WindowSizeX/3)*2 { // Blanc
				p.ColorRed = 1 
				p.ColorGreen = 1
				p.ColorBlue = 1
			} else { // Rouge
				p.ColorRed = 1
				p.ColorGreen = 0
				p.ColorBlue = 0
			}
		}

	case len(mode) > 4 && mode[:4] == "img-":
		// Lorsque le mode est "img-" (image), on va changer la couleur de chaques particules pour que celle-ci corresponde à la couleur au même endroit sur l'image
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			// Récupérer les valeurs RGB de la couleur à la même position ainsi que l'opacité du pixel
			r, g, b, a := s.img.At(int(p.PositionX+4), int(p.PositionY+4)).RGBA()

			p.ColorRed = float64(uint8(r)) / 255
			p.ColorGreen = float64(uint8(g)) / 255
			p.ColorBlue = float64(uint8(b)) / 255
			p.Opacity = float64(uint8(a)) / 255
		}
	case len(mode) > 4 && mode[:4] == "gif-":
		// Lorque le mode est "gif-" (image animée), le processus va être le même, cependant tous les n ticks on va changer l'image pour la prochaine du GIF
		if s.Tick%(s.gif.Delay[s.PresetIndex]+1) == 0 {
			if s.PresetIndex == len(s.gif.Image)-1 {
				s.PresetIndex = 0
			} else {
				s.PresetIndex++
			}
			s.img = s.gif.Image[s.PresetIndex]
		}
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			// Récupérer les valeurs RGB de la couleur
			r, g, b, a := s.img.At(int(p.PositionX+4), int(p.PositionY+4)).RGBA()
			p.ColorRed = float64(uint8(r)) / 255
			p.ColorGreen = float64(uint8(g)) / 255
			p.ColorBlue = float64(uint8(b)) / 255
			p.Opacity = float64(uint8(a)) / 255
		}
	}
}
