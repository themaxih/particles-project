package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

func (s *System) CreateParticle() {
	// POSITION

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ParticleX int
	var ParticleY int

	// Si la générationaléatoire en X est activé, alors assigner un nombre aléatoire qui ne dépasse pas la taille de l'écran à la variable "ParticleX", sinon, assigner "ParticleSpawnX" à "ParticleX"
	if config.General.RandomSpawnX {
		ParticleX = random.Int() % (config.General.WindowSizeX - 11)
	} else {
		// Faire en sorte que quand on rentre une coordonné plus grande que la taille de l'écran, elle soit automatiquemet réduite à la limite de la taille de l'écran.
		if config.General.SpawnX > config.General.WindowSizeX {
			ParticleX = config.General.WindowSizeX
		} else {
			ParticleX = config.General.SpawnX
		}

		if config.General.SpawnX < 0 {
			ParticleX = 0
		}
	}

	// Même chose pour "ParticleY"
	if config.General.RandomSpawnY {
		ParticleY = random.Int() % (config.General.WindowSizeY - 11)
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
	var ParticleSpeedX float64 = random.Float64() * config.General.SpeedMultiplier // Générer une valeur aléatoire
	var ParticleSpeedY float64 = random.Float64() * config.General.SpeedMultiplier

	// Si un nombre généré ets paire, inverser la vitesse de la particule
	if random.Int()%2 == 0 {
		ParticleSpeedX = -ParticleSpeedX
	}
	if random.Int()%2 == 0 {
		ParticleSpeedY = -ParticleSpeedY
	}

	// COLOR

	// Couleur de base de la particle
	var ParticleColorR float64 = 0
	var ParticleColorG float64 = 0
	var ParticleColorB float64 = 0

	// Activer le fait qu'une particle apparaise avec une couleur aléatoire
	if config.General.RandomSpawnColor {
		ParticleColorR = float64(random.Int()%100) / 100
		ParticleColorG = float64(random.Int()%100) / 100
		ParticleColorB = float64(random.Int()%100) / 100
	} else {
		switch {

		// Différent pre-sets de couleurs	
		case s.ParticleColor == "white":
			ParticleColorR = 1
			ParticleColorG = 1
			ParticleColorB = 1
		case s.ParticleColor == "red":
			ParticleColorR = 1
		case s.ParticleColor == "blue":
			ParticleColorB = 1
		case s.ParticleColor == "green":
			ParticleColorG = 1
		case s.ParticleColor == "aqua":
			ParticleColorG = 1
			ParticleColorB = 1
		case s.ParticleColor == "yellow":
			ParticleColorR = 1
			ParticleColorG = 1
		case s.ParticleColor == "magenta":
			ParticleColorR = 1
			ParticleColorB = 1
		case s.ParticleColor == "orange":
			ParticleColorR = 1
			ParticleColorG = 0.65
		case s.ParticleColor == "random":
			ParticleColorR = float64(random.Int()%100) / 100
			ParticleColorG = float64(random.Int()%100) / 100
			ParticleColorB = float64(random.Int()%100) / 100
		}
	}

	// Si le mode "Pierre, feuille, ciseau" est activé, faire en sorte que le premier tier des particules soit rouge, le deuxième soit vert et le troisième soit bleu.
	if config.General.RockPaperScissors {
		if s.ParticleNumber <= config.General.InitNumParticles/3-1 {
			ParticleColorR = 1
			ParticleColorG = 0
			ParticleColorB = 0
		}
		if s.ParticleNumber > config.General.InitNumParticles/3-1 && s.ParticleNumber <= (config.General.InitNumParticles/3)+config.General.InitNumParticles/3 {
			ParticleColorR = 0
			ParticleColorG = 1
			ParticleColorB = 0
		}
		if s.ParticleNumber > config.General.InitNumParticles/3+config.General.InitNumParticles/3-1 {
			ParticleColorR = 0
			ParticleColorG = 0
			ParticleColorB = 1
		}
	}

	var ParticleOpacity float64 = 1

	if config.General.RandomSpawnOpacity {
		ParticleOpacity = float64(random.Int()%100) / 100 // Mettre une opacité aléatoire
	}

	// PUSH_FRONT
	s.Content.PushFront(&Particle{
		PositionX: float64(ParticleX),
		PositionY: float64(ParticleY),

		SpeedX:    ParticleSpeedX,
		SpeedY:    ParticleSpeedY,

		ScaleX:    1, ScaleY: 1, // Taille de base = 1

		ColorRed: ParticleColorR, ColorGreen: ParticleColorG, ColorBlue: ParticleColorB,

		Opacity: ParticleOpacity, LifeSpan: config.General.LifeSpan,

		KillState: 0, // Etat de la particule : 0 -> Vivante | 1 -> Morte
	})
}
