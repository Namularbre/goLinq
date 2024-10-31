# Go linq

Linq is an acronym for Language Integrated Query. It's originally a technology made by Microsoft, 
you can find it in C#: https://learn.microsoft.com/en-us/dotnet/csharp/linq/.

It's a very practical feature I want to bring in my go project, so I want to reimplement a lib the contains a part of it.

## Features
<li>Select</li>
<li>Where</li>
<li>Limit</li>
<li>Skip</li>

All the methods works with slices, but I will work on a struct that works like a query, to make go linq easier to write and read.

## Install

````
go get github.com/Namularbre/goLinq
````

## Author
[Namularbre](https://github.com/Namularbre)
