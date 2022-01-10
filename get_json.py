#!/usr/bin/env python3

import json
from poketes import pokes

del pokes["__fallback__"]

print(json.dumps(pokes))
