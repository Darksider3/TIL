Python 3(atleast?) got a neat `__call__` method for classes, which you can use to call a class like a function!

This is awesome, because now it's possible for us to treat a class like a function.
Example class:
```python
class example:
  def __init__(self, say=u"Hello World!"):
    self.phrase=say
    print(u"__init__(self, say=u\"{}\")".format(self.phrase))
  
  def say(self):
    print(self.phrase)
x=example()
x.say()
```
The output is simple:
```python
__init__(self, say=u"Hello World!")
```
## Now adding `__call__` to call the instance of the object!

We can now add `__call__` to call the instance of the object as a **function** itself:
```python
class example:
  def __init__(self, say=u"Hello World!"):
    self.phrase=say
    print(u"__init__(self, say=u\"{}\")".format(self.phrase))
  
  def say(self):
    print(self.phrase)
  
  def __call__(self, say):
    self.phrase=say
    print(u"__call__(self, say=u\"{}\")".format(self.phrase))
    self.say()

xmpl=example()
xmpl.say()
xmpl(u"new Hello World!")
```

Output:
```python
__init__(self, say=u"Hello World!")
Hello World!
__call__(self, say=u"new Hello World!")
new Hello World!
```