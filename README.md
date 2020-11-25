# EzBot
EzBot is a web driver bot made to be easy to use and quick to deploy without any programming experience.
EzBot makes use of the [Rod chromedriver package](https://github.com/go-rod/rod). I made it to give people without programming knowledge access to create a custom web bot.

# Getting Started
### Installing
Go to 'Releases' and download ezbot.zip. If you are on desktop and on the main page for samtrouy/ezbot, the 'Releases' link should be on the right of the page. Extract the files to your preferred directory, just make sure they stay in the same folder always.

### .rod file
EzBot needs two files in the same directory to run. The first file is ".rod" and yes, there is supposed to be nothing before the period. This file includes configurations for the bot.
The default contents look like
```text
show
slow=0.5s
```
On the first line, 'show' will make the bot run in a viewable window. By removing this, EzBot will still run as normal but it will not be visible.
The second option, 'slow=0.5s' indicates the time in between each action that EzBot will sleep for. There needs to be time in between actions in order to not be flagged by a site as a bot.

It is a good idea to have 'show' in the .rod file while testing your bot so you can see what is happening. After you have run your bot and are confident it works, you may choose to remove the line 'show' from the file and let your program run in the background without a viewable window.

### instructions.txt
The instructions file, 'instructions.txt' will be the main file you will work with. This file will include all of the actions, or instructions, that your bot will carry out. Currently, there are 4 types of instructions the bot can perform: 'navigate', 'click', 'input', and 'wait'.

Let's start by writing our first bot. We will need three of the four types of instructions for this bot. Our bot will navigate to wikipedia.org and perform a search.
Our bot's first instruction will look like this:
```text
navigate
	https://www.wikipedia.org/
```
First, we define the type of action for the bot to take. This is 'navigate'. The action will be in lowercase letters.
Each action needs some type of value. Navigate only needs one value, which is the URL. To tell the bot which site to navigate to, we press enter to start a new line, press tab to indent, and type the full URL. The bot needs the full URL to work properly.
Values for actions are indented for readability. To recap, we started by typing the action we wanted the bot to take, entered a new line, entered a tab and then type the URL. All instructions will follow this same basic path, though different instructions may require different values, whereas 'navigate' only requires a URL.

So we said after the bot went to wikipedia we wanted it to perform a search. The next step will be to input some text into a field. Let's start by writing the next action we want the bot to take. Starting a new line, type 'input'.

```text
navigate
	https://www.wikipedia.org/
input
```
Great. Now we are going to need two values for every input function. Those values will be the 'target', where we want text to be input, and the 'text', which is the text we want to input. The target will always come before the value.

To get the target, we need to get the selector of the text field. To do this, we can use Chrome. Open up Chrome and go to wikipedia.org. Next, press ctrl+shift+i to open up the inspector. Now, right click the text input area you want to insert text in. The element should now be highlighted in the inspector, which will likely be on the right side of your browser. Right click the highlighted item in the inspector, click copy > copy selector. The selector is now copied to your clipboard and you can paste it into your instructions file. To correctly insert it into the file, we will start a new line after 'input', press tab and then paste. It will look as follows:


```text
navigate
	https://www.wikipedia.org/
input
	#searchInput
```

Now we need one more value for our input action, which as mentioned before is the text we want to input. This is simple, just type what you want to put into the search bar.


```text
navigate
	https://www.wikipedia.org/
input
	#searchInput
	Futurama
```

The last step is to click the search button. In your Chrome browser, while you still have the inspector opened, right click the search button and click inspect. There should now be a new element highlighted in the inspector. Right click that and click copy > copy selector. We can add the 'click' action to our file now. Click only needs one value, which is the target. For this we will put the selector of the item we want to click, which is what you just copied from the inspector.


```text
navigate
	https://www.wikipedia.org/
input
	#searchInput
	Futurama
click
	#search-form > fieldset > button
```

Great, now save your instructions.txt file and run bot.exe.
You should see your bot perform all of the actions that we originally set out for it to do.


There is one action that we didn't use in the example, and I will work on a tutorial that includes it as well. It is 'wait' and it takes one value, a selector, and it refreshes the page until it sees that element. An example use would be using it to check stock of an item. 

Quick example:
```
wait
	#ProductBuy > div > div:nth-child(2) > button
```
You could use this by copying the selector for the element that shows 'add to cart' for an item that is in stock, and use that selector with the wait action on the page of an item that is out of stock.


# Help

I am working on adding to the readme so please stand by. This project is also a work in progress and I am adding/fixing things every day. If you would like to ask for help, report a problem, or request a feature, please join the discord https://discord.gg/9zKuxcy8


