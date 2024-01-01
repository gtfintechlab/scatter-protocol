import numpy as np
from pymongo import MongoClient
import os
import matplotlib.pyplot as plt


def get_trainer_token_data(workspace_name):
    client = MongoClient("mongodb://127.0.0.1:27017")
    db = client["scatter-protocol"]

    workspace = db["workspaces"].find_one({"name": workspace_name})
    workspace_id = workspace["_id"]
    trainers = db["protocolnodes"].find(
        {"workspaceId": workspace_id, "peerType": "trainer"}
    )
    trainer_data = {}
    for index, trainer in enumerate(trainers):
        events = db["logevents"].find(
            {
                "blockchainAddress": trainer["blockchainAddress"],
                "workspaceId": workspace_id,
                "logType": "Token Balance",
            }
        )

        xData, yData = [], []
        for event in events:
            xData.append(event["xDataPoint"])
            yData.append(event["yDataPoint"])
        trainer_data[trainer["blockchainAddress"]] = {
            "x": xData,
            "y": yData,
            "type": "Trainer",
        }

    return trainer_data


def get_validator_token_data(workspace_name):
    client = MongoClient("mongodb://127.0.0.1:27017")
    db = client["scatter-protocol"]

    workspace = db["workspaces"].find_one({"name": workspace_name})
    workspace_id = workspace["_id"]
    validators = db["protocolnodes"].find(
        {"workspaceId": workspace_id, "peerType": "validator"}
    )
    validator_data = {}
    for index, validator in enumerate(validators):
        events = db["logevents"].find(
            {
                "blockchainAddress": validator["blockchainAddress"],
                "workspaceId": workspace_id,
                "logType": "Token Balance",
            }
        )

        xData, yData = [], []
        for event in events:
            xData.append(event["xDataPoint"])
            yData.append(event["yDataPoint"])
        validator_data[validator["blockchainAddress"]] = {
            "x": xData,
            "y": yData,
            "type": "Validator",
        }

    return validator_data


def create_overlayed_graph(
    data,
    output_file,
    x_axis_label,
    y_axis_label,
    fig_size,
    title,
    label_prefix=None,
    base_line=None,
    base_line_label="",
    secondary_base_line=None,
    secondary_base_line_label="",
):
    fig, ax = plt.subplots(figsize=fig_size)
    trainer_count = 0
    for index, (key, values) in enumerate(data.items()):
        if values["type"] == "Trainer":
            trainer_count += 1

        x_values = values["x"]
        y_values = values["y"]
        ax.plot(
            x_values,
            y_values,
            label=f"{label_prefix if label_prefix else values['type']} {(index+1) if values['type'] == 'Trainer' else index+1 - trainer_count}",
        )

    ax.set_xlabel(x_axis_label)
    ax.set_ylabel(y_axis_label)
    if base_line:
        ax.axhline(y=base_line, color="red", linestyle="--", label=base_line_label)

    if secondary_base_line:
        ax.axhline(
            y=secondary_base_line,
            color="blue",
            linestyle="--",
            label=secondary_base_line_label,
        )

    ax.legend()
    plt.title(title)
    plt.savefig(output_file)


def create_separated_graph(
    data,
    output_file,
    num_rows,
    num_cols,
    x_axis_label,
    y_axis_label,
    fig_size,
    title,
    label_prefix=None,
    subtitle_suffix="",
    base_line=None,
    base_line_label="",
    secondary_base_line=None,
    secondary_base_line_label="",
):
    fig, axes = plt.subplots(nrows=num_rows, ncols=num_cols, figsize=fig_size)
    if len(axes.shape) == 1:
        axes = np.array([axes])

    trainer_count = 0
    for idx, (key, values) in enumerate(data.items()):
        if values["type"] == "Trainer":
            trainer_count += 1

        row_idx = idx // num_cols
        col_idx = idx % num_cols

        x_values = values["x"]
        y_values = values["y"]

        axes[row_idx, col_idx].plot(x_values, y_values, label=title)
        axes[row_idx, col_idx].set_xlabel(x_axis_label)
        axes[row_idx, col_idx].set_ylabel(y_axis_label)

        if base_line:
            axes[row_idx, col_idx].axhline(
                y=base_line, color="red", linestyle="--", label=base_line_label
            )

        if secondary_base_line:
            axes[row_idx, col_idx].axhline(
                y=secondary_base_line,
                color="blue",
                linestyle="--",
                label=secondary_base_line_label,
            )

        axes[row_idx, col_idx].legend()
        axes[row_idx, col_idx].set_title(
            f"{label_prefix if label_prefix else values['type']} {(idx+1) if values['type'] == 'Trainer' else idx+1 - trainer_count} {subtitle_suffix}"
        )

    plt.tight_layout()
    plt.savefig(output_file)


if __name__ == "__main__":
    trainer_data = get_trainer_token_data("Paper: Scenario 1")
    validator_data = get_validator_token_data("Paper: Scenario 1")

    # Trainer Token Graphs
    create_separated_graph(
        data=trainer_data,
        output_file="trainer/scenario-1-seperated.png",
        num_cols=3,
        num_rows=2,
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(12, 8),
        title="Trainer Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=trainer_data,
        output_file="trainer/scenario-1-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(8, 6),
        title="Trainer Token Supply",
    )

    # Validator Token Graphs
    create_separated_graph(
        data=validator_data,
        output_file="validator/scenario-1-seperated.png",
        num_cols=3,
        num_rows=1,
        label_prefix="Validator",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(12, 4),
        title="Validator Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=validator_data,
        output_file="validator/scenario-1-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(8, 6),
        title="Validator Token Supply",
    )

    # Trainer Token Graphs
    create_separated_graph(
        data={**trainer_data, **validator_data},
        output_file="combined/scenario-1-seperated.png",
        num_cols=3,
        num_rows=3,
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(12, 8),
        title="Node Token Supply",
        subtitle_suffix="Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
    )
    create_overlayed_graph(
        data={**trainer_data, **validator_data},
        output_file="combined/scenario-1-overlayed.png",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(8, 6),
        title="Node Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        secondary_base_line=75000,
        secondary_base_line_label="Initial Token Supply (After Validator Stake)",
    )
