from helpers import get_woman_ids
import pytest


@pytest.mark.parametrize('id, answer_name',
                         [('1', 'A-Bomb'),
                          ('125', 'Blue Beetle II')])
# @smoke_check
def test_smoke_superheroapi(get_hero_by_id, answer_name):
    assert get_hero_by_id.json()['name'] == answer_name


@pytest.mark.parametrize('id', get_woman_ids())
def test_woman_is_woman(get_hero_by_id):
    assert get_hero_by_id.json()['appearance']['gender'] == 'Female'


def test_1():
    print(get_woman_ids().text)