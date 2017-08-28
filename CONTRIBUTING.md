# Contributing to Docker CE

Want to contribute on Docker CE? Awesome!

This page contains information about reporting issues as well as some tips and
guidelines useful to experienced open source contributors. Finally, make sure
you read our [community guidelines](#docker-community-guidelines) before you
start participating.

## Topics

* [Reporting Security Issues](#reporting-security-issues)
* [Reporting Issues](#reporting-other-issues)
* [Submitting Pull Requests](#submitting-pull-requests)
* [Community Guidelines](#docker-community-guidelines)

## Reporting security issues

The Docker maintainers take security seriously. If you discover a security
issue, please bring it to their attention right away!

Please **DO NOT** file a public issue, instead send your report privately to
[security@docker.com](mailto:security@docker.com).

Security reports are greatly appreciated and we will publicly thank you for it.
We also like to send gifts&mdash;if you're into Docker schwag, make sure to let
us know. We currently do not offer a paid security bounty program, but are not
ruling it out in the future.

## Reporting other issues

There are separate issue-tracking repos for the end user Docker CE
products specialized for a platform. Find your issue or file a new issue
for the platform you are using:

* https://github.com/docker/for-linux
* https://github.com/docker/for-mac
* https://github.com/docker/for-win
* https://github.com/docker/for-aws
* https://github.com/docker/for-azure

When reporting issues, always include:

* The output of `docker version`.
* The output of `docker info`.

If presented with a template when creating an issue, please follow its directions.

Also include the steps required to reproduce the problem if possible and
applicable. This information will help us review and fix your issue faster.
When sending lengthy log-files, consider posting them as a gist (https://gist.github.com).
Don't forget to remove sensitive data from your logfiles before posting (you can
replace those parts with "REDACTED").

## Submitting pull requests

Please see the corresponding `CONTRIBUTING.md` file of each component for more information:

* Changes to the `engine` should be directed upstream to https://github.com/moby/moby
* Changes to the `cli` should be directed upstream to https://github.com/docker/cli
* Changes to the `packaging` should be directed upstream to https://github.com/docker/docker-ce-packaging

## Docker community guidelines

We want to keep the Docker community awesome, growing and collaborative. We need
your help to keep it that way. To help with this we've come up with some general
guidelines for the community as a whole:

* Be nice: Be courteous, respectful and polite to fellow community members.
  Regional, racial, gender, or other abuse will not be tolerated. We like
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

* Don't send email to the maintainers: There's no need to send email to the
  maintainers to ask them to investigate an issue or to take a look at a
  pull request. Instead of sending an email, GitHub mentions should be
  used to ping maintainers to review a pull request, a proposal or an
  issue.

### Guideline violations — 3 strikes method

The point of this section is not to find opportunities to punish people, but we
do need a fair way to deal with people who are making our community suck.

1. First occurrence: We'll give you a friendly, but public reminder that the
   behavior is inappropriate according to our guidelines.

2. Second occurrence: We will send you a private message with a warning that
   any additional violations will result in removal from the community.

3. Third occurrence: Depending on the violation, we may need to delete or ban
   your account.

**Notes:**

* Obvious spammers are banned on first occurrence. If we don't do this, we'll
  have spam all over the place.

* Violations are forgiven after 6 months of good behavior, and we won't hold a
  grudge.

* People who commit minor infractions will get some education, rather than
  hammering them in the 3 strikes process.

* The rules apply equally to everyone in the community, no matter how much
        you've contributed.

* Extreme violations of a threatening, abusive, destructive or illegal nature
        will be addressed immediately and are not subject to 3 strikes or forgiveness.

* Contact abuse@docker.com to report abuse or appeal violations. In the case of
        appeals, we know that mistakes happen, and we'll work with you to come up with a
        fair solution if there has been a misunderstanding.