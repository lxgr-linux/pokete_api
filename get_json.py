#!/usr/bin/env python3

import json
import sys
from pokete_poketes import pokes
from pokete_attacks import attacks
from pokete_types import types

del pokes["__fallback__"]

match sys.argv[1]:
    case "pokes":
        print(json.dumps(pokes))
    case "attacks":
        print(json.dumps(attacks))
    case "types":
        print(json.dumps(types))
    case _:
        print({})
