# The Diamond Problem

The Diamond problem occurs where you have multiple superclasses that come from the same base, that you need
to build your latest child class. For example, when you've got `Base` which is the base for `BaseChild1`, and
`BaseChild2`, and want to reference both in one more class: `LatestChild`:
```
class Base{};
class BaseChild1 : public Base{};
class BaseChild2 : public Base{};

class Lastchild : public BaseChild1, public BaseChild2
```


See for a solution [virtual\_inhiterance.md]( at Virtual Inhiterance.)
