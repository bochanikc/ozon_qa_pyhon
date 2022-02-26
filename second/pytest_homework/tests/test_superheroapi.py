from second.pytest_homework.helpers import get_woman_ids, gen_param
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


def test_two_heroes_stronger(who_stronger):
    who_stronger
    # id1, id2, hero = who_stronger


@pytest.mark.parametrize('a,b,c', [gen_param() for _ in range(10)])
def test_sum(a, b, c):
    assert a + b == c
