# covidb

This problem statement has been inspired by the applications - https://www.covid19india.org/ and https://cluster.covid19india.org/ 


The problem statement is only trying to understand if the student can solve the problem with a good design, so does not bring in the remorse part of COVID which is the stats on deaths


On the whole this looks like an easy problem - not much complexity in terms of algorithm. Modelling too - it feels it’s a bit easy. Maybe it can be tweaked.


Currently if I think in a simple way - I can solve this probably with just a list data structure, though graphs can also be used? Considering there are relationships among people. That was the initial idea. But graphs can make things quite complex. Gotta check the performance of the different approaches, different data structures but I think we aren’t going to test the candidate on performance of the program much, so using any data structures might be okay

----

Problem Statement:
With the recent COVID situation, statistics and information regarding the spread is all that people are looking at in news, social media and everywhere else on the Internet. There are still people who are going out like it’s a usual time, and the government wants to show the seriousness of the situation, so they want you to build a system that will tell people some information regarding people who have been positively tested with COVID


You are going to build a prototype for it. For the sake of simplicity and for prototyping fast, you will build a command line tool (command line interface/CLI)


Build an application
* That will allow a person to add information about a positively tested person. Information includes - gender, age, area and city of the person. For privacy purposes, the person's name must not be included in the information, as this information is going to be open to the public. Generate a unique ID for every person to differentiate people. If the person has a relationship with another already positively tested person, the user should be able to capture it. Examples of a relationship are “friend of”, “father of”, “son of”. So, if the new person whose information is being stored is X, and Y is an already positively tested person, if Y is the father of X, you can put the information for X as “son of Y”. So, relationship information includes relation name and person’s unique ID, and in the example the values of these two are “son of” and Y while giving information about X and that’s how the user gives information to the application. A person can be related to more than one positively tested person.
* That will allow a person to view the information of all the people who have been tested positive as a list. Since the list can grow very large, some sort of pagination can be done so that the user sees only 10 people’s information at a time, and if the user wants to see more, they can move to the next page by giving input
* That will work even after restarts by persisting information


Extensions
* While listing the information of the people, the user can choose a person from the list to know complete information about them and their relationship with other positively tested people. 
* Ability to update any information about people who are already present in our database
* Ability to type an area name and then see how many people are affected in that area

---

The name of project `covidb` has been derived by combining the words `COVID` and
`DB` (Database), since we are building a database of positively tested people.

---

Honestly, after trying out the program, it feels a bit boring. Mostly getting
and showing input. Not much to think about actually. Hmm. 
