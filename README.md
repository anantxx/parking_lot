# Parking Lot

Implementation of automated parking lot.

Language : GoLang
Version : go1.12.1

##### It is necessary that the code should be placed in $GOPATH/src/github.com/ folder.

## Problem Statement
 
I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is
given a number starting at 1 increasing with increasing distance from the entry point
in steps of one. I want to create an automated ticketing system that allows my
customers to use my parking lot without human intervention.
When a car enters my parking lot, I want to have a ticket issued to the driver. The
ticket issuing process includes us documenting the registration number (number
plate) and the colour of the car and allocating an available parking slot to the car
before actually handing over a ticket to the driver (we assume that our customers are
nice enough to always park in the slots allocated to them). The customer should be
allocated a parking slot which is nearest to the entry. At the exit the customer returns
the ticket which then marks the slot they were using as being available.
Due to government regulation, the system should provide me with the ability to findout:

- [x] Registration numbers of all cars of a particular colour.
- [x] Slot number in which a car with a given registration number is parked.
- [x] Slot numbers of all slots where a car of a particular colour is parked.

We interact with the system via a simple set of commands which produce a specific
output. Please take a look at the example below, which includes all the commands
you need to support - they're self explanatory. The system should allow input in two
ways. Just to clarify, the same codebase should support both modes of input - we
don't want two distinct submissions.
1) It should provide us with an interactive command prompt based shell where
commands can be typed in
2) It should accept a filename as a parameter at the command prompt and read the
commands from that file


## Architecture

Application is made up of following main Entities/ Layers :

##### model: has aggregate, entity and value object
##### repository: has repository interfaces of aggregate
##### service: has application services that depend on several models

## SetUp

```
$ ./bin/setup
```

## Run

```
$ ./bin/parking_lot <File_Name>
```


## Test

```
$ go test ./...
```
