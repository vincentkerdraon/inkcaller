This is a sample ink file to use the ink caller library. 
Please check the official doc on how to write ink syntax.

This file is used in code examples.

The lib is designed for:
_ BeginStory() to init ink state (maybe with the random seed), and load the state
_ GoToKnot() to call any knot. For example knot="Hub" to find all the storylets.
_ GetResourceText() for simple text, like "Credits" in this file.
_ ContinueStory() when a choice was available and to move to next step.

Each call is independent, and we provide the ink state each time. Meaning multiple stories can be executed at the same time. 
It is possible to inject external data into the ink state, and read it from ink to make conditions. The variables must be declared initially in ink.


//Example of using ink internal variables.

LIST ScenesID =  SceneAvailable_2

VAR ScenesAvailable = ()

//Once this variable declared, it can be modified in the ink state between calls.
//For example, we could display Scene2 only if {level>1}
VAR Level = 0 

// DEBUG mode adds a few shortcuts - remember to set to false in release!
VAR DEBUG = true

* [Visit the Hub] ->Hub
* {DEBUG} [INK_DEBUG] -> INK_DEBUG

=== Hub

+ [Start Scene1] ->Scene1
+ {ScenesAvailable?SceneAvailable_2} [Start Scene2 (only visible after Scene1)] ->Scene2
+ {Level==1} [Start Scene3 (only visible by changing ink internal state)] ->Scene3
+ ->DONE

=== Scene1
Welcome to Scene1.
+ Go to Scene1_1
->Scene1_1

=== Scene1_1
Welcome to Scene1_1.
That's the end of this scene. (no choice available)
Now Scene2 is available from the hub!

Note you can't go back to the hub from here if you are using inky editor or the web export. To mitigate that, you could for example create a "INK_DEBUG" knot available in all the dead-ends to test outside the lib. And then parse and remove this choice when in production.
 ~ ScenesAvailable += (SceneAvailable_2)
 
->DONE

=== Scene2
Welcome to Scene2.
You will need to visit Scene1 to see Scene2 again.
 ~ ScenesAvailable -= (SceneAvailable_2)
->DONE

=== Scene3
Welcome to Scene3.
You can only access it from the hub by changing the variable Level.
(So that when you need to input data into ink)
Remember that "INK_DEBUG"? Maybe you also want to arbitrary change the level in this knot.

And if you want to change something in your external data model, just write it and parse the text. For example like this:
>>> level+=1 
(The syntax is up to you, as long as it does not conflict with ink.)
->DONE

=== Credits
Maelle & Vincent
->DONE


=== INK_DEBUG
IN DEBUG MODE!

+ [Hub] ->Hub
+ [Hub NO DEBUG] 
    ~DEBUG=false
    ->Hub
+ [Credits] ->Credits
+ [Hub with Scene2] 
    ~ ScenesAvailable += (SceneAvailable_2) 
    ->Hub
+ [Hub with Scene3] 
    ~ Level = 1
    ->Hub

- ->DONE
