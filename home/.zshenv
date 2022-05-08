# Java

export JAVA_HOME_8=/usr/lib/jvm/java-8-openjdk/
export JAVA_HOME_11=/usr/lib/jvm/java-11-openjdk/
export JAVA_HOME_17=/usr/lib/jvm/java-17-openjdk/

alias java8='/usr/lib/jvm/java-8-openjdk/bin/java'
alias java11='/usr/lib/jvm/java-11-openjdk/bin/java'
alias java17='/usr/lib/jvm/java-17-openjdk/bin/java'

alias javac8='/usr/lib/jvm/java-8-openjdk/bin/javac'
alias javac11='/usr/lib/jvm/java-11-openjdk/bin/javac'
alias javac17='/usr/lib/jvm/java-17-openjdk/bin/javac'

alias javah8='export JAVA_HOME=$JAVA_HOME_8'
alias javah11='export JAVA_HOME=$JAVA_HOME_11'
alias javah17='export JAVA_HOME=$JAVA_HOME_17'

# NODE

source /usr/share/nvm/init-nvm.sh

# custom scripts

export PATH=/home/quirinecker/.scripts/:$PATH

# Android

# export ANDROID_HOME=~/Android
# export PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin

# Node Project Setup

export PATH=./node_modules/.bin:$PATH
