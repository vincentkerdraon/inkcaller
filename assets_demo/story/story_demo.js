﻿var storyContent ={"inkVersion":20,"root":[["^This is a sample ink file to use the ink caller library.","\n","^Please check the official doc on how to write ink syntax.","\n","^This file is used in code examples.","\n","^The lib is designed for:","\n","^_ BeginStory() to init ink state (maybe with the random seed), and load the state","\n","^_ GoToKnot() to call any knot. For example knot=\"Hub\" to find all the storylets.","\n","^_ GetResourceText() for simple text, like \"Credits\" in this file.","\n","^_ ContinueStory() when a choice was available and to move to next step.","\n","^Each call is independent, and we provide the ink state each time. Meaning multiple stories can be executed at the same time.","\n","^It is possible to inject external data into the ink state, and read it from ink to make conditions. The variables must be declared initially in ink.","\n","ev","str","^Visit the Hub","/str","/ev",{"*":"0.c-0","flg":20},"ev","str","^INK_DEBUG","/str",{"VAR?":"DEBUG"},"/ev",{"*":"0.c-1","flg":21},{"c-0":["^ ",{"->":"Hub"},"\n",{"->":"0.g-0"},{"#f":5}],"c-1":["^ ",{"->":"INK_DEBUG"},"\n",{"->":"0.g-0"},{"#f":5}],"g-0":["done",null]}],"done",{"Hub":[["ev","str","^Start Scene1","/str","/ev",{"*":".^.c-0","flg":20},"ev","str","^Start Scene2 (only visible after Scene1)","/str",{"VAR?":"ScenesAvailable"},{"VAR?":"SceneAvailable_2"},"?","/ev",{"*":".^.c-1","flg":21},"ev","str","^Start Scene3 (only visible by changing ink internal state)","/str",{"VAR?":"Level"},1,"==","/ev",{"*":".^.c-2","flg":21},{"*":".^.c-3","flg":24},{"c-0":["^ ",{"->":"Scene1"},"\n",{"#f":5}],"c-1":["^ ",{"->":"Scene2"},"\n",{"#f":5}],"c-2":["^ ",{"->":"Scene3"},"\n",{"#f":5}],"c-3":["done","\n",{"#f":5}]}],null],"Scene1":[["^Welcome to Scene1.","\n",["ev",{"^->":"Scene1.0.2.$r1"},{"temp=":"$r"},"str",{"->":".^.s"},[{"#n":"$r1"}],"/str","/ev",{"*":".^.^.c-0","flg":18},{"s":["^Go to Scene1_1",{"->":"$r","var":true},null]}],{"c-0":["ev",{"^->":"Scene1.0.c-0.$r2"},"/ev",{"temp=":"$r"},{"->":".^.^.2.s"},[{"#n":"$r2"}],"\n",{"->":"Scene1_1"},{"#f":5}]}],null],"Scene1_1":["^Welcome to Scene1_1.","\n","^That's the end of this scene. (no choice available)","\n","^Now Scene2 is available from the hub!","\n","^Note you can't go back to the hub from here if you are using inky editor or the web export. To mitigate that, you could for example create a \"INK_DEBUG\" knot available in all the dead-ends to test outside the lib. And then parse and remove this choice when in production.","\n","ev",{"VAR?":"ScenesAvailable"},{"list":{"ScenesID.SceneAvailable_2":1}},"+",{"VAR=":"ScenesAvailable","re":true},"/ev","done",null],"Scene2":["^Welcome to Scene2.","\n","^You will need to visit Scene1 to see Scene2 again.","\n","ev",{"VAR?":"ScenesAvailable"},{"list":{"ScenesID.SceneAvailable_2":1}},"-",{"VAR=":"ScenesAvailable","re":true},"/ev","done",null],"Scene3":["^Welcome to Scene3.","\n","^You can only access it from the hub by changing the variable Level.","\n","^(So that when you need to input data into ink)","\n","^Remember that \"INK_DEBUG\"? Maybe you also want to arbitrary change the level in this knot.","\n","^And if you want to change something in your external data model, just write it and parse the text. For example like this:","\n","^>>> level+=1","\n","^(The syntax is up to you, as long as it does not conflict with ink.)","\n","done",null],"Credits":["^Maelle & Vincent","\n","done",null],"INK_DEBUG":[["^IN DEBUG MODE!","\n","ev","str","^Hub","/str","/ev",{"*":".^.c-0","flg":20},"ev","str","^Hub NO DEBUG","/str","/ev",{"*":".^.c-1","flg":20},"ev","str","^Credits","/str","/ev",{"*":".^.c-2","flg":20},"ev","str","^Hub with Scene2","/str","/ev",{"*":".^.c-3","flg":20},"ev","str","^Hub with Scene3","/str","/ev",{"*":".^.c-4","flg":20},{"c-0":["^ ",{"->":"Hub"},"\n",{"#f":5}],"c-1":["^ ","\n","ev",false,"/ev",{"VAR=":"DEBUG","re":true},{"->":"Hub"},{"#f":5}],"c-2":["^ ",{"->":"Credits"},"\n",{"#f":5}],"c-3":["^ ","\n","ev",{"VAR?":"ScenesAvailable"},{"list":{"ScenesID.SceneAvailable_2":1}},"+",{"VAR=":"ScenesAvailable","re":true},"/ev",{"->":"Hub"},{"#f":5}],"c-4":["^ ","\n","ev",1,"/ev",{"VAR=":"Level","re":true},{"->":"Hub"},"done",{"#f":5}]}],null],"global decl":["ev",{"list":{},"origins":["ScenesID"]},{"VAR=":"ScenesID"},{"list":{}},{"VAR=":"ScenesAvailable"},0,{"VAR=":"Level"},true,{"VAR=":"DEBUG"},"/ev","end",null]}],"listDefs":{"ScenesID":{"SceneAvailable_2":1}}};
