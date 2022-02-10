import superheroapi as sh

if __name__ == '__main__':
    super_hero = sh.get_superhero()
    sh.save_picture_on_disc(super_hero)
