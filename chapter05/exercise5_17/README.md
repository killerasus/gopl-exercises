# Exercise 5.17

Write a variadic function `ElementsByTagName` that, given an HTML node tree and zero or more names, returns all the elements that match one of those names. Here are two example calls:

<code>
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node<br/>
images := ElementsByTagName(doc, "img")<br/>
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")<br/>
</code>
