from superhero import superheroapi

if __name__ == '__main__':
    super_hero = superheroapi.get_superhero()
    superheroapi.save_picture_on_disc(super_hero)
