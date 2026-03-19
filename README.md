# Project 01 — Go_Bootcamp

**Summary:** In this project, you will get to know the Go language, study the basics of the language, its syntax, and data types: int, string, slice, map, struct. You will work with input/output and learn how to write small Go programs to implement console utilities.

💡 [Click here](https:/\new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) to share your feedback on this project. It’s anonymous and will help our team make learning better. We recommend completing the survey immediately after finishing the project.

## Contents

  - [Chapter I](#chapter-i)
  - [Instructions](#instructions)
  - [Chapter II](#chapter-ii)
    - [General Information](#general-information)
  - [Chapter III](#chapter-iii)
    - [Project «Smart Utilities»](#project-smart-utilities)
    - [Task 1. Console Calculator](#task-1-console-calculator)
    - [Task 2. Most Frequent Words](#task-2-most-frequent-words)
    - [Task 3. Slice Intersection](#task-3-slice-intersection)
    - [Task 4. Visit Log](#task-4-visit-log)


## Chapter I
## Instructions

1. Throughout the course, you will experience a sense of uncertainty and a sharp lack of information — this is normal. Don’t forget that the information in the repository and Google is always at your side. Just like your peers and Rocket.Chat. Communicate. Search. Use common sense. Don’t be afraid to make mistakes.
2. Be careful with your sources. Verify. Think. Analyze. Compare.
3. Read the tasks attentively. Read them several times.
4. It's also better to read the examples attentively. They may contain things that are not explicitly stated in the task itself.
5. You may come across inconsistencies — when something in the task conditions or example contradicts what you already know. If that happens — try to figure it out. If you can’t — write your question in the open questions and clarify it during the process. Don’t leave open questions unresolved.
6. If a task seems unclear or unachievable — it only seems so. Try breaking it down. Most parts will likely become understandable.
7. On your way you’ll encounter various types of tasks. Bonus tasks are for the most meticulous and curious. These tasks are more difficult and optional, but if you complete them, you will gain extra experience and knowledge.
8. Don’t try to trick the system or those around you. You’ll only be deceiving yourself.
9. Got a question? Ask your neighbor to the right. If that doesn’t help — ask the one to the left.
10. When using help — always make sure you understand the why, how, and what for. Otherwise, the help won’t be meaningful.
11. Always push only to the develop branch! The master branch will be ignored. Work in the src directory.

## Chapter II
### General Information

**A Brief History of the Language**

Go (its actual name, though the term “Golang” is used for easier searching to avoid conflicts with the English verb “go”) is a compiled, multithreaded programming language developed by Google.

The first version of Go was released in November 2009. It was designed as a simpler alternative to C++, as there were a number of problems within the company that were hard to solve quickly. Go was developed to meet modern requirements for a language — namely, native support for multi-core systems and the extended capabilities of the standard library.

You can read more [here](https://go.dev/doc/).  
The source code of the language is available [here](https://github.com/golang/go).

**Installing the Compiler**

Installation files for various operating systems, including the source code with build instructions, can be found in the **Downloads** section on the [official Go website](https://go.dev/dl/).  
We recommend installing the latest stable version of the language.

**Toolchain**

The compiler comes with a set of programs and utilities that can be accessed using the go command. These are referred to as the **Go Toolchain**.

For example, the go env command allows you to check your Go installation and at the same time displays the list of environment variables used by the toolchain, along with their values.

The go fmt and go test commands are also part of the built-in toolkit. To compile source code into an executable file, use the go build command. For local execution, the go run command is more commonly used — it combines compilation and execution.

To display help for a specific command, use: `go help <command>`.

You can read more about the Go Toolchain [here](https://pkg.go.dev/cmd/go).

Before version 1.11, when modules were introduced, all source code files had to be located in $GOPATH/src.  
With modules, Go project source code can reside in any directory.  
Since version 1.16, modules are enabled by default. Use modules in your projects — to do this, simply run: `go mod init` in the package directory.

**Code Editors**

We suggest installing an IDE or connecting a plugin to your preferred editor:

If you want an IDE — install **JetBrains GoLand**.  
However, the free trial version is only available for one month.

A good choice is also to install the **Go plugin** for the **Visual Studio Code** editor.

**Where to Start**

We recommend beginning by reading the article [Get Started](https://go.dev/doc/tutorial/getting-started).

The [Go Tour](https://go.dev/tour/welcome/1) is an excellent starting point for learning the language.  
All three links are static, as this is the main Go learning resource.

**Articles for Studying Golang**

**Recommended Reading List**

- **Awesome-Go** — you can find it by searching “awesome-go github”.
- **Go Proverbs** — search for the article using keywords like “Go proverbs” (in Russian), original version: “go proverbs”.

## Chapter III
### Project «Smart Utilities»

### Task 1. Console Calculator

You need to implement a console calculator:

1. The user enters the left operand and presses Enter.
2. The user enters the operation and presses Enter.
3. The user enters the right operand and presses Enter.
4. The program performs the selected operation and outputs the result to the console.
5. The calculator must support addition, subtraction, multiplication, and division.
6. Division accuracy — 3 decimal places.
7. Before entering each operand or operation, the program should display a prompt to the user.
8. Operands and operation are entered via the console by pressing Enter after each input.
9. Operands must be of type Float64.
10. If invalid input is provided, the program should print an error message (Invalid input) to stdout and prompt the user to enter the data again.

**Functional requirements:**

- Example:

  Input left operand: 
  > `10`

  Input operation 
  > `+`

  Input right operand
  > `15`

Only the standard library may be used.

_Hint:_ Parameters are read from the console as strings and must be converted to the required type.

### Task 2. Most Frequent Words

You need to implement a program that receives a list of words and a number K as input. The output must be a list of the K most frequently occurring words, sorted in descending order of frequency.

1. The program reads a string of words.
2. The program reads the number K, which is the limit on the number of output words.
3. The program determines how many times each word appears in the input string.
4. The program sorts the resulting array in descending order and returns the first K words.
5. The delimiter between words is the space character.
6. If the number of unique words is less than K, the program returns a list of all unique words, sorted by frequency.
7. If an empty string is given as input, the output must also be an empty string.
8. A word is any sequence of characters separated by spaces.
9. If multiple words have the same frequency, they must be sorted lexicographically.
10. You must implement tests that cover the following cases:  
    a. Normal behavior when K is less than the number of unique words.  
    b. Input is an empty list of words.  
    c. Input list contains fewer unique words than K.

**Functional requirements:**

- Example:

Input:

```
aa bb cc aa cc cc cc aa ab ac bb
3
```

Output:

```
cc aa bb
```

Because:  
`cc` appears 4 times,  
`aa` — 3 times,  
`bb` — 2 times.

Only the standard library is allowed.  
We recommend using the testing package for writing your tests.

### Task 3. Slice Intersection

1. The program reads the first list of numbers from the console, separated by spaces; input is finalized by pressing Enter.
2. The program reads the second list of numbers from the console, separated by spaces; input is finalized by pressing Enter.
3. The input lists are unordered.
4. The program works only with numbers of type int.
5. The program finds the intersection of the two lists and returns it in the order the values appear in the first list.

**Input:**

```
5 3 4 2 1 6
6 4 2 4
```

**Output:**

```
4 2 6
```

1. If the intersection is empty, the program outputs: Empty intersection.
2. If the list contains values that are not of type int, the program outputs the error message Invalid input and terminates.

### Task 4. Visit Log

1. The program implements 3 commands in interactive mode (commands are entered in the console without stopping the main execution).
2. The **Save** operation allows saving the visitor’s full name, the doctor's specialization they visited, and the visit date into the log.  
    a. The doctor’s specialization is provided as a string.  

    b. The visit date is in the format YYYY-MM-DD. 

    c. `Example: Save \n Ivanov Ivan Ivanovich \n orthopedist \n 2024-04-13`

    d. `\n` means a line break, i.e. Enter.

1. The GetHistory operation allows viewing the patient's visit history. It takes the patient's full name and returns a list of specialization–date pairs.  
    a. Example: `GetHistory \n Ivanov Ivan Ivanovich`

    b. Output: `orthopedist 2024-04-13 \n neurologist 2024-05-24 \n`

    c. `\n` means a line break, i.e. Enter.

1. The GetLastVisit operation returns the last visit to a specific specialist. It takes the patient’s full name and the doctor’s specialization and returns the last visit date.  
    a. Example: `GetLastVisit \n Ivanov Ivan Ivanovich \n orthopedist`

    b. Output: `2024-04-13`
    
    c. `\n` means a line break, i.e. Enter.

1. If the patient is not found, the program returns an error of type `PatientNotFoundError` with the message `patient not found` (this must be implemented manually).

Data is stored in memory only; long-term data persistence is not required.

**Hint:** you may use the map type as a storage structure.  
To store visit data, define an additional structure with the fields: specialization and visit date.