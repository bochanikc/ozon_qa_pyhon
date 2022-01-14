import pytest


@pytest.mark.parametrize('id, answer_name',
                         [('1', 'A-Bomb'),
                          ('125', 'Blue Beetle II')])
def test_smoke_superheroapi(get_hero_by_id, id, answer_name):
    result = get_hero_by_id
    assert result['name'] == answer_name
