# Iterate over Arrays in Ruby

Given following:
``` ruby
array = [1, 2, 3, 4, 5, 6]
array.each { |x| puts x }
```
[Source http://stackoverflow.com/a/310638/4721352](SO Answer)
It iterates trough the `array`-variable.
The `|x|` means here, that it uses `x` as the value. On every rotation, it tries to get the next variable.
Outputs:
``` sh
1
2
3
4
5
6
```

Another variant with key and value stored:
```ruby 
array = ["A", "B", "C"]
array.each_with_index {|val, index| puts "#{val} => #{index}" }
```

Outputs:
``` sh
A => 0
B => 1
C => 2
```

First example in C would be:
``` c
#include <stdio.h>

int main(int argc, char **argv)
{
	int x=[1,2,3,4,5];
  size_t size=4;
	for(int i=0; i<=size; i++)
  {
  	printf("%i\n", x[i]);
  }
}
```
Pretty handy, heh?