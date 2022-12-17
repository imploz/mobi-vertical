# Go MOBI

[![Github Actions CI](https://github.com/leotaku/mobi/workflows/check/badge.svg)](https://github.com/leotaku/mobi/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/leotaku/mobi)](https://goreportcard.com/report/github.com/leotaku/mobi)
[![Go Reference](https://pkg.go.dev/badge/github.com/leotaku/mobi.svg)](https://pkg.go.dev/github.com/leotaku/mobi)

This package implements facilities to create KF8-formatted MOBI and AZW3 books.
We also export the raw PalmDB writer and various PalmDoc, MOBI and KF8 components as subpackages, which can be used to implement other formats that build on these standards.

## This fork

This fork was created from [the original](https://github.com/leotaku/mobi) to add support for **vertical** right-to-left text
(and also vertical left-to-right text for completeness).
(in addition to support for **horizontal** RTL which seems to be in the code, though I haven't tested it and didn't modify it's behavior in this fork).
This allows this library to generate japanese azw3 ebooks with the same text orientation as in typical Japanese printed novels to enjoy them on a Kindle reader.

Also, I added `ExtraHead` field to `Book` to allow adding e.g. extra styles directly to the generated html files.

### Usage

In order to get vertical RTL text, use mobi.Book's `RightToLeft=true` and `Vertical=true`:

```go
  mb := mobi.Book{
    // ...
    RightToLeft: true,
    Vertical:    true,
  }
```

In order to get emphasis markers from [yomukaku](https://yomukaku.jp) work use the following for Book's  

```go
  mb := mobi.Book{
    // ...
    ExtraHead: `
      <style> 
        .emphasisDots {
          font-weight: normal;
          text-emphasis-style: dot;
          -webkit-text-emphasis-style: dot;
          }
      </style>`
  }
```

### Possible improvements of this fork

It would probably be cleaner to introduce the necessary styles as CSS files and refer to them in each html file using Book.CSSFlows instead of this dirty yet easy hack in `templating.go`.
The same goes for injecting e.g. the style for `emphasisDots` class via `ExtraHead`.

## Known issues

+ Chapters are supported but subchapters are not
+ Old readers without KF8 are not supported (Kindle 1, 2 and DX)
+ Books without any text content are always malformed
+ Errors during template expansion result in a panic

## References

+ MobileRead Wiki
  + [MOBI format](https://wiki.mobileread.com/wiki/MOBI)
  + [PDB format](https://wiki.mobileread.com/wiki/PDB)
  + [KF8 format](https://wiki.mobileread.com/wiki/KF8)
+ Calibre source code for [MOBI support](https://github.com/kovidgoyal/calibre/tree/master/src/calibre/ebooks/mobi)
+ Vladimir Konovalov's [Golang MOBI writer](https://github.com/766b/mobi)
+ Library of Congress on the [MOBI format](https://www.loc.gov/preservation/digital/formats/fdd/fdd000472.shtml)

Huge thanks to the authors of these resources.

## License

[MIT](./LICENSE) © Leo Gaskin 2020-2022
