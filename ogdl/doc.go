// (C) Copyright 2012-2013, Rolf Veen.
// See the LICENCE file.

/*
Package ogdl is used to process OGDL, the Ordered Graph Data Language.

OGDL is a simple textual format to write trees or graphs of text, where indentation
and spaces define the structure. Here is an example:

  network
    ip 192.168.1.100
    gw 192.168.1.9
    
The languange is simple, either in its textual representation or its
number of productions (the specification rules), allowing for compact
implementations.

OGDL character streams are normally formed by Unicode characters, and 
encoded as UTF-8 strings, but any encoding that is ASCII transparent
is compatible with the specification and the implementations.  
    
See the full spec at http://ogdl.org.
*/
package ogdl