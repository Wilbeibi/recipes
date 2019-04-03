#!/bin/bash

# -r: do not allow backslashes to escape any characters
# -s: do not echo input coming from a terminal
# -p output the string PROMPT without a trailing newline before attempting to read

passwd=$1
read -r -s -p "Enter password: " passwd
echo -e "\nPassword is $passwd"
