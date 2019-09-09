# Difflib

Difflib is a standard python library that enables different types of visual diffs.
One of that are unified\_diffs with a helper function from multi-line `string`s:
```
import sys
def _unidiff_output(expected, actual):
    """
    Helper function. Returns a string containing the unified diff of two multiline strings.
    """

    import difflib
    
    #expected = [line for line in expected.split('\n')]
    #actual = [line for line in actual.split('\n')]
    #diff=difflib.unified_diff(expected, actual)
    #d = difflib.Differ()
    result = list(difflib.unified_diff(expected.splitlines(keepends=True), actual.splitlines(keepends=True), fromfile="tmpBefore", tofile="tmpAfter"))
    #result = list(d.compare(expected.splitlines(keepends=True), actual.splitlines(keepends=True)))
    return result

s1 = "Test123"
s2 = "Test\nHahaha"

print(sys.stdout.writelines(_unidiff_output(s1, s2)))
```

Which could output something like:
```
--- tmpBefore
+++ tmpAfter
@@ -1 +1,2 @@
-Test123+Test
+HahahaNone
```
