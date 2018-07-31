## Tips 
1. `git ls-files | xargs cloc`: Count lines of code in git repo
2. `twistd -n ftp -r ./`: FTP server on port 2121; `twistd -n web -p 8080 --path .`: web server on port 8080 (`twistd -n -h` for more services like ssh, dns, mail, socks)
3. `watch -n 0 <command>`: run any command every 0.1s
4. `rsync -abviuzP src/ dest/` from https://superuser.com/a/547316/173890
    - `-i` turns on the itemized format, which shows more information than the default format
    - `-b` makes rsync backup files that exist in both folders, appending `~` to the old file. You can control this suffix with `--suffix .suf`
    - `-u` makes rsync transfer skip files which are newer in dest than in src
    - `-z` turns on compression, which is useful when transferring easily-compressible files over slow links
    - `-P` turns on `--partial` and `--progress`
    - `--partial` makes rsync keep partially transferred files if the transfer is interrupted
    - `--progress` shows a progress bar for each transfer, useful if you transfer big files
5. netstat: -l/--listening, -a/--all, -t/--tcp, -u/--udp, -n/--numeric
    - `netstat -atn` check listening TCP ports
    - `netstat -atn` check listening UCP ports
    - `netstat -atun` check listening on both ports
6. `cat /proc/sys/kernel/random/entropy_avail`: Check entropy pool size (below 200 is not good)
7. `curl -w "TCP handshake: %{time_connect}， SSL handshake: %{time_appconnect}\n" -so /dev/null https://www.google.com`
    - HTTP time: TCP handshake
    - HTTPS time: TCP handshake + SSL handshake
8. `pv`: "pipe viewer", show stats on data goi ng through a pipe
9. `sed s/foo/bar/g file.txt`, `foo` can be a regular expression
    - GNU and BSD sed behaviors differently in `-i`, so [always use](https://stackoverflow.com/a/22084103/1035859) `-i.bak`
    - `sed -n 12p` print 12th line, `sed -n 5, 30p` print liens 5-30. `-n` suppresses output so only `p`'s part gets printed 
    - `sed 5d` delete 5th line, `sed /foo/d` delete lines matching `/foo/`
    - `sed '/foo/a bar'` append 'bar' after lines containing 'foo'
10. [awk](https://coolshell.cn/articles/9070.html): `docker ps | awk {'print $1'}` print first column container ID
    - `$NF` for total number, last column
11. [less](https://twitter.com/b0rk/status/1005470181240508417): '/' search, 'n/N' next/prev match, 'j/K' down/up a line, 
    'g/G' begining/end of file
    - `less -r` display bash esxape codes as colors
    - `v` to edit file, `F` to keep reading as updated
    - `less +F` follow updates, `less +G` start 20% into file, `less +/foo` search for 'foo'
12. `fc`: fix command, open the last command you ran in an editor, and rerun
13. `date -u`: show UTC time now
14. `set -x` enables a mode of the shell where all executed commands are printed to the terminal
15. `jq`: 'fromjson' to unescape, 'tojson' to escape when parsing
16. `MY_IP=$(dig +short myip.opendns.com @resolver1.opendns.com)` [refer](https://unix.stackexchange.com/a/81699/36211)
17. `&&` in bash is "AND", statement to the left as well as right of `&&` should be run in sequence, `&` means preceding commands, 
    to the immediate left of the `&`, should simply run in the background
18. `rg -F -e` to search literal string

=======
### ripgrep
- `rg -F -e` to search literal string
- `rg -t go <keyword> -g '!vendor'`: only search in go files, ignore path contains vendor
- `rg fast README.md --replace FAST`:  replace all occurrences of `fast` with `FAST`
- `rg 'type((\s\w+\s)+)interface' -g '!vendor'`: find all interface definitions
- `rg -z/--search-zip`: search compressed files
- [Add custom type](https://github.com/BurntSushi/ripgrep/blob/master/GUIDE.md#configuration-file): --type-add web:\*.{html,css,js}\*

## References
+ [Advanced Bash-Scripting Guide](https://www.tldp.org/LDP/abs/abs-guide.pdf)
