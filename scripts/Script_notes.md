## Tips 
- [deprecated]`git ls-files | xargs cloc`: Count lines of code in git repo
- `twistd -n ftp -r ./`: FTP server on port 2121; `twistd -n web -p 8080 --path .`: web server on port 8080 (`twistd -n -h` for more services like ssh, dns, mail, socks)
- `watch -n 0 <command>`: run any command every 0.1s
- `rsync -abviuzP src/ dest/` from https://superuser.com/a/547316/173890
    - `-i` turns on the itemized format, which shows more information than the default format
    - `-b` makes rsync backup files that exist in both folders, appending `~` to the old file. You can control this suffix with `--suffix .suf`
    - `-u` makes rsync transfer skip files which are newer in dest than in src
    - `-z` turns on compression, which is useful when transferring easily-compressible files over slow links
    - `-P` turns on `--partial` and `--progress`
    - `--partial` makes rsync keep partially transferred files if the transfer is interrupted
    - `--progress` shows a progress bar for each transfer, useful if you transfer big files
- netstat: -l/--listening, -a/--all, -t/--tcp, -u/--udp, -n/--numeric -p/--program
    - `netstat -atn` check listening TCP ports
    - `netstat -atn` check listening UCP ports
    - `netstat -tunpal` check listening on both ports for TCP/UDP of program(pid)
    - `ss -tunpal` newer version of netstat (remember as "tuna, please") from [b0rk](https://twitter.com/b0rk/status/1090058524137345025)
- `cat /proc/sys/kernel/random/entropy_avail`: Check entropy pool size (below 200 is not good)
- `curl -w "TCP handshake: %{time_connect}， SSL handshake: %{time_appconnect}\n" -so /dev/null https://www.google.com`
    - HTTP time: TCP handshake
    - HTTPS time: TCP handshake + SSL handshake
- `pv`: "pipe viewer", show stats on data going through a pipe
- `sed s/foo/bar/g file.txt`, `foo` can be a regular expression
    - GNU and BSD sed behaviors differently in `-i`, so [always use](https://stackoverflow.com/a/22084103/1035859) `-i.bak`
    - `sed -n 12p` print 12th line, `sed -n 5, 30p` print liens 5-30. `-n` suppresses output so only `p`'s part gets printed 
    - `sed 5d` delete 5th line, `sed /foo/d` delete lines matching `/foo/`
    - `sed '/foo/a bar'` append 'bar' after lines containing 'foo'
- [awk](https://coolshell.cn/articles/9070.html): `docker ps | awk {'print $1'}` print first column container ID
    - `$NF` for total number, last column
- [less](https://twitter.com/b0rk/status/1005470181240508417): '/' search, 'n/N' next/prev match, 'j/K' down/up a line, 
    'g/G' begining/end of file
    - `less -r` display bash escape codes as colors
    - `v` to edit file, `F` to keep reading as updated
    - `less +F` follow updates, `less +G` start 20% into file, `less +/foo` search for 'foo'
- [bat](https://github.com/sharkdp/bat): `alias lessy=bat -l yaml -p` as less alternative for yaml
- `fc`: fix command, open the last command you ran in an editor, and rerun
- `date -u`: show UTC time now
- `set -x` enables a mode of the shell where all executed commands are printed to the terminal
- `MY_IP=$(dig +short myip.opendns.com @resolver1.opendns.com)` [refer](https://unix.stackexchange.com/a/81699/36211)
- `&&` in bash is "AND", statement to the left as well as right of `&&` should be run in sequence, `&` means preceding commands, 
    to the immediate left of the `&`, should simply run in the background
- `dig -t A wilbeibi.com +trace` trace DNS resolve
- `vim scp://user@server/~/data.txt`: edit file on remote server
-  `tr : '\n' <<<$PATH`: show directories in PATH, one per line (`<<<` for redirects the string to stdin of the command), or `echo $PATH|tr ":" "\n"`
- `ssh -L 80:localhost:8080 remotehost`: browser visit localhost:80 is actually get response from remotehost:8080
- `ssh -R 8080:localhost:80 remotehost`: vice versa
- [ssh jump host config](https://gist.github.com/wilbeibi/1505fcd81f7376cdd91cd370d2dd9204)
- [ncdu](https://dev.yorhel.nl/ncdu): check disk usage interactively
- [moreutils: the utilities package every UNIX/Linux/Mac OS developer should know](https://rentes.github.io/unix/utilities/2015/07/27/moreutils-package/)
- [rga: ripgrep, but also search in PDFs, E-Books, Office documents, zip, tar.gz, etc](https://phiresky.github.io/blog/2019/rga--ripgrep-for-zip-targz-docx-odt-epub-jpg/)
- double dash in command signify the end of command options [ref](signify the end of command options)
- [just](https://github.com/casey/just): makefile better alternative, can used to save recipes as well
## ZSH
+ `cd /u/l/b`: path expansion
+ `cd site1 site2`: path replacement, if you were in /srv/www/site1/current/log, it will go to /srv/www/site2/current/log via this command
+ `ls -l **/*.log`: extended globbing, **/ = recursive
+ `zmv '(*).txt' 'template_$1.html'`: rename files
+ `ls -l zsh_demo/**/*(. Lm-2 mh-1 om)`: ls files under zsh_demo recursively, `Lm-2` for less than 2mb (similarly, `Lm+30` for over 30mb, m for megabytes, k for kilobytes, or nothing for just bytes), `mh-1` for files modified in the last hour (M for Months, w for weeks, h for hours, m for minutes, and s for seconds), `om` to o(rder) by modification date, o for most recent, O vise versa, m for modification date, or L to sort by size.
+ <kbd>Ctrl</kbd> + <kbd>X</kbd> + <kbd>Ctrl</kbd> + <kbd>E</kbd>  to edit long command
+ [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions/blob/master/INSTALL.md), → to complete
+ [Master Your Z Shell with These Outrageously Useful Tips](http://reasoniamhere.com/2014/01/11/outrageously-useful-tips-to-master-your-z-shell/)
+ `sh -c "$(wget https://example.com/execute.sh -0 -)"`: download and execute command
+ <kbd>⌘</kbd> + <kbd>.</kbd>: recall the last argument of any of the previous commands
+ theme: spaceship, font: 14pt Fira Mono for Powerline

## jq
- `jq`: 'fromjson' to unescape, 'tojson' to escape when parsing

## lnav: log viewer


## Git
### Common
+ GIT_BRANCH="$(git rev-parse --abbrev-ref HEAD)" to get current git branch
+ COMMIT_HASH_SHORT="$(git rev-parse --short HEAD)" to get current commit id
+ `git clean -fXd` to remove untracked files and directories [ref](https://stackoverflow.com/a/64966/1035859)
+ `git fetch origin; git rebase -i origin/master; git push --force-with-lease origin dev` in dev branch to rebase from master
+ `git commmit --fixup <fix to which commit>` and `git rebase -i --autosquash <prev commit>`, refer [fixup and autosquash](https://fle.github.io/git-tip-keep-your-branch-clean-with-fixup-and-autosquash.html)
### `tig` interactive commits viewer in cli
+ [manual](https://jonas.github.io/tig/)
### Clean commits [ref](https://about.gitlab.com/2018/06/07/keeping-git-commit-history-clean/)
#### Change most recent commit
+ `git add file_to_change.go`
+ `git commit --amend` to amend changes in stage to most recent commit
+ `git push --force-with-lease <remote_name> <branch_name>` to push (about [--force-with-lease](https://developer.atlassian.com/blog/2015/04/force-with-lease/))
#### Change a specific commit
+ `git rebase -i <one commit before the one to modify>`, then in interactive editor, change first `pick <commit to modify> <description>` to `edit <commit to modify> <description>`
+ modify file_to_change.go, `git add file_to_change.go`
+ `git rebase --continue`
+ `git push --force-with-lease <remote_name> <branch_name>`
#### Add, remove, or combine commits
+ `git rebase -i <commit>`
+ change `pick` to `squash`(keeps the commits messages in the description), `fixup`(forget the commit messages of the fixes and keep the original) or `drop`(delete that commit)
### Delete branches [ref](https://stackoverflow.com/a/46412667/1035859)
+ `git push origin --delete <branch_name>`: delete remote branch
+ `git branch -D <branch>`: delete local branch
### config per repository [ref](https://stackoverflow.com/questions/18181439/git-different-config-for-different-repository)
### Change commit author
+ `git commit --amend --author="Author Name <email@address.com>"`
### git gc is bad, use repack instead
+ `git repack -a -d --depth=250 --window=250` [refer Linus](http://gcc.gnu.org/ml/gcc/2007-12/msg00165.html)
### store credentials
+ `git config credential.helper store` to avoid keep been asked for username and password
### git console show non-ascii path name
+ `git config --global core.quotepath off`, [refer](https://stackoverflow.com/a/22828826/1035859)
### update upstream from forked repo
+ `git remote add upstream <original repo>.git`
+ `git pull upstream master --rebase`
+ Always `git pull --rebase`: [refer](https://coderwall.com/p/yf5-0w/like-git-pull-rebase-make-it-default)
### global gitignore
+ `git config --global core.excludesfile ~/.gitignore` ('.vscode', '.idea', '.DS_Store')
### stash recipes
+ `git stash save <msg>`: save a stash with a message
+ `git stash -u/--include-untracked`: save untracked files
+ `git stash -p/--patch`: stash just a single file
+ `git stash show -p` show a stash or `git stash show` view the latest stash
+ `gits stash drop <stash_id>` or `git stash clear` to remove stash(s)
### alias
+ `lg = log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr)%C(bold blue)<%an>%Creset' --abbrev-commit`
+ `hl =  "!f() { git log --pretty=short -u -L \"$1\",\"$1\":\"$2\"; }; f"`

## ripgrep & silver searcher
- `rg -F -e` to search literal string
- `ag --go --ignore vendor --ignore mod --ignore '*_test.go' <keyword>`: only search in go files, ignore path contains vendor and mod
- `rg fast README.md --replace FAST`:  replace all occurrences of `fast` with `FAST`
- `rg 'type((\s\w+\s)+)interface' -g '!vendor'`: find all interface definitions suit pattern "type <interface_name> interface"
- `rg "defer (.+)Rollback"`: match string like "defer tx.Rollback()"
- `rg -z/--search-zip`: search compressed files
- [Add custom type](https://github.com/BurntSushi/ripgrep/blob/master/GUIDE.md#configuration-file): --type-add web:\*.{html,css,js}\*
### Third party git tools
- [Git Interactive Rebase Tool](https://github.com/MitMaro/git-interactive-rebase-tool)
- ⭐ [tj/git-extras](https://github.com/tj/git-extras)
- ⭐ [nvie/git-toolbelt](https://github.com/nvie/git-toolbelt)

## Good code
- [docker-gc](https://github.com/spotify/docker-gc/blob/master/docker-gc)
## References
+ [Advanced Bash-Scripting Guide](https://www.tldp.org/LDP/abs/abs-guide.pdf)
+ :star: [special dollar sign shell variables](https://stackoverflow.com/questions/5163144/what-are-the-special-dollar-sign-shell-variables)
