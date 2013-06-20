clison
======


Usage
---


####Prettify
clison without any parameter will prettify the json. Below, the worst example ever, because github prettify the json by default.
```shell
curl -s https://api.github.com/legacy/repos/search/clison | clison
```
####Select value
If you want to dump the git_url of this repository to put it in a variable:
```shell
url=$(curl -s https://api.github.com/repos/gnicod/clison | clison git_url)
```
####Select value in array
You can use [\<key\>=\<value\>] to select the first occurence of this pattern in an array  
```shell
curl -s https://api.github.com/legacy/repos/search/clison | clison repositories.[language=Go]
# or 
curl -s https://api.github.com/legacy/repos/search/clison | clison "repositories.[language=Go].description"
# or even
curl -s https://api.github.com/legacy/repos/search/clison | clison repositories.0.description
```

Regexp may be coming soon
