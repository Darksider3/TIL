# cxxtest - Testing

Last week i discovered [cxxTest](http://cxxtest.com/). It's some testing suite for C++ Files, that generates Runners and registers with Python(previously
with Perl) automatic testing sources.

For example, given File HelloWorld.hh
```cpp
#include <iostream>
class Blub
{
  Blub()
  {
    //nothing
  }
  char *returnMe(char *me)
  {
    return me;
  }
};
```
To generate the runner and generate the tests, we have to write first the tests
To setup our env that we need for the tests, we can use setUp(), which runs before EVERY test, and teatDown() which runs after every.

Lets name it testHello.hh
```cpp
//our source
#include "HelloWorld.hh"
//needed by cxxtest
#include <cxxtest/TestSuite.h>

class BlubTestSuite : public CxxTest::TestSuite
{
  Blub *test;
public:
  void setUp()
  {
    this->test=new Blub();
  }
  void tearDown()
  {
    delete this->test;
  }
 //test functions must have "test" in the name!!
  void test_Me_Dont_Fail(void)
  {
    //Assert the given statement and return error when false, continue if returns true
    TS_ASSERT(this->test.returnMe("Hello!")=="Hello!"); //true, ok. Pass it.
    TS_ASSERT_EQUALS(this->test.returnMe("TS_ASSERT_EQUALS", "TS_ASSERT_DOESNT_EQUALS")); //ERROR, it will spit out it.
  }
};
```
No we only have to run the cxxtestgen command:
```sh
cxxtestgen --error-printer -f -o test.cpp testHello.hh
```
We got now a test.cpp. Thus we've to compile with our "favourite" compiler.
```sh
g++ test.cpp -o test
```
Done! Magic happened. We got already registered tests, dont have to mess with register functions and dont even need to write some
obscure main function that runs everything!
Execute the `test` binary.
Output from a bigger file:

```
Running cxxtest tests (6 tests)...OK!
```

An error looks like:
```
testFile.hpp:62: Error: Assertion failed: "ABC"==true
``` 
