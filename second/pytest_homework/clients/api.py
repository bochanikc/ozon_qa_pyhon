from urllib.parse import urljoin
import requests
from second.pytest_homework.constants import BASE_URL, ACCESS_TOKEN

API_URL = urljoin(BASE_URL, 'api/')
IDS_URL = urljoin(BASE_URL, 'ids.html')

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

    def get_powerstats_by_id(id):
        result = ApiClient(base_url=API_URL, access_token=ACCESS_TOKEN).get(
            path=f'/{id}/powerstats'
        )
        print(result.text)
        return result

    def get_hero(id):
        result = ApiClient(base_url=API_URL, access_token=ACCESS_TOKEN).get(
            path=f'/{id}'
        )
        print(result.text)
        return result
