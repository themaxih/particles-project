package main

import (
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
1 - configGravity.json
2 - configLifeSpan.json
3 - configOptiMemoire.json
4 - configColor.json [! Epilepsie]
5 - configFR.json
6 - configRPS.json
7 - configSouris.json
8 - configImage.json
Voici quelques liens à tester dans le fichier configImage.json:
- https://upload.wikimedia.org/wikipedia/en/thumb/a/ae/Flag_of_the_United_Kingdom.svg/640px-Flag_of_the_United_Kingdom.svg.png
- https://web.gcompostela.org/wp-content/uploads/2019/02/Universit%C3%A9-de-Nantes.png
- https://opengameart.org/sites/default/files/landscape.png
9 - configGIF.json
Voici quelques liens à tester dans le fichier configGIF.json:
- https://i0.wp.com/www.printmag.com/wp-content/uploads/2021/02/4cbe8d_f1ed2800a49649848102c68fc5a66e53mv2.gif?fit=476%2C280&ssl=1
- https://images.ctfassets.net/l3l0sjr15nav/NCqT6EmgLiydETEgvagXG/24e1571ad345881afc6e775bf6f86a82/200611-EN-GIF-GIFs-Rule.gif
- https://media.tenor.com/mbjDoHXX-T0AAAAC/finn-the-human-adventure-time.gif
*/

func main() {

	config.Get("configSouris.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
