# Requirements

First install libinjection:

```sh
git clone https://github.com/libinjection/libinjection
gcc -std=c99 -Wall -Werror -fpic -c libinjection/src/libinjection_sqli.c -o libinjection/libinjection_sqli.o 
gcc -std=c99 -Wall -Werror -fpic -c libinjection/src/libinjection_xss.c -o libinjection/libinjection_xss.o
gcc -std=c99 -Wall -Werror -fpic -c libinjection/src/libinjection_html5.c -o libinjection/libinjection_html5.o
gcc -dynamiclib -shared -o libinjection/libinjection.so libinjection/libinjection_sqli.o libinjection/libinjection_xss.o libinjection/libinjection_html5.o
cp libinjection/*.so /usr/local/lib
cp libinjection/*.o /usr/local/lib
cp libinjection/src/*.h /usr/local/include/
chmod 444 /usr/local/include/libinjection*
ldconfig
```
