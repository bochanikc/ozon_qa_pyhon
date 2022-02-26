import pytest
from helpers import get_random_id
from second.pytest_homework.clients.api import ApiClient

@pytest.fixture()
def get_hero_by_id(id):
    return ApiClient.get_hero(id)

@pytest.fixture()
def get_random_hero_powerstats(get_powerstats_by_id):
    while True:
        result = get_powerstats_by_id(id=get_random_id())
        if result is None:
            continue
        return result

@pytest.fixture()
def who_stronger(get_random_hero_powerstats):
    print(get_random_hero_powerstats)
    return get_random_hero_powerstats

