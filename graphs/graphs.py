from sys import stdin
from typing import cast

import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns

sns.set_theme()



def read_df() -> pd.DataFrame:
    df = pd.read_json(stdin)

    df["dataType"] = df["dataType"].apply(
        lambda typ: "[]" + typ.removesuffix("Slice").lower()
        if typ.endswith("Slice")
        else typ
    )
    df["millisPerOp"] = df["nanosPerOp"] // 1_000_000
    df.drop(columns=["nanosPerOp"], inplace=True)

    df = df[df["dataType"] != "[]string"]
    return cast(pd.DataFrame, df)


if __name__ == "__main__":
    df = read_df()

    for typ, group in df.groupby("benchmarkType"):
        plt.figure()

        plot = sns.barplot(
            data=group, x="dataType", y="millisPerOp", hue="benchmarkLibrary"
        )

        plt.xticks(rotation=45)

        plt.title(f"{str(typ).capitalize()}ling performance for primitive slices (lower is better)")

        plt.xlabel("Type")
        plt.ylabel("Duration (ms)")

        plot.legend(title="Library")

        plt.tight_layout()
        plt.savefig(f"benchmark_{typ}.svg", dpi=600)
