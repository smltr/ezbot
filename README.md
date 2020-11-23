# EzBot
EzBot is a web driver bot made to be easy to use and quick to deploy without any programming experience.

EzBot makes use of the Rod chromedriver package. I made it to give people without programming knowledge access to create a custom web bot.

Open up bot\instructions.txt and you will see a plain text file that gives the bot instructions. There are currently 3 types of instructions.
"navigate"
"click"
"input"
The example instructions.txt file makes use of all 3 instruction types.

The instructions will be executed in the order they are listed in instructions.txt.

Following the example instructions.txt file, you will see the first instruction "navigate". On the next line is an indention followed by the URL "
Every navigate instruction will have one field needed, which is the URL. All instructions have at least one field, some have 2. Fields are always indented for readability.
This first instruction tells the browser to navigate to https://wikipedia.org.

```
navigate
	https://wikipedia.org
```

The next example instruction is input. Input will always have 2 fields needed, a target and a value. The selector of an object can be easily obtained by opening up the chrome inspector, inspecting an element, right clicking the element in the inspector, and selecting copy > copy selector.
The second field is the value you would like to input. The example show would input the value "Futurama" into the object #searchInput.

```
input
	#searchInput
	Futurama
```

The last instruction in the example is a click. A click instruction has one field, which is the selector of the object you would like to input text into.

```
click
	#search-form > fieldset > button
```

This is still a work in progress. I am going to add a "wait" instruction which will be used to check the value of something repeatedly (example: if an item is in stock).
Please let me know of any bugs or requests.
