from os import getenv

from github import Github
from github.Repository import Repository


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
