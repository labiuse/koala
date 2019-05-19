## Shell/CLI

1. a
- (for i in {1..100}; do echo $i; done; ) : A loop over numbers in range of 1 till 100 which are produced with built-in bash sequence expression and are echoed to stdout one per line. - grep 3: catches numbers that contain digit 3
- grep -v 1 : only select lines that doesn't contain digit 1
- paste -s -d+ - : removes new lines characters (-s) and concatenate them by '+' sign, '-' at the end means use stdin instead of file for input.
- bc: bc is calculator command

Print the nubmers 1 till 100 on separate lines, select the ones that contains 3, filter out the ones that contain 1, sum the remaining numbers.
Pipe ('|') operator connects stdout of first command to stdin of second command.

1. b

 [ ! -f /var/lock/myscript.lock ] && touch /var/lock/myscript.lock && (yum -y update >> /var/log/mylog.log 2>&1; ) && rm -f /var/lock/myscript.lock

if `/var/lock/myscript.lock` file doesn't exists create it, then updates the installed packages with yum package manager (-y suppress the questions from command with replying yes to all of them) and appends stdout and stderr of yum update command into `/var/log/mylog.log` file and at the end removes the lock file. Basically it tries to lock the `yum update` and keep the output in a log file.

But it has some flaws, `touch` command is not atomic so there can be race condition and one of the processes will create the file while the other one still would not fail (touch won't fail if file exists, it just updates the file's access and modification time). A better solution can be to use `mkdir` command which is atomic and also use `trap` (only works inside a bash script) to make sure the lock is removed even if `yum` is exited unexpectedly or with non 0 exit code:

> #!/bin/bash
> 
> if mkdir -p /var/lock/myscript; then
> 
>   trap 'rm -R /var/lock/myscript' 0;
> 
> yum -y update >> /var/log/mylog.log 2>&1;
>
>  fi;


2.a

(\+61|0)([- ()]*[2378]){1}([- ()]*[0-9]){8}

2. b
- +61312345678 
- +61212345678
- +612-12345678
- +61-2-12345678
- +61-2(1234)(5678)
- 02-12345678
- 02-1234(5678)

2. c

`$ grep -iEo '(\+61|0)([- ()]*[2378]){1}([- ()]*[0-9]){8}' /var/www/site1/uploads/phnumbers/ | sed 's/[-() ]//g' | sed 's/+61/0/g' | sort | uniq > numbers.txt`

Find the numbers, remove extra characters "-() ", replace +61 with 0 so all numbers would be in a unified format, and finally use `sort` and `uniq` commands to only list uniqe numbers.


## Software development

1. 

`// IsFirstUpper returns true if first character is uppercase`

`func IsFirstUpper(s string) bool {`

`  return len(s) != 0 && s[0] >= 65 && s[0] <= 90`

`}`

2. 
a) PHP

b) Variable $d

c) In the main scope variable $a, and in the scope of anonymous function variables $b and $c.

d) $d = [ 'age' => 2, 'experience' => 4, 'interview-1' => 4, 'interview-2' => 6 ];

In this example array $d represents the weight of each criteria in a hiring process, its keys are criteria and the values are weight. The callback shows each criteria with its weight in more human friendly format of `[criteria]x[weight]`.

e) `agex2,experiencex4,interview1x4,interview2x6`

f) `array_map` applies provided callback function on its input array parameters which in this case are the keys of array $d and values of array $d respectively. The callback function is removing '-','_',',' characters from keys of array d and concatenate them with character `x` and array $d values. Implode concatenate items of resulting array with ',' character as string. Array $d is not modified in this process.

3. Write a function in Go which returns the top two most frequent numbers from a list, in order of frequency first

For this question I've used priority queue implemented with heap, leveraging built-in heap package, the performance is O(logn).

4. Describe the main tenets of a microservice.


- Isolation and Service boundary: each service should have well defined boundary and single responsibility so their domains should not collide and their internal implementation should be abstracted in form of API contract.
- Services should not share data storage (Database, ...) and should communicate through API (Restful, GRPC, ...)
- Services should be able to be tested, deployed and scaled separately hence each one should have its own CI/CD pipeline.
- Each service should be monitored and have alerting in place.


5. How would you decide between using a monolithic vs micro-services architecture?

Things I would consider are:
- The type of project. For example is it a POC project or is it a application contributing to business goals. How big is it in terms of business domains? Can those domains become decoupled and separated? Microservices is more suited for a system that can be broken down to multiple sub systems with decoupled functionalities and separated concerns but those sub systems should still have enough logics or specific technical needs to be justified to be considered a different sub system.
- Size of the team. Usually micro-services is for a big scope application where there's many sub systems with different business domains that have many logics. Implementing such big and complicated system requires more than few teams and more than a handful of developers. 
- Team experience. If there's no one with experience in the team that can at least lay down the architecture and help with ground work then it is better to go with monolithics where challenges are more familiar for the team and it's not a unknown land for them.

At the end it's important to have all the aspects in mind and one can't rule out the others. For example if we have enough developers and they are all experienced with micro-services but the project is a simple POC that have one or two simple domains then we better go with monolithic. 

6. Describe your favorite design pattern.

Each design pattern is created to address specific concerns and problems and be a general solution to a common problem. They help with problems and decisions regarding to decoupling, separation of concerns, abstraction and etc. So in response to each particular problem there's usually only one and in some rare cases more that one design pattern that we can refer to. The patterns I've mostly used are Singleton, Object pool, Factory and Facade.

7. How would you modify an API without breaking backwards compatibility?
 
 From the moment an API is published we can't easily change its public contract hence removing or changing API parameters or expected response/behavior is usually out of picture. Instead of changing/removing parameters we can introduce new ones without making any change to old behavior. The presence of this new parameter(s) can indicated a new use case with different processing and business logic which also can have a new different response.
 
8. What's the difference between streaming and queuing?

They are different concepts and not easily comparable. 

Streaming refers to a stream of data with unknown size. Processing a stream needs special care in regards and care to connection since we, as processor of the data, don't know the size of the data and should process the data as it's coming.

Queuing: can refer to queue data structure (FIFO) or message queuing. 
Message queuing is a system that can receive data in form of a predefined message and keep them till the messages are delivered to systems knowns as consumers which will receive the messages and will process them. Message queuing have some features like keep message in disk, re-queuing, queue priority and ...

