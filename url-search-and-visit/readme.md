# URL Search and Visit

This problem has been inspired by the URL bars in Internet Browsers, like
Firefox, Chrome.

When a user types something in a browser's URL bar, they see a drop with
multiple results for URLs based on their history and the first result is also
auto completed in the URL bar, based on this concept, the problem has been
formed.

This is kind of a general search problem and can have variations, for example -
search engine search problem, searching in a blog for articles, searching
in an ecommerce site for products, searching for a contact in a digital phone
book etc.

Problem:
There is a finite list of website URLs.

A user can type an input, which is a string of characters - alphabets only.
Other characters like `-`, `.`, numbers etc can be part of extensions.

With this input, the user can search for website(s) based on site URL.

Search logic is - if a URL contains the search string, then the URL is part
of the search result. Case insensitive search, so `A` and `a` are the same.

User types the search string and presses enter to see the set of results.

User then chooses the site, to tell they are visiting the site.

Next time, when the search results contain the site that the user visited, the
site has higher rank and hence comes to the top in the search result.

Every time a user visits a site, the particular site's rank increases.

The search results are always sorted by the rank of each site.

Example Input and Output

```
$ ./url-search
Menu
1. Search a website
2. Exit
Choice: 1

Type search string: net

Search Results:
1. https://netflix.com
2. https://app.diagrams.net
3. https://alternativeio.net

Menu
1. Visit a website from search result
2. Search a website
3. Exit
Choice: 2

Type search string: flix

Search Results:
1. https://flix.com
2. https://netflix.com
3. https://alternativeio.net

Menu
1. Visit a website from search result
2. Search a website
3. Exit
Choice: 1

Which website do you want to visit? 2
You have visited https://netflix.com !

Menu
1. Visit a website from search result
2. Search a website
3. Exit
Choice: 2

Type search string: flix

Search Results:
1. https://netflix.com
2. https://flix.com
3. https://www.flixo.com/

Menu
1. Visit a website from search result
2. Search a website
3. Exit
Choice: 3

Exited!
```

In the above example, you can see how `https://netflix.com` comes on the top
once the user has visited the website once, because it has higher ranking,
and other sites don't have any rank yet because the user hasn't visited any
of them.

One thing to note is, in this case there's already a list of websites in the
program, hardcoded. Another possible extension is below

Extension: Allow user to type any URL and visit the site directly instead of
searching for sites. This can be an extra option to the user
`Visit website with URL` apart from the `Visit a website from search result`.
This new option can add to the list of websites the program knows and can be
used in the search to show in the search results. This way we have a dynamic
list of websites rather than a static list.

```
$ ./url-search
Menu
1. Visit a website with URL
2. Search a website
3. Exit
Choice: 1

Which website do you want to visit? https://keep.google.com
You have visited https://keep.google.com !

Menu
1. Visit a website with URL
2. Search a website
3. Exit
Choice: 3

Exited!
```

In the above examples you can also see how the `Visit a website from search result`
shows up only when a user has done the search at least once. Also, it remembers
the last search results and the user can visit any of the sites in the last search
result.

Extension: Implement the same for multiple users. The program asks the user to
login, and user can also logout. Each user goes to different sites different
number of times, so the rank of site depends on the user doing the search

Example Input and Output

```
$ ./url-search
Menu
1. Login
2. Exit
Choice: 1

Username: Karuppiah
Logged in as Karuppiah!

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 2

Type search string: flix

Search Results:
1. https://flix.com
2. https://netflix.com
3. https://alternativeio.net

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 1

Which website do you want to visit? 2
You have visited https://netflix.com !

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 2

Type search string: flix

Search Results:
1. https://netflix.com
2. https://flix.com
3. https://www.flixo.com/

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 3

Logged out!

Menu
1. Login
2. Exit
Choice: 1

Username: Bharath
Logged in as Bharath!

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 2

Type search string: flix

Search Results:
1. https://flix.com
2. https://netflix.com
3. https://alternativeio.net

Menu
1. Visit a website from search result
2. Search a website
3. Logout
4. Exit
Choice: 4

Exited!
```

As you can see, when `Karuppiah` logs in and searches after visiting
`netflix.com`, it shows him that website on the top in the search result,
but for the same search, for `Bharath`, it shows a different result as
with `netflix.com` _**not**_ coming at the top `Bharath` has not yet visited any
website

Extension: Apart from website URL, capture website title and search based on
title too. The search results are first sorted by rank and if two sites have
the same rank, then sort them by the following parameters:
* If the search string matches both website title and website URL, then it
has the highest priority and comes on the top
* If the search string matches website URL alone then it has higher priority
than if the search string matches website title alone

Example Input and Output

```
$ ./url-search
Menu
1. Search a website
2. Exit
Choice: 1

Type search string: recruitment

Search Results:
1. Recruitment.com | Talent Acquisition and Recruiting Ideas https://recruitment.com/
2. Recruitment - Google Drive https://drive.google.com/drive/u/0/folders/123454

Menu
1. Visit a website from search result
2. Search a website
3. Exit
Choice: 3

Exited!
```

In the above example, all sites have rank as 0, but the search string
`recruitment` matches the site URL `recruitment.com` and the site's title too
and hence has higer priority than the other site, which seems to be a google
drive docs link with the doc named `Recruitment` and hence the title for it
is that, but only the title matches the search string

Extension: For searching a string in the URL, do not search the TLD part (Top
Level Domain) of the domain. TLD examples - `.com`, `.net`, `.in`, `.io`,
`.works` etc. This is the last part of the domain name in a website URL.
If a search string matches only this part of the URL, for example `net` search
string matches `https://app.diagrams.net`, show it in the results only when
there is no other search result left out, based on other URL parts, and website
titles.

---

Modelling:
1. Site - URL, Title
2. User
3. Visit/Hit Count - for a given user visiting a site, what's the number of
hits / visits to the site by the user.

Problem Solving:
1. Filtering for searching
2. Sorting for relevance of search results
3. Context based sorting based on the user
