...


##  abbrev




Implementation: github.com/Masterminds/sprig/v3.abbrev

##  abbrevboth




Implementation: github.com/Masterminds/sprig/v3.abbrevboth

##  add
Sums numbers. Accepts two or more inputs. `add 1 2 3` will return `6`.



Implementation: github.com/Masterminds/sprig/v3.init.func15

##  add1
Increments by 1. `add1 3` will return `4`.



Implementation: github.com/Masterminds/sprig/v3.init.func14

##  add1f
Increments float number by 1.



Implementation: github.com/Masterminds/sprig/v3.init.func21

##  addf
Sums float numbers.



Implementation: github.com/Masterminds/sprig/v3.init.func22

##  adler32sum




Implementation: github.com/Masterminds/sprig/v3.adler32sum

##  ago
Returns duration from current time in seconds resolution.



Implementation: github.com/Masterminds/sprig/v3.dateAgo

##  all




Implementation: github.com/Masterminds/sprig/v3.all

##  any




Implementation: github.com/Masterminds/sprig/v3.any

##  append




Implementation: github.com/Masterminds/sprig/v3.push

##  atoi
Converts a string to an integer.



Implementation: github.com/Masterminds/sprig/v3.init.func12

##  b32dec
Decodes string from Base32 format.



Implementation: github.com/Masterminds/sprig/v3.base32decode

##  b32enc
Encodes string with Base32 format.



Implementation: github.com/Masterminds/sprig/v3.base32encode

##  b64dec
Decodes string from Base64 format.



Implementation: github.com/Masterminds/sprig/v3.base64decode

##  b64enc
Encodes string with Base64 format.



Implementation: github.com/Masterminds/sprig/v3.base64encode

##  base




Implementation: path.Base

##  bcrypt




Implementation: github.com/Masterminds/sprig/v3.bcrypt

##  biggest




Implementation: github.com/Masterminds/sprig/v3.max

##  buildCustomCert




Implementation: github.com/Masterminds/sprig/v3.buildCustomCertificate

##  camelcase




Implementation: github.com/huandu/xstrings.ToCamelCase

##  cat




Implementation: github.com/Masterminds/sprig/v3.cat

##  ceil
Returns greatest float value greater than or equal to input value. `ceil 123.001` will return `124.0`.



Implementation: github.com/Masterminds/sprig/v3.ceil

##  chunk




Implementation: github.com/Masterminds/sprig/v3.chunk

##  clean




Implementation: path.Clean

##  coalesce




Implementation: github.com/Masterminds/sprig/v3.coalesce

##  compact
Accepts a list and removes entries with empty values.



Implementation: github.com/Masterminds/sprig/v3.compact

##  concat
Concatenates arbitrary number of lists into one.



Implementation: github.com/Masterminds/sprig/v3.concat

##  contains
Tests if one string is contained inside of another. `contains "cat" "catch"` will return `true`.



Implementation: github.com/Masterminds/sprig/v3.init.func9

##  date
Formats date.



Implementation: github.com/Masterminds/sprig/v3.date

##  dateInZone
Same as `date` but with a timezone.



Implementation: github.com/Masterminds/sprig/v3.dateInZone

##  dateModify




Implementation: github.com/Masterminds/sprig/v3.dateModify

##  date_in_zone




Implementation: github.com/Masterminds/sprig/v3.dateInZone

##  date_modify




Implementation: github.com/Masterminds/sprig/v3.dateModify

##  decryptAES
Receives a Base64 string encoded by the AES-256 CBC algorithm and returns the decoded text.



Implementation: github.com/Masterminds/sprig/v3.decryptAES

##  deepCopy




Implementation: github.com/Masterminds/sprig/v3.deepCopy

##  deepEqual




Implementation: reflect.DeepEqual

##  default




Implementation: github.com/pluralsh/polly/template.dfault

##  derivePassword




Implementation: github.com/Masterminds/sprig/v3.derivePassword

##  dict




Implementation: github.com/Masterminds/sprig/v3.dict

##  dig




Implementation: github.com/Masterminds/sprig/v3.dig

##  dir




Implementation: path.Dir

##  div
Performs integer division.



Implementation: github.com/Masterminds/sprig/v3.init.func17

##  divf




Implementation: github.com/Masterminds/sprig/v3.init.func24

##  duration
Formats a given amount of seconds as a `time.Duration`.



Implementation: github.com/Masterminds/sprig/v3.duration

##  durationRound




Implementation: github.com/Masterminds/sprig/v3.durationRound

##  empty




Implementation: github.com/Masterminds/sprig/v3.empty

##  encryptAES
Encrypts text with AES-256 CBC and returns a Base64 encoded string.



Implementation: github.com/Masterminds/sprig/v3.encryptAES

##  env
Reads environment variable.



Implementation: os.Getenv

##  expandenv
Substitutes environment variable in a string.



Implementation: os.ExpandEnv

##  ext




Implementation: path.Ext

##  fail
Unconditionally returns an empty string and an error with the specified text. This is useful in scenarios where other conditionals have determined that template rendering should fail.



Implementation: github.com/Masterminds/sprig/v3.init.func26

##  first




Implementation: github.com/Masterminds/sprig/v3.first

##  float64
Converts to a `float64`.



Implementation: github.com/Masterminds/sprig/v3.toFloat64

##  floor
Returns the greatest float value greater than or equal to input value. `floor 123.9999` will return `123.0`.



Implementation: github.com/Masterminds/sprig/v3.floor

##  fromJson


Aliases: from_json

Implementation: github.com/Masterminds/sprig/v3.fromJson

##  genCA
Generates a new, self-signed x509 SSL Certificate Authority using 2048-bit RSA private key. It takes subject common name (CN) and cert validity duration in days as parameters. It returns object with PEM-encoded certificate and key. Note that the returned object can be passed to the `genSignedCert` function to sign a certificate using this CA.



Implementation: github.com/Masterminds/sprig/v3.generateCertificateAuthority

##  genCAWithKey
Generates a new, self-signed x509 SSL Certificate Authority using given private key. It takes subject common name (CN), cert validity duration in days and private key (PEM-encoded; DSA keys are not supported) as parameters. It returns object with PEM-encoded certificate and key. Note that the returned object can be passed to the `genSignedCert` function to sign a certificate using this CA.



Implementation: github.com/Masterminds/sprig/v3.generateCertificateAuthorityWithPEMKey

##  genPrivateKey




Implementation: github.com/Masterminds/sprig/v3.generatePrivateKey

##  genSelfSignedCert
Generates an SSL self-signed certificate.



Implementation: github.com/Masterminds/sprig/v3.generateSelfSignedCertificate

##  genSelfSignedCertWithKey




Implementation: github.com/Masterminds/sprig/v3.generateSelfSignedCertificateWithPEMKey

##  genSignedCert
Generates an SSL certificate and key based on a given CA.



Implementation: github.com/Masterminds/sprig/v3.generateSignedCertificate

##  genSignedCertWithKey




Implementation: github.com/Masterminds/sprig/v3.generateSignedCertificateWithPEMKey

##  get




Implementation: github.com/Masterminds/sprig/v3.get

##  getHostByName




Implementation: github.com/Masterminds/sprig/v3.getHostByName

##  has
Checks if a list has a particular element.



Implementation: github.com/Masterminds/sprig/v3.has

##  hasKey




Implementation: github.com/Masterminds/sprig/v3.hasKey

##  hasPrefix




Implementation: github.com/Masterminds/sprig/v3.init.func10

##  hasSuffix




Implementation: github.com/Masterminds/sprig/v3.init.func11

##  htmlDate
Formats a date for inserting into HTML date picker input field.



Implementation: github.com/Masterminds/sprig/v3.htmlDate

##  htmlDateInZone
Same as `htmlDate` but with a timezone.



Implementation: github.com/Masterminds/sprig/v3.htmlDateInZone

##  htpasswd




Implementation: github.com/Masterminds/sprig/v3.htpasswd

##  indent




Implementation: github.com/pluralsh/polly/template.indent

##  initial




Implementation: github.com/Masterminds/sprig/v3.initial

##  initials




Implementation: github.com/Masterminds/sprig/v3.initials

##  int
Converts to a `int`.



Implementation: github.com/Masterminds/sprig/v3.toInt

##  int64
Converts to a `int64`.



Implementation: github.com/Masterminds/sprig/v3.toInt64

##  isAbs




Implementation: path.IsAbs

##  join




Implementation: github.com/Masterminds/sprig/v3.join

##  kebabcase




Implementation: github.com/huandu/xstrings.ToKebabCase

##  keys




Implementation: github.com/Masterminds/sprig/v3.keys

##  kindIs




Implementation: github.com/Masterminds/sprig/v3.kindIs

##  kindOf




Implementation: github.com/Masterminds/sprig/v3.kindOf

##  last




Implementation: github.com/Masterminds/sprig/v3.last

##  list




Implementation: github.com/Masterminds/sprig/v3.list

##  lower
Converts the entire string to lowercase. `upper "HELLO"` will return `hello`.



Implementation: strings.ToLower

##  max
Returns the largest of a series of integers. `max 1 2 3` will return `3`.



Implementation: github.com/Masterminds/sprig/v3.max

##  maxf




Implementation: github.com/Masterminds/sprig/v3.maxf

##  merge




Implementation: github.com/Masterminds/sprig/v3.merge

##  mergeOverwrite




Implementation: github.com/Masterminds/sprig/v3.mergeOverwrite

##  min
Returns the smallest of a series of integers. `min 1 2 3` will return `1`.



Implementation: github.com/Masterminds/sprig/v3.min

##  minf




Implementation: github.com/Masterminds/sprig/v3.minf

##  mod




Implementation: github.com/Masterminds/sprig/v3.init.func18

##  mul
Multiples numbers. Accepts two or more inputs. `mul 1 2 3` will return `6`.



Implementation: github.com/Masterminds/sprig/v3.init.func19

##  mulf




Implementation: github.com/Masterminds/sprig/v3.init.func25

##  mustAppend




Implementation: github.com/Masterminds/sprig/v3.mustPush

##  mustChunk




Implementation: github.com/Masterminds/sprig/v3.mustChunk

##  mustCompact
Accepts a list and removes entries with empty values.



Implementation: github.com/Masterminds/sprig/v3.mustCompact

##  mustDateModify




Implementation: github.com/Masterminds/sprig/v3.mustDateModify

##  mustDeepCopy




Implementation: github.com/Masterminds/sprig/v3.mustDeepCopy

##  mustFirst




Implementation: github.com/Masterminds/sprig/v3.mustFirst

##  mustFromJson




Implementation: github.com/Masterminds/sprig/v3.mustFromJson

##  mustHas
Checks if a list has a particular element.



Implementation: github.com/Masterminds/sprig/v3.mustHas

##  mustInitial




Implementation: github.com/Masterminds/sprig/v3.mustInitial

##  mustLast




Implementation: github.com/Masterminds/sprig/v3.mustLast

##  mustMerge




Implementation: github.com/Masterminds/sprig/v3.mustMerge

##  mustMergeOverwrite




Implementation: github.com/Masterminds/sprig/v3.mustMergeOverwrite

##  mustPrepend




Implementation: github.com/Masterminds/sprig/v3.mustPrepend

##  mustPush




Implementation: github.com/Masterminds/sprig/v3.mustPush

##  mustRegexFind




Implementation: github.com/Masterminds/sprig/v3.mustRegexFind

##  mustRegexFindAll




Implementation: github.com/Masterminds/sprig/v3.mustRegexFindAll

##  mustRegexMatch




Implementation: github.com/Masterminds/sprig/v3.mustRegexMatch

##  mustRegexReplaceAll




Implementation: github.com/Masterminds/sprig/v3.mustRegexReplaceAll

##  mustRegexReplaceAllLiteral




Implementation: github.com/Masterminds/sprig/v3.mustRegexReplaceAllLiteral

##  mustRegexSplit




Implementation: github.com/Masterminds/sprig/v3.mustRegexSplit

##  mustRest




Implementation: github.com/Masterminds/sprig/v3.mustRest

##  mustReverse
Produces a new list with the reversed elements of the given list.



Implementation: github.com/Masterminds/sprig/v3.mustReverse

##  mustSlice




Implementation: github.com/Masterminds/sprig/v3.mustSlice

##  mustToDate
Converts a string to a date. The first argument is the date layout and the second is the date string. If the string can’t be converted it returns the zero value.



Implementation: github.com/Masterminds/sprig/v3.mustToDate

##  mustToJson




Implementation: github.com/Masterminds/sprig/v3.mustToJson

##  mustToPrettyJson




Implementation: github.com/Masterminds/sprig/v3.mustToPrettyJson

##  mustToRawJson




Implementation: github.com/Masterminds/sprig/v3.mustToRawJson

##  mustUniq
Generates a list with all of the duplicates removed.



Implementation: github.com/Masterminds/sprig/v3.mustUniq

##  mustWithout
Filters items out of a list.



Implementation: github.com/Masterminds/sprig/v3.mustWithout

##  must_date_modify




Implementation: github.com/Masterminds/sprig/v3.mustDateModify

##  nindent




Implementation: github.com/pluralsh/polly/template.nindent

##  nospace
Removes all whitespace from a string. `nospace "hello w o r l d"` will return `helloworld`.



Implementation: github.com/Masterminds/goutils.DeleteWhiteSpace

##  omit




Implementation: github.com/Masterminds/sprig/v3.omit

##  osBase




Implementation: path/filepath.Base

##  osClean




Implementation: path/filepath.Clean

##  osDir




Implementation: path/filepath.Dir

##  osExt




Implementation: path/filepath.Ext

##  osIsAbs




Implementation: path/filepath.IsAbs

##  pick




Implementation: github.com/Masterminds/sprig/v3.pick

##  pluck




Implementation: github.com/Masterminds/sprig/v3.pluck

##  plural




Implementation: github.com/Masterminds/sprig/v3.plural

##  prepend




Implementation: github.com/Masterminds/sprig/v3.prepend

##  push




Implementation: github.com/Masterminds/sprig/v3.push

##  quote




Implementation: github.com/Masterminds/sprig/v3.quote

##  randAlpha




Implementation: github.com/Masterminds/sprig/v3.randAlpha

##  randAlphaNum




Implementation: github.com/Masterminds/sprig/v3.randAlphaNumeric

##  randAscii




Implementation: github.com/Masterminds/sprig/v3.randAscii

##  randBytes




Implementation: github.com/Masterminds/sprig/v3.randBytes

##  randInt
Returns a random integer value from min (inclusive) to max (exclusive). `randInt 12 30` will produce a random number in the range from 12 to 30.



Implementation: github.com/Masterminds/sprig/v3.init.func20

##  randNumeric




Implementation: github.com/Masterminds/sprig/v3.randNumeric

##  regexFind




Implementation: github.com/Masterminds/sprig/v3.regexFind

##  regexFindAll




Implementation: github.com/Masterminds/sprig/v3.regexFindAll

##  regexMatch




Implementation: github.com/Masterminds/sprig/v3.regexMatch

##  regexQuoteMeta




Implementation: github.com/Masterminds/sprig/v3.regexQuoteMeta

##  regexReplaceAll




Implementation: github.com/Masterminds/sprig/v3.regexReplaceAll

##  regexReplaceAllLiteral




Implementation: github.com/Masterminds/sprig/v3.regexReplaceAllLiteral

##  regexSplit




Implementation: github.com/Masterminds/sprig/v3.regexSplit

##  repeat




Implementation: github.com/Masterminds/sprig/v3.init.func2

##  replace




Implementation: strings.ReplaceAll

##  rest




Implementation: github.com/Masterminds/sprig/v3.rest

##  reverse
Produces a new list with the reversed elements of the given list.



Implementation: github.com/Masterminds/sprig/v3.reverse

##  round
Returns a float value with the remainder rounded to the given number to digits after the decimal point. `round 123.55555 3` will return `123.556`.



Implementation: github.com/Masterminds/sprig/v3.round

##  semver




Implementation: github.com/Masterminds/sprig/v3.semver

##  semverCompare


Aliases: semver_compare

Implementation: github.com/Masterminds/sprig/v3.semverCompare

##  seq
Works like Bash `seq` command. Specify 1 parameter (`end`) to generate all counting integers between 1 and `end` inclusive. Specify 2 parameters (`start` and `end`) to generate all counting integers between `start` and `end` inclusive incrementing or decrementing by 1. Specify 3 parameters (`start`, `step` and `end) to generate all counting integers between `start` and `end` inclusive incrementing or decrementing by `step`.



Implementation: github.com/Masterminds/sprig/v3.seq

##  set




Implementation: github.com/Masterminds/sprig/v3.set

##  sha1sum




Implementation: github.com/Masterminds/sprig/v3.sha1sum

##  sha256sum


Aliases: sha26sum

Implementation: github.com/Masterminds/sprig/v3.sha256sum

##  shuffle




Implementation: github.com/huandu/xstrings.Shuffle

##  slice




Implementation: github.com/Masterminds/sprig/v3.slice

##  snakecase




Implementation: github.com/huandu/xstrings.ToSnakeCase

##  sortAlpha




Implementation: github.com/Masterminds/sprig/v3.sortAlpha

##  split




Implementation: github.com/Masterminds/sprig/v3.split

##  splitList




Implementation: github.com/Masterminds/sprig/v3.init.func13

##  splitn




Implementation: github.com/Masterminds/sprig/v3.splitn

##  squote




Implementation: github.com/Masterminds/sprig/v3.squote

##  sub




Implementation: github.com/Masterminds/sprig/v3.init.func16

##  subf




Implementation: github.com/Masterminds/sprig/v3.init.func23

##  substr




Implementation: github.com/Masterminds/sprig/v3.substring

##  swapcase




Implementation: github.com/Masterminds/goutils.SwapCase

##  ternary
Takes two values and a test value. If the test value is true, the first value will be returned. If the test value is false, the second value will be returned. This is similar to the C ternary operator. `ternary "foo" "bar" true` or `true | "foo" "bar"` will return `"foo"`.



Implementation: github.com/pluralsh/polly/template.ternary

##  title
Converts a string to title case. `title "hello world"` will return `"Hello World"`.



Implementation: strings.Title

##  toDate
Converts a string to a date. The first argument is the date layout and the second is the date string. If the string can’t be converted it returns the zero value.



Implementation: github.com/Masterminds/sprig/v3.toDate

##  toDecimal
Converts a Unix octal to a `int64`.`"0777" | toDecimal` will convert `0777` to `511` and return the value as `int64`.



Implementation: github.com/Masterminds/sprig/v3.toDecimal

##  toJson


Aliases: to_json

Implementation: github.com/Masterminds/sprig/v3.toJson

##  toPrettyJson




Implementation: github.com/Masterminds/sprig/v3.toPrettyJson

##  toRawJson
Encodes an item into JSON string with HTML characters unescaped. `toRawJson .Item` will return unescaped JSON string representation of `.Item`.



Implementation: github.com/Masterminds/sprig/v3.toRawJson

##  toString
Converts to a string.



Implementation: github.com/Masterminds/sprig/v3.strval

##  toStrings
Converts a list, slice or array to a list of strings. `list 1 2 3 | toString` converts `1`, `2` and `3` to strings and then returns them as a list.



Implementation: github.com/Masterminds/sprig/v3.strslice

##  trim
Removes space from either side of a string. `trim "  hello  "` will return `hello`.



Implementation: strings.TrimSpace

##  trimAll
Removes given characters from the front or back of a string. `trimAll "$" "$5.00"` will return `5.00` (as a string).



Implementation: github.com/Masterminds/sprig/v3.init.func4

##  trimPrefix
Trims just the prefix from a string. `trimPrefix "-" "-hello"` will return `hello`.



Implementation: github.com/Masterminds/sprig/v3.init.func6

##  trimSuffix
Trims just the suffix from a string. `trimSuffix "-" "hello-"` will return `hello`.



Implementation: github.com/Masterminds/sprig/v3.init.func5

##  trimall




Implementation: github.com/Masterminds/sprig/v3.init.func3

##  trunc




Implementation: github.com/Masterminds/sprig/v3.trunc

##  tuple




Implementation: github.com/Masterminds/sprig/v3.list

##  typeIs




Implementation: github.com/Masterminds/sprig/v3.typeIs

##  typeIsLike




Implementation: github.com/Masterminds/sprig/v3.typeIsLike

##  typeOf




Implementation: github.com/Masterminds/sprig/v3.typeOf

##  uniq
Generates a list with all of the duplicates removed.



Implementation: github.com/Masterminds/sprig/v3.uniq

##  unixEpoch
Returns the seconds since the Unix epoch.



Implementation: github.com/Masterminds/sprig/v3.unixEpoch

##  unset




Implementation: github.com/Masterminds/sprig/v3.unset

##  until
Builds a range of integers. `until 5` will return a list `[0, 1, 2, 3, 4]`.



Implementation: github.com/Masterminds/sprig/v3.until

##  untilStep
Like `until` generates a list of counting integers but it allows to define a start, stop and step. `untilStep 3 6 2` will return `[3, 5]` by starting with 3 and adding 2 until it is equal or greater than 6.



Implementation: github.com/Masterminds/sprig/v3.untilStep

##  untitle
Removes title casing. `untitle "Hello World"` will return `"hello world"`.



Implementation: github.com/Masterminds/sprig/v3.untitle

##  upper
Converts the entire string to uppercase. `upper "hello"` will return `HELLO`.



Implementation: strings.ToUpper

##  urlJoin
Joins map produced by `urlParse` to produce URL string. `urlJoin (dict "fragment" "fragment" "host" "host:80" "path" "/path" "query" "query" "scheme" "http")` will return `proto://host:80/path?query#fragment`.



Implementation: github.com/Masterminds/sprig/v3.urlJoin

##  urlParse
Parses string for URL and produces dict with URL parts. For more info check https://golang.org/pkg/net/url/#URL.



Implementation: github.com/Masterminds/sprig/v3.urlParse

##  values




Implementation: github.com/Masterminds/sprig/v3.values

##  without
Filters items out of a list. It can take more than one filter.



Implementation: github.com/Masterminds/sprig/v3.without

##  wrap
Wraps text at a given column count. `wrap 80 $text` will wrap the string in `$text` at 80 columns.



Implementation: github.com/Masterminds/sprig/v3.init.func7

##  wrapWith
Works as `wrap` but lets you specify the string to wrap with (`wrap` uses `
`). `wrapWith` 5 "\t" "Hello world"` will return `hello world` (where the whitespace is an ASCII tab character).



Implementation: github.com/Masterminds/sprig/v3.init.func8
