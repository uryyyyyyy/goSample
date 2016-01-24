
```

# install dependencies
sudo apt-get install curl git mercurial make binutils bison gcc build-essential

# install gvm
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.bashrc

# all available version
gvm listall

# install 1.5
gvm install 1.5.3

# error...

# install 1.4 first
gvm install 1.4.3
gvm use go1.4.3

# then install 1.5
gvm install 1.5.3
gvm use go1.5.3

# check
go version

```
