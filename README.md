# Spelling Bee Solver

I've been playing the NYT Spelling Bee with my partner lately, and thought it would be interesting to build a solver.
Since I've been interested in trying out Go for a while, this was my excuse.

## The Problem Statement

For those unaware, the basics of the NYT Spelling Bee is you have a honeycomb of letters; the centre letter, and six outer letters.
You can use any number of these letters in any combination to make a word, provided: 
1. Each word includes the centre letter
2. Words must be four letters or longer

Points are scored based on their length / letters used.
Unfortunately, valid words are a little arbitrary, with no set word list. However, as I understand it the bulk are sourced from Merriam-Webster, a JSON file of which I have included and use for constructing the word set.

## Solution Idea

In truth, a solution looping through dictionary words and checking if they can be made with the problem set, for example using a bitvector representing letters present, is a good solution (especially since the search domain is relatively small). However, I thought it would be interesting to leverage a data structure as a lookup table, to minimise the work done when given a spelling bee problem.
The approach is as follows:
1. Create a map to store words and a corresponding key.
2. Make a function to create a key from a word that only represents the information we care about matching, so in this case the unique letters used to generate a word (ordered so that they always map to the same bucket if they have the same letters).
3. When given a Spelling Bee problem, generate all unique combinations of letters, add the center letter to each, and look up the words under each permutation in the map (the union of these makes up the solution set).

It would have made sense to have done a similar principle with a custom hashmap implementation, using a hash function that is position and letter-count agnostic i.e: multiplying each of the unique Unicode values present together. However this would have involved implementing a lot of hashmap specific considerations such as resizing, collision handling, bucket distribution etc. that I thought it best to allow the default map implementation to handle (hence the 'bolt-on' approach taken).
The datastructure could also be extended to handle the finding of words made from letter subsets faster, by keeping a mapping of buckets with supersets of available letters to buckets with subsets (e.g: the bucket used to store words made up of 'abcd' would map to the buck of words made up from 'abc'), to really minimise the word done at query time. However given generating unique permutations of six letters is a trivial amount of work, I felt the way I handled this still kept in the spirit of the overall approach.

## How to Use

The dictionary package includes a Merriam-Webster JSON file and a JSON Parser helper function that pulls words from it. It looks for a "word" field, so isn't very format agnostic.

The solver package includes the actual code for the data structure, which is the solver struct. It has defined on it functions for taking a slice of strings to store a word list, and for solving a problem set given string for the centre letter and the outer letters.

A code example of how these are used: 
``` Go
	// Read the dictionary into a slice
	words, err := dictionary.ParseJSON("dictionary/dictionary.json")

	if err != nil {
		panic(err)
	}

	// Initialise a input the wordlist into a new solver
	s := solver.NewSolver()
	s.ParseWords(words)

	// Generate the solution set
	sol := s.Solve("i", "adehlp")
```
