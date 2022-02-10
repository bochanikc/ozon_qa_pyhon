import json
import random

import generator

if __name__ == '__main__':
    json_data = generator.get_items(random.randint(1, 21))
    print(json.dumps(json_data, indent=4, sort_keys=False))
