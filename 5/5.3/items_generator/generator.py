import json
import random
from faker import Faker


def get_article():
    fake = Faker()
    category = get_category()
    article = fake.bothify(text=f'OZON_{category}_??##')
    return article


def get_category():
    categories = ['outerwear', 'shirt', 'socks']
    return random.choice(categories)


def get_weight(category):
    if category == 'outerwear':
        return random.randint(1000, 10000)
    elif category == 'shirt':
        return random.randint(100, 2000)
    elif category == 'socks':
        return random.randint(10, 300)


def get_type(category):
    if category == 'outerwear':
        types_outerwear = ['coat', 'jacket', 'windbreaker']
        return random.choice(types_outerwear)
    elif category == 'shirt':
        types_shirt = ['casual', 'dress', 'loosefit']
        return random.choice(types_shirt)


def get_colour():
    fake = Faker()
    color = fake.color_name()
    return color


def get_textile():
    materials = ['cotton', 'silk', 'polyester', 'viscose', 'acrylic']
    persents = 100
    material = random.choice(materials)
    persent = random.randint(10, 100)
    textile = [
        {
            "matherial": material,
            "percent": persent
        }
    ]
    return textile


def get_description():
    fake = Faker()
    return fake.text()


def get_items(num):
    items = []
    for i in range(1, num + 1):
        id = i
        article = get_article()
        category = get_category()
        weight = get_weight(category)
        textile = get_textile()
        description = get_description()
        item = {
            "id": id,
            "article": article,
            "category": category,
            "weight": weight,
            "textile": textile,
            "description": description
        }
        items.append(item)
    json_items = json.dumps(items)
    return json_items
