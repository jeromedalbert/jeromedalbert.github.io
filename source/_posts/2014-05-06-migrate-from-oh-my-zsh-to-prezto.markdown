---
layout: post
title: "Migrate from Oh-my-zsh to Prezto"
date: 2014-05-06 22:44
comments: true
categories: 
---

{% img header https://farm6.staticflickr.com/5553/14847770315_4dc2f5c4ef_m.jpg Pipe the whole result to the "lolcat" Ruby gem for better effect. %}

Zsh is an amazing shell, and <a href="https://github.com/robbyrussell/oh-my-zsh" target="blank">Oh-my-zsh</a> is very good at showcasing its power without having to dive into too much configuration.
I recommend it to Zsh first-timers: within minutes they can enjoy autocompletion on steroids, crazy prompt tweaking, git integration as well as many other plugins.

BUT. It's God. Damn. Slow!<!--more-->

Wanna spend the night happily coding and hacking in the "zone"?
Here, wait for 1 long second (or more) whenever you open up a new terminal window/tab.
Are you triggering tab autocompletion? Oh boy. Come back tomorrow.

Enter <a href="https://github.com/sorin-ionescu/prezto" target="blank">Prezto</a>, a.k.a the "Instantly Awesome Zsh". It began as a fork of Oh-my-zsh, but was later completely rewritten with optimization in mind, and moved to its own repo.  

Here's how I migrated in order to have the exact same behavior as before, while still having a fast shell. Yes, I like to have my cake and eat it too.

## Installation

Very straightforward:

- Backup your existing configuration: `~/.oh-my-zsh`, `~/.zshrc` and your possible other dotfiles
- Uninstall Oh-my-zsh by running `uninstall_oh_my_zsh`
- Install Prezto by following the quick instructions on the <a href="https://github.com/sorin-ionescu/prezto" target="blank">github repo</a>

Done. You now have a bunch of symlinked `.z<something>` dotfiles in your home folder.

## Fiddling with the dotfiles

- If it annoys you, remove the possible example instructions in `~/.zlogin` and `~/.zlogout` (such as displaying a fortune every time you login)
- Put back all your aliases and other customizations into `~/.zshrc`:

``` sh
source ~/.zprezto/init.zsh

export SUBLIME=subl
export EDITOR="$SUBLIME --wait"
export VISUAL=$EDITOR
export LSCOLORS="exfxcxdxbxegedabagacad"
...

alias ll="ls -lh"
alias la="ls -A"
alias gl="git pull --rebase"
...
```

- If you like to organize your exports/aliases/etc into separate files, have a look at the other dotfiles. I personally don't do that and just have one massive `~/.zshrc`.

## Customizing the prompt

Now to the fun part!  
Prezto contains half a dozen <a href="https://github.com/sorin-ionescu/prezto#themes" target="blank">example themes</a> compared to the 140 (!) ones of Oh-my-zsh. If you have a custom prompt like me, you will need to tinker a bit.

My theme is quite minimalist: just display the current path followed by `$` and a space.
When in a Git repo, I like to display the branch name followed by a dirty/clean indicator like so:

{% img https://farm6.staticflickr.com/5562/14685304517_58427bb488_o.png Dirty vs Clean %}

At the time of writing, most Prezto themes use `vcs_info` which can detect staged and unstaged changes (i.e. a dirty state).
But there's no hook for a clean state, so when there are no changes, the prompt will just display nothing.
I would much rather have a satisfying green little tick that means "Nothing else to do! Well done mate, you can have a beer\* now".

For this, we need to use `git-info` from the Prezto `git` plugin. I ended up with this:

```
function prompt_jerome_precmd {
  git-info
}

function prompt_jerome_setup {
  setopt LOCAL_OPTIONS
  unsetopt XTRACE KSH_ARRAYS
  prompt_opts=(cr percent subst)

  # Load required functions.
  autoload -Uz add-zsh-hook

  # Add hook for calling git-info before each command.
  add-zsh-hook precmd prompt_jerome_precmd

  # Set git-info parameters.
  zstyle ':prezto:module:git:info' verbose 'yes'
  zstyle ':prezto:module:git:info:branch' format '%F{green}%b%f'
  zstyle ':prezto:module:git:info:clean' format ' %F{green}✔%f'
  zstyle ':prezto:module:git:info:dirty' format ' %F{red}✗%f'
  zstyle ':prezto:module:git:info:keys' format \
    'prompt' ' %F{green}(%f$(coalesce "%b" "%p" "%c")${git_info[rprompt]}%s%F{green})%f' \
    'rprompt' '%C%D'

  # Define prompts.
  PROMPT='%~${(e)git_info[prompt]}$ '
  RPROMPT=''
}

prompt_jerome_setup "$@"
```

I then saved the theme in `~/.zprezto/modules/prompt/functions/prompt_jerome_setup`
and made Prezto use it by having this line in `~/.zpreztorc`:
``` sh
zstyle ':prezto:module:prompt' theme 'jerome'
```

Done.

## Last bits and pieces

- I added `history-substring-search` and `git` to the `~/.zpreztorc` plugin list:
``` sh
...
zstyle ':prezto:load' pmodule \
  'environment' \
  'terminal' \
  'editor' \
  'history' \
  'directory' \
  'spectrum' \
  'utility' \
  'completion' \
  'prompt' \
  'history-substring-search' \
  'git'
...
```

`history-substring-search`: when pressing up/down arrows, completes the beginning of a command by searching in the history.  
`git`: adds useful aliases and functions like the `git-info` one that we discussed above.

- When pressing tab to complete a command, I didn't like the category menus, so I commented these lines out in `~/.zprezto/modules/completion/init.zsh`:
```sh
...
# Group matches and describe.
zstyle ':completion:*:*:*:*:*' menu select
# zstyle ':completion:*:matches' group 'yes'
# zstyle ':completion:*:options' description 'yes'
# zstyle ':completion:*:options' auto-description '%d'
# zstyle ':completion:*:corrections' format ' %F{green}-- %d (errors: %e) --%f'
# zstyle ':completion:*:descriptions' format ' %F{yellow}-- %d --%f'
# zstyle ':completion:*:messages' format ' %F{purple} -- %d --%f'
# zstyle ':completion:*:warnings' format ' %F{red}-- no matches found --%f'
# zstyle ':completion:*:default' list-prompt '%S%M matches%s'
# zstyle ':completion:*' format ' %F{yellow}-- %d --%f'
# zstyle ':completion:*' group-name ''
# zstyle ':completion:*' verbose yes
...
```

I also modified this line in order to have red diretories when autocompleting the `cd` command:
```
# Directories
zstyle ':completion:*:default' list-colors ''
```

- In my `~/.zshrc`, I added these zsh options:
```sh
unsetopt CORRECT                      # Disable autocorrect guesses. Happens when typing a wrong
                                      # command that may look like an existing one.

expand-or-complete-with-dots() {      # This bunch of code displays red dots when autocompleting
  echo -n "\e[31m......\e[0m"         # a command with the tab key, "Oh-my-zsh"-style.
  zle expand-or-complete
  zle redisplay
}
zle -N expand-or-complete-with-dots
bindkey "^I" expand-or-complete-with-dots
```

## Wrapping up

That's it, you now have a fast zsh that behaves like before!

I skipped some very minor configurations in this article, so if you want to review exactly what I've done (or if you want to know what kind of crazy stuff I put in my `~/.zshrc`),
have a look at the commits in my <a href="https://github.com/jeromedalbert/prezto/commits/master" target="blank">Prezto fork</a>.

Enjoy!

*EDIT: <a href="https://news.ycombinator.com/item?id=8158707" target="blank">Hacker News thread</a>*

<div class="references">
*I prefer cider.
</div>
