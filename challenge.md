# Backend Engineer - Coding Challenge

This coding exercise is an opportunity for you to show us how you break down requirements into actual code as well as to demonstrate quality, clean coding style, and creativity in problem-solving. The task is designed to represent an example of the kind of work Bitly engineers do. You should not expect there to be any gotchas.

For the purposes of this challenge, we will be working with CSV and JSON files rather than databases/data streams/APIs; however, it does reflect data typical of what we work with.

We are providing two data files to be used with the exercise which can be found here: https://bit.ly/BitlyBackendCodingChallengeFiles 

The files are as follows:
* `encodes.csv` - This file contains information about shortened links or "encodes". It is a mapping of long URL to short domain and bitlink hash.
* `decodes.json` - This file contains raw data representing clicks on bitlinks in JSON format.

Use this dataset to answer the questions laid out in the Problem Statement below.

## Problem Statement

**Problem:** Using the "click event data in decodes.json, calculate the number of clicks from 2021 for each record in the encodes.csv dataset.

Your solution should include a runnable program.

## Acceptance Criteria

* When run, your program should output the following to the console:
    ```
    A sorted array of JSON objects containing the long URL as the key and the click count as the value. 
    The array should be sorted descending by click count.
    
    The form should be as follows: [{"LONG_URL": count}, {"LONG_URL": count}]
    
    Example: [{"https://google.com": 3}, {"https://www.twitter.com" : 2}]
    ```
* Count clicks only if they occurred in the year 2021
* Solution should log relevant actions or unexpected conditions
* Solution should include unit tests for each function 
* Solution should be well documented 
* Solution should include a README covering:
  - A list of dependencies of your project, as well as how to install them (we may not be experts in your chosen language, framework, and tools)
  - Instructions for running your application/script (you may include a Dockerfile or a Makefile, but this is not a requirement)
  - A description of any design decisions 

## Language and Framework

Feel free to choose any language and framework you are comfortable with. The language that we primarily use at Bitly on the backend is Go, but you should not feel pressured to write in Go if you are not currently comfortable with it. We want you to be able to focus more on your solution than the tools and language.

If you advance to the next stage of interviews, a Live Coding interview session will involve making improvements to and building on your Coding Challenge submission. This is one reason we emphasize using a language that you are comfortable with.

## Bitly Glossary

* bitlink: A short link created by the Bitly platform
* hash: The system-generated backhalf of a bitlink
* encode: The act of shortening a long link into a bitlink
* decode: A redirect or click event 


### Good luck!