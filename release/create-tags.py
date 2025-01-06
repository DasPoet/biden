import git

from sys import argv
from changelogs import parse_changelog


def create_git_tag_if_not_exists(name: str, description: str | None = None):
    repo = git.Repo(".")

    repo.config_writer().set_value("user", "name", "DasPoet").release()
    repo.config_writer().set_value("user", "email", "cederk2306@gmail.com").release()

    if name not in [tag.name for tag in repo.tags]:
        print(f"creating tag '{name}'")
        repo.create_tag(name, message=description)
    else:
        print(f"skipping tag '{name}', because it already exists")


if __name__ == "__main__":
    if len(argv) < 2:
        print(f"usage: python {argv[0]} <CHANGELOG.md>")
        exit(1)

    with open(argv[1], "r") as file:
        changelog = parse_changelog(file.read())

    for version in changelog.versions:
        description = version.to_markdown()

        create_git_tag_if_not_exists(name=version.name, description=f"Version {version.name}")
