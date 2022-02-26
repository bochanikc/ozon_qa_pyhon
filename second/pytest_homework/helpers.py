import random
import requests
from lxml import html


def gen_param():
    a = random.random()
    b = random.random()
    c = a + b
    return a, b, c

def get_random_id():
    return random.randint(1, 701)

def get_heroes():
    heroes = []
    result = requests.get(url='https://www.superheroapi.com/ids.html').text
    result = html.fromstring(result)
    data = result.xpath('//table//tr')
    print(len(data))
    for step in data:
        heroes.append([step.xpath('./td')[0].text, step.xpath('./td')[1].text])
    print(heroes)
    return heroes


def get_woman_ids():
    women = []
    heroes = get_heroes()
    for hero in heroes:
        if 'woman' in hero[1].lower():
            women.append(hero[0])
    return women
