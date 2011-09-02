unifun: A Silly Unicode Text Transformer
========================================

Sometimes it is fun to abuse Unicode. For example, the phrase "The Quick Brown
Fox jumped over the Lazy Dog" can be rendered using only Unicode codepoints as:

> 𝔗𝔥𝔢 𝔔𝔲𝔦𝔠𝔨 𝔅𝔯𝔬𝔴𝔫 𝔉𝔬𝔵 𝔧𝔲𝔪𝔭𝔢𝔡 𝔬𝔳𝔢𝔯 𝔱𝔥𝔢 𝔏𝔞𝔷𝔶 𝔇𝔬𝔤 (Blackletter)
> 𝕿𝖍𝖊 𝕼𝖚𝖎𝖈𝖐 𝕭𝖗𝖔𝖜𝖓 𝕱𝖔𝖝 𝖏𝖚𝖒𝖕𝖊𝖉 𝖔𝖛𝖊𝖗 𝖙𝖍𝖊 𝕷𝖆𝖟𝖞 𝕯𝖔𝖌 (Blackletter bold)
> ɓo◖ ʎzɐ˥ ǝɥʇ ɹǝʌo pǝdɯnɾ xoℲ uʍoɹ𐐒 ʞɔınΌ ǝɥ⊥ (Fipped)

The `unifun` program allows you to convert text from vanilla Latin characters
to these exotic Unicode beasts.

Compiling
---------

You need to have the Go SDK installed. (See the [getting
started](http://golang.org/doc/install.html) section of the Go website.) In
addition, you'll need to have the `gomake` progam in your PATH.

If the SDK is installed, you should be able to compile and install the utility
with one command:

    $ cd /home/foo/bar/unifun # This is the source directory for unifun
    $ ./all.bash install

Usage
-----

    usage: unifun [options] text
      -f="blackletter": Specify which 'font' to use. Use -l to see all.
      -l=false: List fonts and exit.
      -n=false: Don't print trailing newline.

Implementation
--------------

The command is implemented with the help of a package written in
[Go](http://golang.org).

    PACKAGE

    package uctricks
    import "."

    A selection of fun unicode 'tricks' to apply to text


    FUNCTIONS

    func FontList() []string
    Return a slice containing all font names.


    TYPES

    type Font struct {
	// contains filtered or unexported fields
    }
    A 'font' is a mapping from rune to rune and (optionally) a function to apply to the text before mapping.

    func FontNamed(name string) (*Font, os.Error)
    Return a pointer to the font named.

    func (f *Font) Apply(s string) string

License and Copyright
---------------------

This code is Copyright 2011 Rich Wareham and is licensed under the 2-clause
BSD-license. Full information can be found in the COPYING file.
