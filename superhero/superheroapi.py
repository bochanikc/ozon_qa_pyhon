import random
import requests

ACCESS_TOKEN = 2042120175955355


def get_superhero():
    number = random.randint(1, 700)
    response = requests.get(f'https://superheroapi.com/api/{ACCESS_TOKEN}/{number}')
    if response.status_code == 200:
        res_json = response.json()
        print('superheho: ', res_json['name'])
        return res_json
    else:
        print('Error')


def get_superhero_img_url(superhero):
    superhero = superhero
    superhero_img_url = superhero['image']['url']
    return superhero_img_url


def get_superhero_name(superhero):
    superhero = superhero
    superhero_name = superhero['biography']['full-name']
    superhero_name = superhero_name.replace(' ', '_')
    return superhero_name


def save_picture_on_disc(superhero):
    superhero_img_url = get_superhero_img_url(superhero)
    superhero_name = get_superhero_name(superhero)
    # response = requests.get(superhero_img_url)
    response = requests.get('https://i08.fotocdn.net/s107/5a11941e18370178/public_pin_l/2356233773.jpg')
    # www.superherodb.com умер(
    if response.status_code == 200:
        with open(f'{superhero_name}.jpg', 'wb') as f:
            f.write(response.content)
            print('image download')
    else:
        print('Error image download')
