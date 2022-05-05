# user feed algorithm


## Engagement algorithm

1. categorizing questions

based on the words used on a question we are categorizing the question. each post gonna have 11 relatable words


2. user upvote post

when user upvote post that means he like this type of questions / interested on topics like this so the keywords of 
the question will append in to the user  interested on topics


so for each users there will be topics that he's interested on

---

finding posts based on these intereted keywords

1. we can fetch then from the database without a problem. well by doing some sql magic

2. sorting these values of user to find the most interested ones

as a example first we get all the interests of the user, imagine if the following list is the
interests of user

```json
[the of lorem ipsum and it type with dummy text the lorem ipsum of and it a dummy has type]
```

then we get the count of duplicates and sort them like below

```go
[{of 2} {and 2} {type 2} {it 2} {dummy 2} {the 2} {lorem 2} {ipsum 2} {with 1} {text 1} {a 1} {has 1}]
```
now we can find the most interested text of different users, like as this example this user is mostly interested on 
`of` and `and`

since these words are common words of english we can ignore them. first we need to create a normal words of 
english. in fact i found a array with these common words


```java
{"the","of","and","a","to","in","is","you","that","it","he","was","for","on","are","as","with","his",
"they","I","at","be","this","have","from","or","one","had","by","word","but","not","what","all","were",
"we","when","your","can","said","there","use","an","each","which","she","do","how","their","if","will",
"up","other","about","out","many","then","them","these","so","some","her","would","make","like","him",
"into","time","has","look","two","more","write","go","see","number","no","way","could","people","my",
"than","first","water","been","call","who","oil","its","now","find","long","down","day","did","get",
"come","made","may","part"};

```

probably words like this we can ignore since they are useless.


---

based on remaining values we are fetching the posts
