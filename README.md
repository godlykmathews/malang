# Malang Compiler  
**"They were teaching me compiler design, so a compiler is what I designed."**  
<p align="center">
<img src="./media/logo.png" height=300 width=300>
</p>

## What is Malang?
Malang is... well, let's just say it's a *unique* programming language. Born out of sheer defiance and an unwavering commitment to proving a point, Malang is my way of showing that I *did* in fact pay attention in my compiler design course—whether my professor appreciates the outcome or not.  

Will Malang change the world? No.  
Will it revolutionize software development? Absolutely not.  
Did it make me question my life choices? Almost certainly.  

It’s not just a language. It's a *concept*. A statement. A joke that spiraled completely out of control. But hey, **I** made a compiler!  

Malang is written _Manglish_, Malayalam written in English script—because why not? If you've ever wanted to write code that looks like a casual WhatsApp conversation between two Malayalis, Malang is here to enable that madness.

## Features  
- **Manglish to Go Compilation** - Because writing Malayalam in English letters is an art.  
- **Code Generation** - Magic! and go for the most part. 
- **Execution** - You write Manglish, Malang does the rest. Theoretically.  
- **Serious Yet Not Serious** - It works, but should it? Debatable.

## Installation
Since you clearly have excellent taste, getting started is easy:
```sh
git clone https://github.com/Rohith04MVK/malang
cd malang
go build -o malang
```
## Examples
Let's walk through some examples to see Malang in action. We'll start simple and gradually build up to more complex (well, *relatively* complex) code.

**1. Basic Input and Output:**

```go
parayu("Hello, ninte per entha?")  // Asks the user for their name.
kelk(name)                       // Reads the user's input and stores it in the variable 'name'.
```
- parayu("..."): This is your basic output statement. It prints the string within the double quotes to the console.

- kelk(variable): This statement takes input from the user and stores it in the specified variable.

**2. Conditional Statements (If-Else):**
```go
ith_sheriyano (name == "Rohith") enkil {
    parayu("Eda, ithu ninte thante language alle!") // Special message for Rohith (That's me :D).
} alle {
    parayu("Nannayittanu! Sugamano, " + name + "?") // Greeting for everyone else.
}
```
- **ith_sheriyano (condition) enkil { ... } alle { ... }**: This is your classic if-else statement.
    - ith_sheriyano: Kind of like "if it is true".
    - enkil: Think of it as "then".
    - alle: This means "else".
    - The code inside the enkil block executes if the condition is true. Otherwise, the code inside the alle block executes.

**3. While Loops:**
```go
ennam = 0                         // Initializes a variable 'ennam' to 0.
ellam_sheriyano (ennam < 5) enkil {  // Loops as long as 'ennam' is less than 5.
    parayu("Count: " + ennam)      // Prints the current value of 'ennam'.
    ennam = ennam + 1             // Increments 'ennam'.
}
```
- **ellam_sheriyano (condition) enkil { ... }**: This is your while loop.
    - ellam_sheriyano: Roughly translates to "while everything is okay".
    - The code within the curly braces {} will repeatedly execute as long as the condition (in this case, ennam < 5) remains true.

**4. For Loops:**
```go
oron_ayi i edukk (1..5) {       // Loops from 1 to 5 (inclusive).
    parayu("Value: " + i)     // Prints the current loop variable 'i'.
}
```

- **oron_ayi i edukk (start..end) { ... }**: This is your for loop.
    - oron_ayi: Means something like "one by one".

    - i: The loop variable (you can choose a different name, of course).

    - edukk: Think of it like "take".

    - (1..5): Specifies the range of values for the loop variable. It starts at 1 and goes up to and including 5.

    - The code inside the curly braces {} is executed for each value of i in the range.

More examples can be found in the `/examples` folder :)
## Why Malang?
Because learning is best when it's fun, and nothing says *"I understand compiler design"* quite like creating a language nobody needed.

## Contributing
If you understand this project on a deep level, I have two things to say:
1. I'm impressed.
2. I'm impressed.

But hey, PRs are welcome.

## License
MIT—because even legally, I refuse to take this too seriously. You're free to use, modify, and distribute Malang, but if you actually deploy it in production, that's on you.
