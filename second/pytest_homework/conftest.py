import pytest
from helpers import get_random_id
from second.pytest_homework.clients.api import ApiClient

@pytest.fixture()
def get_hero_by_id(id):
    return ApiClient.get_hero(id)

@pytest.fixture()
def get_data_for_who_stronger():
    hero = []
    count = 0
    while True:
        result = ApiClient.get_powerstats_by_id(get_random_id())
        if result.json()['power'] == 'null':
            continue
        hero.append(result)
        count = count + 1
        if count == 2:
            break
    return hero[0], hero[1]
