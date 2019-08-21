## CLion
+ Multiple cursor: <kbd>Alt</kbd> + <kbd>Shift</kbd> + mouse click
+ Dash: <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>D</kbd>
+ Peek Definition: <kbd>⌘</kbd> + <kbd>Y</kbd>
+ `Save Actions`: format on save
+ Move Caret to Code Block End/Start: <kbd>⌥</kbd><kbd>⌘</kbd> + <kbd>[</kbd> / <kbd>]</kbd>
+ Go to line: <kbd>⌘</kbd> + <kbd>L</kbd>
+ Jump to definition: <kbd>⌥</kbd><kbd>⌘</kbd> + <kbd>B</kbd>

## VSCode
+ <kbd>Ctrl</kbd> + <kbd>Tab</kbd> to switch tabs, <kbd>⌘</kbd> + <kbd>Tab</kbd> + LEFT/RIGHT
+ Multiple cursor: <kbd>Alt</kbd> + <kbd>Shift</kbd> + mouse click
    - <kbd>⌘</kbd> + <kbd>Opt</kbd> + <kbd>Shift</kbd> + Up/Down to select multiple lines
+ <kbd>⌘</kbd> + Up/Down to go to head/end of file
+ <kbd>⌘</kbd> + Left/Right to go to first/last character of line
+ <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>N</kbd> to open a new window
+ <kbd>⌘</kbd> + <kbd>+</kbd>/<kbd>-</kbd> to zoom fonts
+ <kbd>Ctrl</kbd> + <kbd>G</kbd>: go to line
+ <kbd>⌘</kbd> + <kbd>Shift</kbd> + <kbd>k</kbd> to delete current line
+ <kbd>Shift</kbd> + <kbd>Ctrl</kbd> + <kbd>`</kbd> to open terminal
+ `"editor.minimap.enabled": false`: to disable minimap
+ [这就是我想要的 VSCode 插件！](https://zhuanlan.zhihu.com/p/36020180)
+ <kbd>⌘</kbd> + click `rg` search result's file name, can open the file in VSCode
+ [Remote - SSH](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-ssh): coding on SSH server
+ [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers): coding in container environment
+ Command line
    + `alias code="/usr/local/bin/code-insiders"`: type `code` instead of `code-insiders`
    + `code -d file1 file2`: Compare file1 and file2 side by side
    + `code -g file:linenum`: Open file on the specific line
### shell integration
+ `code .` : Open the current folder in a new window
+ `code -r .` : Open the current folder in the current window
+ `code -a .` : Add the current folder to the current window

### IntelliJ
+ GitToolBox Plugin

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
+ "chrome://flags" allow tab groups

## Firefox
+ "%<keyword>" in address bar to search in open tabs
### Plugins
+ [Play/Pause](https://addons.mozilla.org/en-US/firefox/addon/play-pause)

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
+ [Tableplus](https://tableplus.io/): Database management
+ Briss2: pdf chop tool
+ `brew switch <formula> <version>`
+ [homebrew-cask-upgrade](https://github.com/buo/homebrew-cask-upgrade)
+ [Fliqlo: Mac 时钟壁纸](https://fliqlo.com)
+ [PodcastMenu: Overcast on Mac](https://github.com/insidegui/PodcastMenu)

## iTerm2
+ [delete world/line in iTerm2](https://coderwall.com/p/ds2dha/word-line-deletion-and-navigation-shortcuts-in-iterm2): deleting a word: 0x17, deleting a line: 0x15
+ `Preference/Advanced`, search `Tabs`, select `Add new tabs at the end of the tab bar, not next to current tab` as `No`

## Dash
+ "Third-party sources" -> "Go Docsets", can add open source godocs

## Perf
+ [fio: Flexible I/O Tester](https://github.com/axboe/fio)
+ [filebench: generated storage benchmark workloads](https://github.com/filebench/filebench/wiki)
+ [Instruments: Mac developer tool](https://help.apple.com/instruments/mac/10.0/#/dev7b09c84f5)


## Mise
+ [My Diigo Programming Notes](https://www.diigo.com/outliner/dzi0kh/Programming?key=a7q47wq9b2)
+ [LaTex cheat sheet](https://wch.github.io/latexsheet/)
+ Debug Chrome CPU hogging: Invoke `Chrome Menu / More Tools / Task Manager` to see what consumes CPU
+ <kbd>Fn</kbd> + <kbd>Q</kbd> to bluetooth pairing HHKB to Mac
+ `neofetch` to show system info
+ [search my gists](https://gist.github.com/search?utf8=%E2%9C%93&q=user%3Awilbeibi&ref=searchresults)