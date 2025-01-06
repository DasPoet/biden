import re

from dataclasses import dataclass, field


@dataclass
class Changelist:
    type: str
    changes: list[str] = field(default_factory=list)

    def to_markdown(self) -> str:
        changes = '\n'.join('- ' + change for change in self.changes)

        return f"## {self.type}\n\n{changes}"


@dataclass
class Version:
    name: str
    changelists: list[Changelist] = field(default_factory=list)

    def to_markdown(self) -> str:
        changes = map(lambda cl: cl.to_markdown(), self.changelists)

        return "# Changes\n\n" + "\n\n".join(changes)

@dataclass
class Changelog:
    versions: list[Version] = field(default_factory=list)


def parse_version_line(line: str) -> str | None:
    pattern = re.compile(r"## \[(?P<version>\d+\.\d+\.\d+)] - \d\d\d\d-\d\d-\d\d")

    match = pattern.match(line)
    if not match:
        return None

    return match.group("version")


def parse_changes_line(line: str) -> str | None:
    pattern = re.compile(r"### (?P<type>Added|Changed|Deprecated|Removed|Fixed|Security)")

    match = pattern.match(line)
    if not match:
        return None

    return match.group("type")
    

def parse_change_line(line: str) -> str | None:
    if line.startswith("* "):
        return line.removeprefix("* ")

    if line.startswith("- "):
        return line.removeprefix("- ")

    return None


def parse_changelog(content: str) -> Changelog:
    versions: list[Version] = []
    changelists_by_type: dict[str, Changelist] = {}

    last_version: Version | None = None
    last_type: str | None = None

    for line in content.splitlines():
        version = parse_version_line(line)
        if version is not None:
            if last_version is not None:
                last_version.changelists = [cl for cl in changelists_by_type.values() if len(cl.changes) > 0]
                versions.append(last_version)

            last_version = Version(name=version)

        changes_type = parse_changes_line(line)
        if changes_type is not None:
            last_type = changes_type.lower()
            changelists_by_type[last_type] = Changelist(type=changes_type)

        change = parse_change_line(line)
        if change is not None:
            assert last_type, line
            assert last_version, line

            changelists_by_type[last_type].changes.append(change)
    
    if last_version is not None:
        last_version.changelists = [cl for cl in changelists_by_type.values() if len(cl.changes) > 0]
        versions.append(last_version)

    return Changelog(versions)
