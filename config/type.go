package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawnX             bool
	RandomSpawnY             bool
	RandomSpawnColor         bool
	DefaultSpawnColor        string
	RandomSpawnOpacity       bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	LifeSpan                 int
	SpeedMultiplier          float64
	ParticleGravity          float64
	MaxParticles             int
	ToggleRespawn            bool
	ToggleVanish             bool
	ToggleEpilepsie          bool
	EventOnClick             bool
	Preset                   string
	ToggleBounce             bool
	RockPaperScissors        bool
}

var General Config
