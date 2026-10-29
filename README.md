# url-slug-generator

Converts arbitrary UTF-8 strings into URL-safe ASCII slugs with a configurable separator.

```go
import url_slug_generator "url-slug-generator"

slug := url_slug_generator.Generate("Hello World")
// slug == "hello-world"

sl := url_slug_generator.New("_")
slug = sl.Generate("Go 1.21 Release")
// slug == "go_1_21_release"
```

This exists because most slug libraries pull in large transliteration tables to convert "é" into "e". That is useful when it works, but it requires a dependency and the results are inconsistent across locales. This library takes the simpler route: keep `[a-z0-9]`, lowercase ASCII, drop everything else, and join the survivors with a separator. Zero dependencies, deterministic output, and the code fits in one file.

The edge you will hit: non-ASCII characters are discarded, not romanized. "Café" becomes "caf", not "cafe". If you need transliteration, this is not the right library. The trade-off is deliberate so the output never surprises you and the build never needs a network.

Exported API:

- `New(separator string) *Slugger` — construct a slugger with a custom separator; empty defaults to `"-"`.
- `(*Slugger).Generate(s string) string` — produce the slug.
- `Generate(s string) string` — package-level helper using the default `-` separator.
- `Slugger` struct with field `Separator string`.
