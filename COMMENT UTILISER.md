PROJET PARTICULES

Le projet particule à pour but de créer un sytème de particules en 2D qui permet de réaliser plusieurs simulations graphiques d'effets.

## Les différentes fonctionnalités

### Les particules

Il faut d'abord savoir qu'une particule possède plusieurs caractéristiques :
    
    - Une couleur
    - Une taille
    - Un angle de rotation
    - Une transparence
    - Une position

Certaines de ces caractéristique peuvent faire évoluer la particule au fil du temps :
    
    - Sa vitesse
    - Son accélération
    - Sa durée de vie
    - Sa capacité d'interaction avec les autres particules

Les particules évoluent dans un système qui va en générer de nouvelles, mettre à jour les particules, en supprimer, ...

### Comment utiliser les différents pre-sets

Avant de vous parlez des différentes fontionnalités, il faut savoir comment les utiliser :

    - Allez dans le fichier "main.go"
    - Modifiez la ligne "config.Get("config.json")" ( l.34 ) en remplacent "config.json" par 1 des 9 pre-sets disponnibles ( dans le même fichier de la ligne 13 à 29 )
    - Il y a aussi possibilité d'utiliser et de modifier directement le fichier "config.json" pour jouer avec les caractéristiques des particules.
    - ATTENTION ! Si vous utilisez les pre-sets "configGIf.json" ou "configImage.json" pensez à laisser "gif-..." ou "img-..." avant le liens de l'image ou du GIF dans la ligne "Preset" ( l.31 ) des fichiers correspondant pour que cela fonctionne. Il faut obligatoirement que l'image soit en format PNG ! Les GIF et images ne doivent pas dépasser les 1920 * 1080 pixels en taille, sinon, vous ne pourrez voir l'image en entier.
    - Voila, vous pouvez maintenant tester les différents pre-sets !

### Les différents pre-sets

Notre système comporte plusieurs fonctionnalités qui utilisent les particules pour éffectuer plusieurs simulations :
    
    - Une simulation où les particules sont affectés par la gravité, cela permet de représenter une cascade d'eau ou de la pluie. Vous pouvez modifier des valeurs dans le fichier "configGravity.json" comme le "SpeedMultiplier" ou le "ParticleGravity" pour jouer avec la gravité des particules.
        Pour activer ce mode, il faut utiliser le pre-set "configGravity.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation où les particules ont une durée de vie et donc sont supprimées au bout d'un moment.
        Pour activer ce mode, il faut utiliser le pre-set "configLifeSpan.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation où plusieurs dizaines de milliers de particules en mouvement peuvent apparaîtrent sans causer de ralentissements.
        Pour activer ce mode, il faut utiliser le pre-set "configOptiMemoire.json' dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation où la couleur des particules change au fil du temps ( 1 fois par "Tick" ).
        Pour activer ce mode, il faut utiliser le pre-set "configColor.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation qui affiche le drapeau français. La couleur des particules change selon leur coordonnée : les particules qui sont dans le premier tier de l'écran sont bleus, celles du deuxième tier sont blanches et celles du troisième tier sont rouges.
        Pour activer ce mode, il faut utiliser le pre-set "configFR.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation du jeu "Pierre, Feuille, Ciseau" ( Rock, Paper, Scissors ) où ici les objets sont remplacé par des particules de 3 couleurs, rouge, bleu et vert où les rouges "mangent" les verts qui eux "mangent" les bleus qui eux "mangent" les rouges. Quand des particules vont entrer en collision, chacune va vérifier sa couleur et la couleur de l'autre et la particule qui possède la couleur gagnante va changer l'autre en sa couleur. Le but est donc de voir à la fin quelle sera l'unique couleur restante.
        Pour activer ce mode, il faut utiliser le pre-set "configRPS.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simualtion où les particules sont attirées par le pointeurs de la souris quand le clic gauche est enfoncé.
        Pour activer ce mode, il faut utiliser le pre-set "configSouris.json' dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

    - Une simulation qui affiche une image, il suffit d'entrer l'adresse d'une image png dans le configImage.json et la fenêtre se mettera automatiquement à la taille de l'image choisi et celle-ci apparaîtera, vous pouvez aller chercher le liens de n'importe quelle image mais on vous en a proposé 3 en commentaire dans le fichier "main.go".
        Pour activer ce mode, il faut utiliser le pre-set "configImage.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).
        
    - On a fait la même simulation mais cettelle affiche un GIF, pareil que pour l'image, il faut entrer l'adresse d'un GIF dans le "configGIF.json" et la fenêtre se mettera automatiquement à sa taille et le GIF apparaîtera, vous pouvez aller chrecher le liens de nimporte quelle GIF mais on vous en a proposé 3 en commentaire dans le fichier "main.go".
        Pour activer ce mode, il faut utiliser le pre-set "configGIF.json" dans le fichier "main.go" à la ligne "config.Get("config.json")" ( l.34 ).

N'oubliez pas que vous pouvez aussi utiliser le fichier "config.json" pour choisir vous même vos paramètres de particules et tester différentes choses.

## Les librairies utilisées

Pour ce projet, nous avons eu besoin de plusieurs librairies :

    - La librairie "math" qui permet d'utiliser des formules mathématiques et notament d'utiliser l'aléatoire avec "math/rand"
    - La librairie "time" qui permet de vérifier l'aléatoire.
    - la librairie "net" qui permet de faire des requètes html avec "net/html" qui permettent d'utiliser des liens Google d'image ou de GIF.
    - La librairie "image" qui permet de décoder des images pour récupérer les coordonnés des pixels, leur couleur, ...
    - La librairie "ebiten" qui la librairie principale qui permet notament de gérer l'interface.
    - La librairie "container" qui ets utilisé pour lister les particules.
    - La librairie "testing" pour effectuer les tests.

**Nous vous remercions de votre attention et vous souhaitons une agréable journée**
