import git

from github import Github
from github.Repository import Repository

from sys import argv
from os import getenv

from changelogs import parse_changelog, Changelog


def get_github_client() -> Github:
    token = getenv("GITHUB_TOKEN")
    return Github(token)

def get_remote_repository(client: Github) -> Repository:
    name = identify_repository()
    return client.get_repo(name)


def identify_repository() -> str:
    name = getenv("GITHUB_REPOSITORY")
    assert name is not None, "missing environment variable 'GITHUB_REPOSITORY'"

    return name


def get_unreleased_tags(client: Github, tags: list[str]) -> list[str]:
    remote_repo = get_remote_repository(client)

    released_tags = [release.tag_name for release in remote_repo.get_releases()]

    return [tag for tag in tags if tag not in released_tags]


def get_release_message(changelog: Changelog, tag: str) -> str:
    version = next(filter(lambda version: version.name == tag, changelog.versions))
    return "\n\n".join(map(lambda cl: cl.to_markdown(), version.changelists))


def create_releases(repo: Repository, tags: list[str]):
    with open(argv[1], "r") as file:
        changelog = parse_changelog(file.read())

    for tag in tags:
        repo.create_git_release(
            tag=tag,
            name=f"Release {tag}",
            message=get_release_message(changelog, tag),
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

    local_tags = [tag.name for tag in local_repo.tags]
    unreleased_tags = get_unreleased_tags(client, local_tags)

    create_releases(remote_repo, unreleased_tags)
