#!/usr/bin/env python3

import json
import sys
from poketes import pokes
from attacks import attacks

del pokes["__fallback__"]

match sys.argv[1]:
    case "pokes":
        print(json.dumps(pokes))
    case "attacks":
        print(json.dumps(attacks))
    case _:
        print({})
