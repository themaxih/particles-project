package particles

import (
	"container/list"
	"image"
	"image/gif"
	"net/http"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

// "NewSystem()" accède à la structure "Système" pour pouvoir définir les caractéristiques des particules.
func NewSystem() System {
	var s System
	s.Tick = 0
	s.Content = list.New()
	s.PresetType = "None"
	s.PresetIndex = 0

	if len(config.General.Preset) > 4 && (config.General.Preset[:4] == "img-" || config.General.Preset[:4] == "gif-") {
		// Faire la requète HTML
		url := config.General.Preset[4:]
		response, _ := http.Get(url)
		
		// Décoder l'image
		if config.General.Preset[:4] == "gif-" {
			s.PresetType = "gif"
			gif, _ := gif.DecodeAll(response.Body)
			s.gif = *gif
			s.img = gif.Image[0]

		} else {
			s.PresetType = "img"
			img, _, _ := image.Decode(response.Body)
			s.img = img
		}

		// Récupérer la taille de l'image
		imgBounds := s.img.Bounds()

		// On change la taille de la fenêtre par rapport à la taille de l'image
		config.General.WindowSizeX = imgBounds.Max.X
		config.General.WindowSizeY = imgBounds.Max.Y
		ebiten.SetWindowSize(imgBounds.Max.X, imgBounds.Max.Y)

		// On change le nombre de particule initial et le nombre de particule maximal selon la taille de l'image pour qu'elle soit le plus détaillé possible ( éviter de laisser des "trous" )
		config.General.MaxParticles = (imgBounds.Max.X * imgBounds.Max.X) / 5
		config.General.InitNumParticles = (imgBounds.Max.X * imgBounds.Max.X) / 1000
		if config.General.MaxParticles > 175000 {
			config.General.MaxParticles = 175000
			config.General.InitNumParticles = config.General.MaxParticles / 1000
		}
		config.General.SpawnRate = float64((imgBounds.Max.X * imgBounds.Max.X) / 1000)
	} else {
		s.img = image.NewRGBA(image.Rect(0, 0, 10, 10))
	}

	s.ParticleSpawnX = config.General.SpawnX
	s.ParticleSpawnY = config.General.SpawnY
	s.ParticleColor = config.General.DefaultSpawnColor

	s.ParticleNumber = 0

	if config.General.RockPaperScissors {
		for nbParticles := 0; nbParticles < config.General.InitNumParticles; nbParticles++ {
			s.CreateParticle()
			s.ParticleNumber++ // Sert à pouvoir connaitre le nombre de particules qui sont déjà apparu pour pouvoir définir leur couleur
		}
	} else {
		// Générer des particules tant que le nombre de particule initials n'a pas été atteind
		for nbParticles := 0; nbParticles < config.General.InitNumParticles; nbParticles++ {
			s.CreateParticle()
		}
	}
	return s
}
