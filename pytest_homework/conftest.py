import random
from urllib.parse import urljoin
import requests
import pytest
from pytest_homework.constants import BASE_URL, ACCESS_TOKEN


class ApiClient:
    """
    Client for work with API
    with GET/POST requests
    """

    def __init__(self, base_url, access_token):
        self.base_url = base_url
        self.access_token = access_token

    def get(self, path="/", params=None, headers=None):
        url = self.base_url + self.access_token + path
        print(url)
        return requests.get(url=url, params=params, headers=headers)

    def post(self, path="/", params=None, headers=None):
        url = self.base_url + self.access_token + path
        return requests.post(url=url, params=params, headers=headers)


API_URL = urljoin(BASE_URL, 'api/')
IDS_URL = urljoin(BASE_URL, 'ids.html')


@pytest.fixture(scope="session")
def api_superheroapi():
    return ApiClient(base_url=API_URL, access_token=ACCESS_TOKEN)


@pytest.fixture()
def get_powerstats_by_id(api_superheroapi, id):
    result = api_superheroapi.get(
        path=f'/{id}/powerstats'
    )
    print(result.text)
    return result


@pytest.fixture()
def get_random_hero_powerstats(get_powerstats_by_id):
    while True:
        result = get_powerstats_by_id(id=get_random_id)
        if result is None:
            continue
        return result


@pytest.fixture()
def get_hero_by_id(api_superheroapi, id):
    result = api_superheroapi.get(
        path=f'/{id}'
    )
    print(result.text)
    return result


@pytest.fixture()
def who_stronger(get_random_hero_powerstats):
    print(get_random_hero_powerstats)
    return get_random_hero_powerstats

@pytest.fixture()
def get_random_id():
    return random.randint(1, 701)
