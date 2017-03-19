Since [http://www.python.org/dev/peps/pep-3107/](PEP 3107), we can "require" types in functions!
However, this one is still true:
> Python does not have variables, like other languages where variables have a type and a value; it has names pointing to objects, which know their type.

But, anyways, here is some function `f` that requires an `list`:
```python
def f(var: list, index: int = 0):
	print(var[index])
```

It won't throw an error. That's obvious, because the famous quote:
>We're all consenting adults here

BUT, we can "parse" the code more easily, without reading the function itself.
It enables us to write "documentation" on the function without to "document" it.

It's pretty neat, but we can also define the function like this, which expects to return an integer:
```python
def f(var: list, index: int=0) -> int:
	print(var[index])
```

**Dont forget: You won't get a `TypeError` exception anyway, but for quick reading trough code this could be pretty useful. Nothing else!**