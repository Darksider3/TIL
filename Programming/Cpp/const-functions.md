# Const functions

Functions with the keyword `const` after them signals a function that doesn't even request to change a class members
value. For example
```cpp
class Ex
{
private:
  int field;
public:
  Ex() const
  {
    field = 123;
  }
};

int main()
{
  const T b();
}

```
would throw an error, because we're accessing an element and even modify it.

This wouldn't throw:
```cpp
class Ex
{
private:
  int field;
public:
  Ex(int val)
  {
    field=val;
  }
  int getField() const
  {
    return field;
  }
};


int main()
{
  T b(123);
  std::cout << b.getField();
}
```
