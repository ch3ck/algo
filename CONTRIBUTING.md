# Contributing to Algo

Want to hack on Algo? Awesome!  We have a contributor's guide that explains
how to add new features to this repository.

![Contributors guide](doc/gopher.jpg)

This page contains information about reporting issues as well as some tips and
guidelines useful to experienced open source contributors. Finally, make sure
you read through before you start participating.


## Topics

* [Data Structures](#)
* [Algorithms](#)

We have different Data structures and algorithms in mind which we created in the
issues.

Create a new 'package' for each new Data structure or Algorithm with it's unit test
integrated in the /test directory. This test should easily be called and tested
from the main package


### Conventions

Fork the repository and make changes on your fork in a feature branch:

- If it's a bug fix branch, name it XXXX-something where XXXX is the number of
	the issue. 
- If it's a feature branch, create an enhancement issue to announce
	your intentions, and name it XXXX-something where XXXX is the number of the
	issue.

Submit unit tests for your changes. Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the test on your branch before
submitting a pull request.

Update the documentation when creating or modifying features. Test your
documentation changes for clarity, concision, and correctness, as well as a
clean documentation build.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plug-ins that do this automatically.

Pull request descriptions should be as clear as possible and include a reference
to all the issues that they address.

Commit messages must start with a capitalized and short summary (max. 50 chars)
written in the imperative, followed by an optional, more detailed explanatory
text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Post
a comment after pushing. New commits show up in the pull request automatically,
but the reviewers are notified only when you comment.

Pull requests must be cleanly rebased on top of master without multiple branches
mixed into the PR.

**Git tip**: If your PR no longer merges cleanly, use `rebase master` in your
feature branch to update your pull request rather than `merge master`.


Include an issue reference like `Closes #XXXX` or `Fixes #XXXX` in commits that
close an issue. Including references automatically closes the issue on a merge.

Please do not add yourself to the `AUTHORS` file, as it is regenerated regularly
from the Git history.

Please see the [Coding Style](#coding-style) for further guidelines.


## Sign your work

The sign-off is a simple line at the end of the explanation for the patch. Your
signature certifies that you wrote the patch or otherwise have the right to pass
it on as an open-source patch. The rules are pretty simple: if you can certify
the below (from [developercertificate.org](http://developercertificate.org/)):



## Communication guidelines

We want to keep things awesome, growing and collaborative. We need
your help to keep it that way. To help with this we've come up with some general
guidelines for the community as a whole:

* Be nice: Be courteous, respectful and polite to fellow community members:
  no regional, racial, gender, or other abuse will be tolerated. We like
  nice people way better than mean ones!

* Encourage diversity and participation: Make everyone in our community feel
  welcome, regardless of their background and the extent of their
  contributions, and do everything possible to encourage participation in
  our community.

* Keep it legal: Basically, don't get us in trouble. Share only content that
  you own, do not share private or sensitive information, and don't break
  the law.

* Stay on topic: Make sure that you are posting to the correct channel and
  avoid off-topic discussions. Remember when you update an issue or respond
  to an email you are potentially sending to a large number of people. Please
  consider this before you update. Also remember that nobody likes spam.


## Coding Style

Unless explicitly stated, we follow all coding guidelines from the Go
community. While some of these standards may seem arbitrary, they somehow seem
to result in a solid, consistent codebase.

It is possible that the code base does not currently comply with these
guidelines. We are not looking for a massive PR that fixes this, since that
goes against the spirit of the guidelines. All new contributions should make a
best effort to clean up and make the code base better than they left it.
Obviously, apply your best judgement. Remember, the goal here is to make the
code base easier for humans to navigate and understand. Always keep that in
mind when nudging others to comply.

The rules:

1. All code should be formatted with `gofmt -s`.
2. Comment the code. Tell us the why, the history and the context.
3. Document _all_ declarations and methods, even private ones. Declare
   expectations, caveats and anything else that may be important. If a type
   gets exported, having the comments already there will ensure it's ready.


## Questions

If you have any further questions please feel free to reach out to [Nyah Check](nyah@ch3ck.me)
