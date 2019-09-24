# Importing from parent directory(or any other above)

You can either do that by referencing the parent with `...`

```
from ... import moduleName
```

or by hardcoding the include path:
```
import os,sys,inspect

currrendir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe()))
parentdir = os.path.dirname(currendir)
sys.path.insert(0, parentdir) # insert parentdir-variable into pythons path

import module
