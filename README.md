# notr

notr is a lightweight cli tool for keeping track of problems and their solutions.

The basic use case is that whenever there is a problem the user can open a problem and then close it later once they have the solution without having to open another application.

This hopes to solve the times when we say ... 

- "I'm sure I'll remember how I fixed x"
- "How did I solve y problem?"
- "How did I fix that error during my builds?"

The list goes on and on ...

## How to install
Run the install script inside the directory.
Following that add the following lines in your .bashrc, .zshrc, or whatever you use. 
```
# Add the user's private bin if it exists
if [ -d "$HOME/bin" ] ; then
  PATH="$HOME/bin:$PATH"
fi
```

From there open up a terminal and you should be able to run the command below to see how the app works!
```
notr --help
```

## How to contribute
1. Fork the repo
2. Create an issue with your planned contribution
3. If you have been assigned an issue, name the branch like `<issue_label>/some-descriptive-name`
4. Write tests if necessary for the changes you have created
5. Push to the branch `git push origin <issue_label>/some-descriptive-name`
6. Create a Pull Request

### Contribution practices 
In order to maintain consistency in the workflow here are some guidelines to follow. PRs and Issues will be declined/asked to be changed if they are not followed. Any improvements to this are encouraged and welcomed!

- Commit messages and PRs should start with some sort of action "Adds, Updates, Removes etc". Please see the example commits as an example as well as [this](https://chris.beams.io/posts/git-commit/)
- If the issue addresses something that is visible to the user please add "feature" otherwise add "enhancement" as a label
- Issues that are straightforward please add "good first issue"
- When making a PR please add in the description or comment which issue it resolves. Something like "Resolves #3".
- If your branch has multiple commits that do the same logical thing ie. edits the README twice then please squash into one
