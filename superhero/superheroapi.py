import random
import requests


def get_superhero():
    ACCESS_TOKEN = 2042120175955355
    number = random.randint(1, 700)
    response = requests.get(f'https://superheroapi.com/api/{ACCESS_TOKEN}/{number}')
    if response.status_code == 200:
        res_json = response.json()
        image_url = res_json['image']['url']
        #response = requests.get(image_url)
        response = requests.get('https://i08.fotocdn.net/s107/5a11941e18370178/public_pin_l/2356233773.jpg')
        print(response.content)
        with open('.', 'wb') as f:
            f.write(response.content)
    else:
        print('Error')


def save_picture_on_disc():
    hero = get_superhero()


get_superhero()