# Virtual Inhiterance

Due to C++s feature of multiple inhiterances, it can cause the so-called "diamond"-Problem: When two superclasses
of a class have a common base class, that the common class get's called twice.
To circumvent that, you can define the inhiterance as virtual in that two superclasses, which leaves exactly one,
but not two, copies of the base class.
Problem example:
```c++
#include <iostream>

class Person {
public:
	Person(int x) { std::cout << "Person::Person(int x) called << std::endl; }
};

class Faculty : public Person {
public:
	Faculty(int x): Person(x) { std::cout << "Faculty::Faculty(int x) called" << std::endl; }
};

class Student : public Person {
public:
	Student(int x): Person(x) { std::cout << "Student::Student(int x) called" << std::endl; }
};

class TA : public Faculty, public Student {
public:
	TA(int x): Faculty(x), Student(x) { std::cout << "TA::TA(int x) called" << std::endl; }
};

int main()
{
	TA tal(309);
}
```

Output:

```
Person::Person(int x) called
Faculty::Faculty(int x) called
Person::Person(int x) called
Student::Student(int x) called
TA::TA(int x) called
```
So the Person::Person get's called twice. To circumvent, we can define the relationship as virtual in that two 
superclasses:
```c++

```
#include <iostream>

class Person {
public:
        Person(int x) { std::cout << "Person::Person(int x) called << std::endl; }
	Person() {std::cout << "Person::Person() called" << std::endl; }
};

class Faculty : virtual public Person {
public:
        Faculty(int x): Person(x) { std::cout << "Faculty::Faculty(int x) called" << std::endl; }
};

class Student : virtual public Person {
public:
        Student(int x): Person(x) { std::cout << "Student::Student(int x) called" << std::endl; }
};

class TA : public Faculty, public Student {
public:
        TA(int x): Faculty(x), Student(x) { std::cout << "TA::TA(int x) called" << std::endl; }
};

int main()
{
        TA tal(309);
}
```

Out:

```
Person::Person() called
Faculty::Faculty(int x) called
Student::Student(int x) called
TA::TA(int x) called
```

As you can see, the grandparents parameterized constructor doesnt get called. Instead, the **default** constructor, __Person::Person()__
gets called. So you are loosing the **second** copy of the base class! Yay!
