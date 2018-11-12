## CLion
+ Multiple cursor: <kbd>Alt</kbd> + <kbd>Shift</kbd> + mouse click
+ Dash: <kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>D</kbd>

## VSCode
+ <kbd>Ctrl</kbd> + <kbd>Tab</kbd> to switch tabs, <kbd>⌘</kbd> + <kbd>Tab</kbd> + LEFT/RIGHT
+ Multiple cursor: <kbd>Alt</kbd> + <kbd>Shift</kbd> + mouse click
    - <kbd>⌘</kbd> + <kbd>Opt</kbd> + <kbd>Shift</kbd> + Up/Down to select multiple lines
+ <kbd>⌘</kbd> + Up/Down to go to head/end of file
+ <kbd>⌘</kbd> + Left/Right to go to first/last character of line
+ <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>N</kbd> to open a new window
+ <kbd>Ctrl</kbd> + <kbd>G</kbd>: go to line
+ <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>k</kbd> to delete current line
+ `"editor.minimap.enabled": false`: to disable minimap
+ [这就是我想要的 VSCode 插件！](https://zhuanlan.zhihu.com/p/36020180)
### CLI
+ `code .` : Open the current folder in a new window
+ `code -r .` : Open the current folder in the current window
+ `code -a .` : Add the current folder to the current window

## psql
+ `\l` or `\list`: list all databases
+ `\c db_name`: connect to database
+ `\q`: disconnect
+ `\conninfo`
+ `\dt`: describe tables
+ `\d table_name`: describe table schema, `\d+` for more
+ `\dF`: describe full text flag, `dF+` for more detailed info

## mysql
+ `SHOW [FULL] PROCESSLIST;`: useful commands to show running threads and connections
+ In MySQL, use "utf8mb4" (real utf-8) instead of "utf8"

## sqlite3
+ `.schema`: Print the database structure
+ `.explain on`: Turn on column names on query results`

## Chrome
+ <kbd>Shift</kbd> + <kbd>fn</kbd> + <kbd>Delete</kbd> Delete address bar suggestion

# Supervisord
```
# /etc/supervisor/conf.d/*.conf
[program:test1]
command=python test1.py
directory=/path/to/test1_dir
user=wilbeibi
autorestart=true
redirect_stderr=true
stdout_logfile=/var/log/test1.log
```

## Alfred workflows
+ `dash python3: os.path`

## Emoji
🛢️ 📰 💾 🖥️ 📱 💣 🚚 🚢 🗿 🚀 🏄 🔧 ☕️ 🦊 🌊

## Mac tools
+ mac2imgur
+ Mac 10.12 不允许第三方来源的app安装了，要`sudo spctl --master-disable`
+ [Monodraw](https://monodraw.helftone.com/): ASCII diagram for illustration, banners
+ <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>Ctrl</kbd> + <kbd>4</kbd>: copy picture of selected area(screenshot) to the clipboard
+ `sshfs <user>@<host>:/remote/path /mount/point` to mount remote filesystem using SSH SFTP.
+ ![命令行自动切换英文输入法](https://i.imgur.com/XQvkiYI.png)
+ Briss2: pdf chop tool
+ `brew switch <formula> <version>`
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
## iTerm2
+ [delete world/line in iTerm2](https://coderwall.com/p/ds2dha/word-line-deletion-and-navigation-shortcuts-in-iterm2): deleting a word: 0x17, deleting a line: 0x15
## Git
### Stash [ref](https://gist.github.com/subchen/3409a16cb46327ca7691)
+ `git stash save "message..."`: Save current work progress, staged/unstaged
+ `git stash list`
+ `git stash pop [--index] [<stash>]`: if not use parameter, recover the latest progress, and delete it from stash
+ `git stash apply [--index] [<stash>]`: same as pop, just not delete it.
+ `git stash clear`
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
### git gc
+ `git gc --aggressive` to compress repo.
### store credentials
+ `git config credential.helper store` to avoid keep been asked for username and password
### git console show non-ascii path name
+ `git config --global core.quotepath off`, [refer](https://stackoverflow.com/a/22828826/1035859)

## Perf
+ [fio: Flexible I/O Tester](https://github.com/axboe/fio)
+ [filebench: generated storage benchmark workloads](https://github.com/filebench/filebench/wiki)


## Mise
+ [My Diigo Programming Notes](https://www.diigo.com/outliner/dzi0kh/Programming?key=a7q47wq9b2)
+ [LaTex cheat sheet](https://wch.github.io/latexsheet/)
+ Debug Chrome CPU hogging: Invoke `Chrome Menu / More Tools / Task Manager` to see what consumes CPU
+ <kbd>Fn</kbd> + <kbd>Q</kbd> to bluetooth pairing HHKB to Mac
+ `neofetch` to show system info