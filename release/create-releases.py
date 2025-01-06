import git

from github.Repository import Repository

from sys import argv

from changelogs import parse_changelog, Changelog, Version
from gh import get_remote_repository, get_github_client


def read_changelog(path: str) -> Changelog:
    with open(path, "r") as file:
        content = file.read()

    return parse_changelog(content)
    

def get_release_message(version: Version) -> str:
    return "\n\n".join(map(lambda cl: cl.to_markdown(), version.changelists))


def create_releases(repo: Repository, changelog: Changelog):
    for version in changelog.versions:
        tag = version.name

        repo.create_git_release(
            tag=tag,
            name=f"v{tag}",
            message=get_release_message(version),
            draft=False,
            prerelease=False
        )


if __name__ == "__main__":
    if len(argv) < 2:
        print(f"usage: python {argv[0]} <CHANGELOG.md>")
        exit(1)

    local_repo = git.Repo(".")

    local_repo.config_writer().set_value("user", "name", "DasPoet").release()
    local_repo.config_writer().set_value("user", "email", "cederk2306@gmail.com").release()

    client = get_github_client()
    remote_repo = get_remote_repository(client)

    changelog = read_changelog(argv[1])

    create_releases(remote_repo, changelog)
