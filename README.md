# WordSalt

This Go module provides a simple function to generate random security keys for WordPress installations. It generates all 8 keys defined in `wp-config.php`:

- `AUTH_KEY`
- `SECURE_AUTH_KEY`
- `LOGGED_IN_KEY`
- `NONCE_KEY`
- `AUTH_SALT`
- `SECURE_AUTH_SALT`
- `LOGGED_IN_SALT`
- `NONCE_SALT`

## Installation

```sh
go install github.com/iamcryptoki/wordsalt@v1.0.1
```

## Usage

### Command-line

Simply run `wordsalt` in your terminal. The output will be the 8 WordPress security keys, each with a define() statement ready to be copied into your `wp-config.php` file.

Example output:

```php
define('AUTH_KEY', 'pUR_8n#vDLgRLn+lA<j>0CK$5mF^US}>,99Tl  eb$)ASF^Ev1P7dMX)X2KJ-t$[');
define('SECURE_AUTH_KEY', 'UrLx|hiUjz76|sH69 f8|C^nbhQ8kC{VyMU>&NX8fI6K<y49I`2|znBD(b9R>vg9');
define('LOGGED_IN_KEY', 'J4{g.Qj.~n0:_)-ie5VW83Vo$5qHy+mt2kwCff&jLZfl{NPoKqc-%@+Q[5:ZaQE8');
define('NONCE_KEY', 'v&vKo)IG|K`*Vj{Z609C%GW$k`z{G0/Ri;I/A@~Kv-RT!+TIRMKQ-vei;n zpH&h');
define('AUTH_SALT', ')E5u#-j,VhE%obE}~YNDTE)%~C8^$t.2X ^2KubHw.<H[VKC[0{Modn=c&$g%q*T');
define('SECURE_AUTH_SALT', 'hp[QTxp3I PRr+!DdRij5-W_< H;~H]_bYu5skGn9{f2S6ZknM1h |]yS{a^t>d`');
define('LOGGED_IN_SALT', 'x6c(`Q(&g_>b7pa2fue$M$#1*]Pe[*}ETBpqr8aRUA/Pe$=;zO;L,vXZQU[):5Y?');
define('NONCE_SALT', 'W}ey@Fn0TG~f.w55$8@Qa<f_t|GT/$*Ykh2e=?r@;[2X;G*wW.2)D`62.*(||&}l');
```

### External module

Import the module into your Go project:

```go
import "github.com/iamcryptoki/wordsalt"
```

Then, call the `GenerateWordPressKeys()` function to generate a map of key names and values:

```go
keys, err := wordsalt.GenerateWordPressKeys()
if err != nil {
    fmt.Println("Error generating keys:", err)
}

for name, key := range keys {
    fmt.Printf("define('%s', '%s');\n", name, key)
}
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

