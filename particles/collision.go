package particles

import (
	"math"
)


// Cette fonction compare les coordonnées d'une particule avec une autre
func (p *Particle) Collides(other *Particle) bool {
	// Renvoie la différence entre les coordonnées d'une particule et d'une autre, et donc si le résultat est inférieure à 10 alors cela vuet dire que les particules se touchent
	return int(math.Sqrt(float64((p.PositionY-float64(other.PositionY))*float64((p.PositionY-float64(other.PositionY)))+float64((p.PositionX-float64(other.PositionX)))*float64((p.PositionX-float64(other.PositionX)))))) < 10 // On vérifie si c'est inférieur à 10 car la particule à une taille de base de 10 px.
}
